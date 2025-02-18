package ethinterface

import (
	"context"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient/simulated"
	"math/big"
)

// EthClientSpec ...
type EthClientSpec interface {
	bind.ContractBackend
	TransactionReceipt(ctx context.Context, txHash common.Hash) (*types.Receipt, error)
	NetworkID(ctx context.Context) (*big.Int, error)
	BalanceAt(ctx context.Context, account common.Address, blockNumber *big.Int) (*big.Int, error)
	HeaderByNumber(ctx context.Context, number *big.Int) (*types.Header, error)
	SyncProgress(ctx context.Context) (*ethereum.SyncProgress, error)
	TransactionByHash(ctx context.Context, hash common.Hash) (tx *types.Transaction, isPending bool, err error)
}

// SimExtend ...
type SimExtend struct {
	simulated.Client
	*simulated.Backend
	//*backends.SimulatedBackend
}

// HeaderByNumber ...
func (sim *SimExtend) HeaderByNumber(ctx context.Context, number *big.Int) (*types.Header, error) {
	currentBlock, err := sim.BlockNumber(ctx)
	if err != nil {
		return nil, err
	}
	return sim.HeaderByNumber(ctx, big.NewInt(int64(currentBlock)))

	//return sim.Blockchain().CurrentBlock(), nil
}

// NetworkID ...
func (sim *SimExtend) NetworkID(ctx context.Context) (*big.Int, error) {
	return big.NewInt(1337), nil
}

func (sim *SimExtend) SyncProgress(ctx context.Context) (*ethereum.SyncProgress, error) {
	return nil, nil
}
