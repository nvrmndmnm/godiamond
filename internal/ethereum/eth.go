package ethereum

import (
	"context"
	"crypto/ecdsa"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

type BoundContract interface {
	Call(opts *bind.CallOpts, results *[]interface{}, method string, params ...interface{}) error
	Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error)
}

type EthereumWrapper struct {
	Client  *ethclient.Client
	Auth    *bind.TransactOpts
	ChainId *big.Int
}

func (eth *EthereumWrapper) Dial(rawurl string) (*ethclient.Client, error) {
	return ethclient.Dial(rawurl)
}

func (eth *EthereumWrapper) NewKeyedTransactorWithChainID(key *ecdsa.PrivateKey, chainID *big.Int) (*bind.TransactOpts, error) {
	return bind.NewKeyedTransactorWithChainID(key, chainID)
}

func (eth *EthereumWrapper) NetworkID(ctx context.Context) (*big.Int, error) {
	return eth.Client.NetworkID(ctx)
}

func (eth *EthereumWrapper) SuggestGasPrice(ctx context.Context) (*big.Int, error) {
	return eth.Client.SuggestGasPrice(ctx)
}

func (eth *EthereumWrapper) HexToECDSA(hexkey string) (*ecdsa.PrivateKey, error) {
	return crypto.HexToECDSA(hexkey)
}

func (eth *EthereumWrapper) Close() {
	eth.Client.Close()
}
