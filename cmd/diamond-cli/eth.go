package main

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
	Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error)
}

type EthereumClient interface {
	Dial(rawurl string) (*ethclient.Client, error)
	NewKeyedTransactorWithChainID(key *ecdsa.PrivateKey, chainID *big.Int) (*bind.TransactOpts, error)
	NetworkID(ctx context.Context) (*big.Int, error)
	SuggestGasPrice(ctx context.Context) (*big.Int, error)
	HexToECDSA(hexkey string) (*ecdsa.PrivateKey, error)
}

type EthereumWrapper struct {
	client  *ethclient.Client
	auth    *bind.TransactOpts
	chainId *big.Int
}

func (eth *EthereumWrapper) Dial(rawurl string) (*ethclient.Client, error) {
	return ethclient.Dial(rawurl)
}

func (eth *EthereumWrapper) NewKeyedTransactorWithChainID(key *ecdsa.PrivateKey, chainID *big.Int) (*bind.TransactOpts, error) {
	return bind.NewKeyedTransactorWithChainID(key, chainID)
}

func (eth *EthereumWrapper) NetworkID(ctx context.Context) (*big.Int, error) {
	return eth.client.NetworkID(ctx)
}

func (eth *EthereumWrapper) SuggestGasPrice(ctx context.Context) (*big.Int, error) {
	return eth.client.SuggestGasPrice(ctx)
}

func (eth *EthereumWrapper) HexToECDSA(hexkey string) (*ecdsa.PrivateKey, error) {
	return crypto.HexToECDSA(hexkey)
}

func (eth *EthereumWrapper) Close() {
	eth.client.Close()
}
