package main

import (
	"context"
	"crypto/ecdsa"
	"encoding/hex"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/spf13/cobra"
	"github.com/tyler-smith/go-bip32"
	"github.com/tyler-smith/go-bip39"
	"log"
	"math/big"
	"runtime"
	"sync"
	"sync/atomic"
	"time"
)

func transferCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "transfer",
		Short: "transfer ju",
		Run:   transfer,
	}
	addTransferFlags(cmd)
	return cmd
}

func addTransferFlags(cmd *cobra.Command) {
	cmd.Flags().IntP("repeat", "r", 1, "repeat number")
	_ = cmd.MarkFlagRequired("repeat")

	cmd.Flags().StringP("mnemonic", "m", "", "mnemonic")
	_ = cmd.MarkFlagRequired("mnemonic")
}

func transfer(cmd *cobra.Command, args []string) {
	rpcLaddr, _ := cmd.Flags().GetString("rpc_laddr")
	repeat, _ := cmd.Flags().GetInt("repeat")
	mnemonic, _ := cmd.Flags().GetString("mnemonic")

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

	fmt.Println("masterPriv:", hex.EncodeToString(masterKey.Key))

	//开始批量构造交易
	client, err := ethclient.Dial(rpcLaddr)
	if err != nil {
		fmt.Println("Failed to connect to the Ethereum client with err:", err)
		return
	}
	//nonce, err := client.PendingNonceAt(context.Background(), ethSender)

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
	mSender.nonde_start = nonce
	mSender.recvTxChan = make(chan *types.Transaction, 50)

	var privkeyArr []*ecdsa.PrivateKey
	var txs []*types.Transaction
	for i := 0; i < repeat; i++ {
		bkey, err := newKeyFromMasterKey(masterKey, TypeEther, bip32.FirstHardenedChild, 0, uint32(i))
		if err != nil {
			fmt.Println("Failed to newKeyFromMasterKey with err:", err)
			return
		}
		newKey, err := crypto.ToECDSA(bkey.Key)
		privkeyArr = append(privkeyArr, newKey)
		cpub, err := crypto.DecompressPubkey(bkey.PublicKey().Key)
		if err != nil {
			fmt.Println("Failed to DecompressPubkey with err:", err)
			return
		}
		addr := crypto.PubkeyToAddress(*cpub)
		//fmt.Println("genAddr:", addr, "genIndex", i)
		tx := mSender.SignJuTx(addr, mSender.nonde_start+uint64(i), big.NewInt(1e10))
		txs = append(txs, tx)
	}
	fmt.Println("+++++++total signed txnum:", len(txs))
	time.Sleep(time.Second * 2)
	mSender.proceeNum = 50
	for i := 0; i < mSender.proceeNum; i++ {
		go func(index int) {
			mSender.sendJuTxV2(index)
		}(i)
	}
	go func() {
		for _, tx := range txs {
			mSender.recvTxChan <- tx
		}
	}()

	for {
		time.Sleep(time.Millisecond * 300)
		fmt.Println("[      send txNum     ]:", mSender.sendTxNum, "process num:", mSender.proceeNum)
		if mSender.sendTxNum == len(txs) {
			close(mSender.recvTxChan)
			break
		}
	}

	for {
		fmt.Println("process num:", mSender.proceeNum)
		if mSender.proceeNum == 0 {
			break
		}
		time.Sleep(time.Second)
	}

	//blockOccupided := make(map[int64]bool)

	fmt.Println("Going to check tx status")
	// check tx process
	numCPUs := runtime.NumCPU()
	txCnt := len(txs)
	txsPerGoroutime := txCnt / numCPUs

	var transferStatics TransferStatics
	transferStatics.TxResult = make(map[string]TxStatus)
	transferStatics.Blocks = make(map[int64]bool)

	checkTxProcessed(txs, client)

	var wg sync.WaitGroup

	if txsPerGoroutime == 0 {
		wg.Add(1)
		go checkTxStatus(txs, client, &transferStatics, &wg)
	} else {
		for i := 0; i < numCPUs-1; i++ {
			wg.Add(1)
			txs2txs2check := txs[i*txsPerGoroutime : (i+1)*txsPerGoroutime]
			go checkTxStatus(txs2txs2check, client, &transferStatics, &wg)
		}

		wg.Add(1)
		txs2txs2check := txs[(numCPUs-1)*txsPerGoroutime:]
		go checkTxStatus(txs2txs2check, client, &transferStatics, &wg)
	}
	wg.Wait()

	//fmt.Println("transferStatics.SuccessCnt", transferStatics.SuccessCnt, transferStatics.Blocks)
	//fmt.Println("transferStatics", transferStatics)

	now := "./" + fmt.Sprintf("%d", time.Now().Unix())
	var transferStatics2DB TransferStatics2DB

	transferStatics2DB.Name = "transferStatics"
	transferStatics2DB.TxResult = make(map[string]TxStatus)
	transferStatics2DB.Blocks = make(map[int64]bool)
	for hash, val := range transferStatics.TxResult {
		transferStatics2DB.TxResult[hash] = val
	}
	for block, val := range transferStatics.Blocks {
		transferStatics2DB.Blocks[block] = val
	}

	transferStatics2DB.SuccessCnt = int(transferStatics.SuccessCnt)
	transferStatics2DB.FailedCnt = int(transferStatics.FailedCnt)
	transferStatics2DB.SendCnt = len(txs)

	writeToFile(now, transferStatics2DB)
	//writeCoinStatics(now, transferStatics2DB)
}

func checkTxProcessed(txs2check []*types.Transaction, client *ethclient.Client) {
	cnt := 0
	for {
		for _, tx := range txs2check {
			receipt, err := client.TransactionReceipt(context.Background(), tx.Hash())
			if err == ethereum.NotFound {
				time.Sleep(time.Second)
				cnt = 0
				continue
			}
			cnt++
			if cnt > 10 && cnt%10 == 0 {
				fmt.Println("checkTxProcessed has processed txs", cnt, "block height", receipt.BlockNumber.Int64())
			}
			continue
		}
		if cnt == len(txs2check) {
			break
		}
		cnt = 0
	}

}

func checkTxStatus(txs2check []*types.Transaction, client *ethclient.Client, transferStatics *TransferStatics, wg *sync.WaitGroup) {
	defer wg.Done()
	sleepCnt := 0
	for _, tx := range txs2check {
		receipt, err := client.TransactionReceipt(context.Background(), tx.Hash())
		if err == ethereum.NotFound {
			fmt.Println("Not find tx with hash", tx.Hash().String())
			transferStatics.RWLock.Lock()
			transferStatics.TxResult[tx.Hash().String()] = TxStatus{
				ErrorInfo: err.Error(),
			}
			transferStatics.RWLock.Unlock()
			time.Sleep(time.Second)
			sleepCnt++
			continue
		} else if err != nil {
			transferStatics.RWLock.Lock()
			transferStatics.TxResult[tx.Hash().String()] = TxStatus{
				ErrorInfo: err.Error(),
			}
			transferStatics.RWLock.Unlock()
			fmt.Println("tx with error", err.Error())
			atomic.AddInt32(&transferStatics.FailedCnt, 1)
			continue
		}

		if receipt.Status != types.ReceiptStatusSuccessful {
			transferStatics.RWLock.Lock()
			transferStatics.Blocks[receipt.BlockNumber.Int64()] = true
			transferStatics.TxResult[tx.Hash().String()] = TxStatus{
				Block:  receipt.BlockNumber.Int64(),
				Status: int(types.ReceiptStatusFailed),
			}
			transferStatics.RWLock.Unlock()
			atomic.AddInt32(&transferStatics.FailedCnt, 1)
			fmt.Println("tx failed with statue", receipt.Status)
			continue
		}

		transferStatics.RWLock.Lock()
		transferStatics.Blocks[receipt.BlockNumber.Int64()] = true
		transferStatics.TxResult[tx.Hash().String()] = TxStatus{
			Block:  receipt.BlockNumber.Int64(),
			Status: int(types.ReceiptStatusSuccessful),
		}
		transferStatics.RWLock.Unlock()
		atomic.AddInt32(&transferStatics.SuccessCnt, 1)
	}
}

type TxStatus struct {
	Block     int64  `json:"block,omitempty"`
	Status    int    `json:"status,omitempty"`
	ErrorInfo string `json:"errorInfo,omitempty"`
}

type TransferStatics struct {
	Name       string
	RWLock     sync.RWMutex
	TxResult   map[string]TxStatus
	Blocks     map[int64]bool
	SuccessCnt int32
	FailedCnt  int32
}

type TransferStatics2DB struct {
	Name       string              `json:"name,omitempty"`
	SuccessCnt int                 `json:"successCnt"`
	SendCnt    int                 `json:"sendCnt,omitempty"`
	FailedCnt  int                 `json:"failedCnt"`
	Blocks     map[int64]bool      `json:"blocks,omitempty"`
	TxResult   map[string]TxStatus `json:"txResult,omitempty"`
}

type multiSender struct {
	nonde_start uint64
	client      *ethclient.Client
	sender      common.Address
	senderKey   *ecdsa.PrivateKey
	recvTxChan  chan *types.Transaction
	proceeNum   int
	sendTxNum   int
	mutex       sync.Mutex
	txMutex     sync.Mutex
	atoTxNum    int32
}

func (m *multiSender) sendJuTxV2(i int) {
	defer func() {
		m.mutex.Lock()
		m.proceeNum--
		m.mutex.Unlock()
	}()
	for {
		tx, ok := <-m.recvTxChan
		if ok {
			err := m.client.SendTransaction(context.Background(), tx)
			if err != nil {
				m.txMutex.Lock()
				m.sendTxNum++
				m.txMutex.Unlock()
				log.Fatalf("Failed to send transaction: %v", err)
				return
			}
			m.txMutex.Lock()
			m.sendTxNum++
			m.txMutex.Unlock()
			atomic.AddInt32(&m.atoTxNum, 1)
			//fmt.Println("Transaction sent:", tx.Hash().Hex(), "process index:", i, "atoTxnum:", m.atoTxNum, "m.sendTxNum:", m.sendTxNum)
		} else {
			//println("channel closed")
			return
		}
	}
}

var gasPrice *big.Int

func (m *multiSender) SignJuTx(to common.Address, nonce uint64, amount *big.Int) *types.Transaction {

	gasLimit := uint64(21000) // 交易的Gas Limit
	var err error

	if gasPrice == nil {
		gasPrice, err = m.client.SuggestGasPrice(context.Background())
		if err != nil {
			log.Fatalf("Failed to suggest gas price: %v", err)
		}
	}

	// 5. 创建交易
	tx := types.NewTransaction(nonce, to, amount, gasLimit, gasPrice, nil)
	// 6. 使用私钥签名交易

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(big.NewInt(chainID)), m.senderKey)
	if err != nil {
		log.Fatalf("Failed to sign transaction: %v", err)
	}
	return signedTx
}

const Purpose uint32 = 0x8000002C
const TypeEther uint32 = 0x8000003c
const chainID = 66633666

func newKeyFromMasterKey(masterKey *bip32.Key, coin, account, chain, address uint32) (*bip32.Key, error) {
	child, err := masterKey.NewChildKey(Purpose)
	if err != nil {
		return nil, err
	}

	child, err = child.NewChildKey(coin)
	if err != nil {
		return nil, err
	}

	child, err = child.NewChildKey(account)
	if err != nil {
		return nil, err
	}

	child, err = child.NewChildKey(chain)
	if err != nil {
		return nil, err
	}

	child, err = child.NewChildKey(address)
	if err != nil {
		return nil, err
	}

	return child, nil
}
