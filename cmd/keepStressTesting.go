package main

import (
	"context"
	"encoding/hex"
	"fmt"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/spf13/cobra"
	"github.com/tyler-smith/go-bip32"
	"github.com/tyler-smith/go-bip39"
	"math/big"
	"sync/atomic"
	"time"
)

func transferStressCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "transferAT",
		Short: "stress transfer ju all time",
		Run:   transferAT,
	}
	addTransferStressFlags(cmd)
	return cmd
}

func addTransferStressFlags(cmd *cobra.Command) {
	cmd.Flags().IntP("repeat", "r", 1, "repeat number")
	_ = cmd.MarkFlagRequired("repeat")

	cmd.Flags().StringP("mnemonic", "m", "", "mnemonic")
	_ = cmd.MarkFlagRequired("mnemonic")
}

func signTxAT(masterKey *bip32.Key, sender *multiSender, signChan chan *types.Transaction) {

	var count int64
	for {

		bkey, err := newKeyFromMasterKey(masterKey, TypeEther, bip32.FirstHardenedChild, 0, uint32(count))
		if err != nil {
			fmt.Println("Failed to newKeyFromMasterKey with err:", err)
			panic(err)
		}

		cpub, err := crypto.DecompressPubkey(bkey.PublicKey().Key)
		if err != nil {
			fmt.Println("Failed to DecompressPubkey with err:", err)
			panic(err)
		}
		addr := crypto.PubkeyToAddress(*cpub)
		tx := sender.SignJuTx(addr, sender.nonce_start+uint64(count), big.NewInt(1e10))
		count++
		signChan <- tx

	}
}

var runChan = make(chan bool, 50)

func waitSignTxAT(signChan chan *types.Transaction, sender *multiSender) {
	for {
		if len(signChan) > int(sender.repeat) {
			for tx := range signChan {
				sender.recvTxChan <- tx
				if len(sender.recvTxChan) == int(sender.repeat) {
					sender.cycle++
					//send signal for mutisend
					//每间隔20000 笔交易就发送信号批量集中发送达到压测的目的
					for i := 0; i < sender.proceeNum; i++ {
						runChan <- true

					}
					break
				}
			}
		}
		time.Sleep(1 * time.Second)
	}
}

func waitSendTxAT(sender *multiSender) {

	if sender.proceeNum == 0 {
		panic("must set proceeNum")
	}
	for i := 0; i < sender.proceeNum; i++ {
		go func(index int) {
			for {

				<-runChan
				for {
					tx, ok := <-sender.recvTxChan
					if ok {
						err := sender.client.SendTransaction(context.Background(), tx)
						if err != nil {
							fmt.Println("Failed to sendTxAT with err:", err, "will retry...")
							err = sender.client.SendTransaction(context.Background(), tx)
							if err != nil {
								fmt.Println("Failed to sendTxAT with err:", err)
							}
						}
						atomic.AddInt64(&sender.sendTxNum, 1)
						if len(sender.recvTxChan) == 0 {
							break
						}
						fmt.Println("send tx at:", tx.Hash().Hex(), "processNum:", index, "tx.Nonce:", tx.Nonce())

					} else {
						return
					}
				}

			}
		}(i)
	}
}

func transferAT(cmd *cobra.Command, args []string) {
	rpcLaddr, _ := cmd.Flags().GetString("rpc_laddr")
	mnemonic, _ := cmd.Flags().GetString("mnemonic")
	repeat, _ := cmd.Flags().GetInt("repeat")
	signChan := make(chan *types.Transaction, 30000)
	seed, err := bip39.NewSeedWithErrorChecking(mnemonic, "")
	if err != nil {
		fmt.Println("NewSeedWithErrorChecking with error:", err.Error())
		return
	}
	masterKey, err := bip32.NewMasterKey(seed)
	masterpriv, err := crypto.ToECDSA(masterKey.Key)
	if nil != err {
		fmt.Println("NewSeedWithErrorChecking with error:", err.Error())
		return
	}

	ethSender := crypto.PubkeyToAddress(masterpriv.PublicKey)
	dpub, _ := crypto.DecompressPubkey(masterKey.PublicKey().Key)
	ethSender2 := crypto.PubkeyToAddress(*dpub)
	fmt.Println("master ethSender:", ethSender.String())
	fmt.Println("master ethSender:", ethSender2.String())
	masterPrivHex := hex.EncodeToString(masterKey.Key)
	fmt.Println("masterPriv:", masterPrivHex)

	//开始批量构造交易
	client, err := ethclient.Dial(rpcLaddr)
	if err != nil {
		fmt.Println("Failed to connect to the Ethereum client with err:", err)
		return
	}

	nonce, err := client.NonceAt(context.Background(), ethSender, nil)
	if err != nil {
		fmt.Println("client.NonceAt with err:", err)
		return
	}
	fmt.Println("current nonce:", nonce, "sender:", ethSender)
	mSender := new(multiSender)
	mSender.client = client
	mSender.sender = ethSender
	mSender.senderKey = masterpriv
	mSender.nonce_start = nonce
	mSender.repeat = int32(repeat)
	mSender.recvTxChan = make(chan *types.Transaction, 20000)
	mSender.proceeNum = 50
	go signTxAT(masterKey, mSender, signChan)
	go waitSignTxAT(signChan, mSender)
	go waitSendTxAT(mSender)
	processStart := time.Now()
	for {
		fmt.Println("signChan capacity", len(signChan), "recvTxChan capacity", len(mSender.recvTxChan), "total send tx:", mSender.sendTxNum, "cycle times:", mSender.cycle, "cost time:", time.Since(processStart))
		time.Sleep(time.Second)
	}

}
