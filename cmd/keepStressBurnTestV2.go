package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/spf13/cobra"
	"github.com/tyler-smith/go-bip32"
	"github.com/tyler-smith/go-bip39"
	"github.com/zhengjunhe/l3-test/cmd/contracts/contracts4juchain/generated"
	"math/big"
	"runtime"
	"sync"
	"time"
)

func crossBurnStressCmdV2() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "burnStressV2",
		Short: "burn from ju all of the time",
		Run:   burnTokenATV2,
	}
	addBurnStressV2Flags(cmd)
	return cmd
}

func addBurnStressV2Flags(cmd *cobra.Command) {

	cmd.Flags().IntP("chainID", "w", 97, "cross to dst chain ID")

	cmd.Flags().StringP("token", "t", "", "token contract address")
	cmd.Flags().StringP("registerAddr", "c", "", "registerAddr contract address")
	cmd.Flags().IntP("repeat", "r", 3000, "repeat number")

	_ = cmd.MarkFlagRequired("registerAddr")

	cmd.Flags().StringP("mnemonic", "m", "", "mnemonic")
	_ = cmd.MarkFlagRequired("mnemonic")

}

func burnTokenATV2(cmd *cobra.Command, args []string) {
	rpcLaddr, _ := cmd.Flags().GetString("rpc_laddr")
	token, _ := cmd.Flags().GetString("token")
	registerAddr, _ := cmd.Flags().GetString("registerAddr")
	repeat, _ := cmd.Flags().GetInt("repeat")
	chainIDWd, err := cmd.Flags().GetInt("chainID")

	fmt.Println("registerAddr:", registerAddr)
	fmt.Println("token:", token)

	if err != nil {
		panic(err)
	}
	mnemonic, _ := cmd.Flags().GetString("mnemonic")
	seed, err := bip39.NewSeedWithErrorChecking(mnemonic, "")
	if err != nil {
		fmt.Println("NewSeedWithErrorChecking with error:", err.Error())
		return
	}

	masterKey, err := bip32.NewMasterKey(seed)
	if err != nil {
		panic(err)
	}
	client, x2EthContracts, x2EthDeployInfo, err := RecoverContractHandler4Dstju(registerAddr, rpcLaddr)
	if err != nil {
		fmt.Println("RecoverContractHandler err:", err)
		return
	}

	var bSender *burnSender
	bSender = &burnSender{
		client:         client,
		bridgeBankIns:  x2EthContracts.BridgeBank,
		bridgeBankAddr: x2EthDeployInfo.BridgeBank.Address,
		tokenAddr:      common.HexToAddress(token),
		amount:         big.NewInt(1),
		chainID2wd:     big.NewInt(int64(chainIDWd)),
		repeat:         repeat,
		recvTxChan:     make(chan *types.Transaction, 2000),
		proceeNum:      50,
		nodeUrl:        rpcLaddr,
		//child:          make(chan *childKeyAddr, 3000),
	}

	bSender.genMutiKeyAddr(masterKey)

	signChan := make(chan *types.Transaction, 30000)
	//循环批量签名
	go signBurnTxATV2(bSender, signChan)
	//等待签名数量达到一定后往发送管道输送
	go waitSignBurnTxATV2(signChan, bSender)
	go waitSendBurnTxAT(bSender)
	processStart := time.Now()

	for {
		fmt.Println("signChan capacity", len(signChan), "recvTxChan capacity", len(bSender.recvTxChan), "total send tx:", bSender.sendTxNum, "cycle times:", bSender.cycle, "cost time:", time.Since(processStart))
		time.Sleep(time.Second)
	}
}

func (b *burnSender) genMutiKeyAddr(masterKey *bip32.Key) {
	for i := 0; i < b.repeat; i++ {
		bkey, err := newKeyFromMasterKey(masterKey, TypeEther, bip32.FirstHardenedChild, 0, uint32(i))
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
		signKey, err := crypto.ToECDSA(bkey.Key)
		if err != nil {
			panic(err)
		}

		b.child = append(b.child, &childKeyAddr{addr: addr, key: signKey})
		//b.child <- &childKeyAddr{addr: addr, key: signKey}
	}
}

func signBurnTxATV2(sender *burnSender, sigChan chan *types.Transaction) {
	numCPUs := runtime.NumCPU()
	txCnt := len(sender.child)
	addrPerGoroutime := txCnt / numCPUs
	fmt.Println("runtime.NumCPU()", numCPUs, "addrPerGoroutime", addrPerGoroutime, "sender num:", txCnt)
	time.Sleep(time.Second * 3)
	for {
		var wg sync.WaitGroup
		if addrPerGoroutime == 0 {
			wg.Add(1)
			for i := 0; i < len(sender.child); i++ {
				txs, err := SignBurn(sender, sender.child[i].addr, sender.child[i].key)
				if err != nil {
					panic(err)

				}
				for _, tx := range txs {
					sigChan <- tx
				}

			}

		} else {
			for i := 0; i < numCPUs-1; i++ {
				partOfchild := sender.child[i*addrPerGoroutime : (i+1)*addrPerGoroutime]
				wg.Add(1)
				go multiSign(partOfchild, sender, sigChan, &wg)

			}

			partOfchild := sender.child[(numCPUs-1)*addrPerGoroutime:]
			wg.Add(1)
			go multiSign(partOfchild, sender, sigChan, &wg)

		}

		wg.Wait()
		fmt.Println("wait for sign completed++++,cycle:", sender.cycle)
		sender.cycle++
	}

}

func multiSign(childKey []*childKeyAddr, bSender *burnSender, sigChan chan *types.Transaction, wg *sync.WaitGroup) {

	for _, child := range childKey {
		txs, err := SignBurn(bSender, child.addr, child.key)
		if err != nil {
			panic(err)

		}
		//fmt.Println("cycle:", bSender.cycle, "txs:", len(txs))
		//fmt.Println("child addr:", child.addr, "index:", index, "cycle:", bSender.cycle, "txsize:", len(txs), "hash:", txs[0].Hash())
		for _, tx := range txs {
			sigChan <- tx

		}

	}
	wg.Done()
}

func SignBurn(bSender *burnSender, senderAddr common.Address, signKey *ecdsa.PrivateKey) ([]*types.Transaction, error) {
	var err error
	var txs []*types.Transaction
	opts := &bind.CallOpts{
		Pending: false,
		From:    common.Address{},
		Context: context.Background(),
	}

	if bSender.bridgeServiceFee == nil {
		bridgeServiceFee, err := bSender.bridgeBankIns.BridgeServiceFee(opts)
		if nil != err {
			return nil, err
		}
		bSender.bridgeServiceFee = bridgeServiceFee
	}

	//aprove
	//tokenAddr := common.HexToAddress(tokenAddrstr)
	if bSender.cycle == 0 { //第一次需要approve
		auth, err := PrepareAuth(bSender.client, signKey, senderAddr)
		if nil != err {

			return nil, err
		}

		tokenInstance, err := generated.NewBridgeToken(bSender.tokenAddr, bSender.client)
		if nil != err {

			panic(err)
		}
		auth.NoSend = true
		//btcbank 是bridgeBank的基类，所以使用bridgeBank的地址
		//fmt.Println("gas:", auth.GasLimit, "sender:", senderAddr)

		tx, err := tokenInstance.Approve(auth, bSender.bridgeBankAddr, big.NewInt(1e18))
		if nil != err {
			panic(err)
		}
		txs = append(txs, tx)
	}

	//fmt.Println("chainIDWd:", b.chainID2wd, "receiver:", receiver.Hex(), "tokenAddr:", b.tokenAddr, "amount:", b.amount)
	auth, err := PrepareAuth(bSender.client, signKey, senderAddr)
	if nil != err {
		return nil, err
	}
	auth.NoSend = true

	tx, err := bSender.bridgeBankIns.BurnBridgeTokens(auth, bSender.chainID2wd, senderAddr, bSender.tokenAddr, bSender.amount)
	if nil != err {
		return nil, err
	}
	txs = append(txs, tx)
	return txs, nil
}

func waitSignBurnTxATV2(sigChan chan *types.Transaction, sender *burnSender) {
	for {

		if sender.cycle == 0 {
			if len(sigChan) >= sender.repeat*2 {
				for tx := range sigChan {
					sender.recvTxChan <- tx
					if len(sender.recvTxChan) == sender.repeat*2 {
						//send signal for mutisend
						//每间隔3000笔交易就发送信号批量集中发送达到压测的目的
						for i := 0; i < sender.proceeNum; i++ {
							runChan <- true

						}
						break
					}
				}

			}

		} else {

			if len(sigChan) >= sender.repeat {

				for tx := range sigChan {

					sender.recvTxChan <- tx

					if len(sender.recvTxChan) == sender.repeat {
						fmt.Println("<<<<<<<<<<<<<<<<<<<<<<<<<<++collect tx completed will send signal ")
						//send signal for mutisend
						//每间隔3000笔交易就发送信号批量集中发送达到压测的目的
						for i := 0; i < sender.proceeNum; i++ {
							runChan <- true

						}
						break
					}
				}
			}

		}
		//time.Sleep(1 * time.Second)
	}
}
