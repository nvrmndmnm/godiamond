package main

import (
	"context"
	"crypto/ecdsa"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/stretchr/testify/mock"
)

// MockEthereumWrapper is a mock object that implements the EthereumWrapper interface.
type MockEthereumWrapper struct {
	mock.Mock
}

func (m *MockEthereumWrapper) Dial(rawurl string) (*ethclient.Client, error) {
	args := m.Called(rawurl)
	return args.Get(0).(*ethclient.Client), args.Error(1)
}

func (m *MockEthereumWrapper) NewKeyedTransactorWithChainID(key *ecdsa.PrivateKey, chainID *big.Int) (*bind.TransactOpts, error) {
	args := m.Called(key, chainID)
	return args.Get(0).(*bind.TransactOpts), args.Error(1)
}

func (m *MockEthereumWrapper) NetworkID(ctx context.Context) (*big.Int, error) {
	args := m.Called(ctx)
	return args.Get(0).(*big.Int), args.Error(1)
}

func (m *MockEthereumWrapper) SuggestGasPrice(ctx context.Context) (*big.Int, error) {
	args := m.Called(ctx)
	return args.Get(0).(*big.Int), args.Error(1)
}

func (m *MockEthereumWrapper) HexToECDSA(hexkey string) (*ecdsa.PrivateKey, error) {
	args := m.Called(hexkey)
	return args.Get(0).(*ecdsa.PrivateKey), args.Error(1)
}

func (m *MockEthereumWrapper) Close() {
	m.Called()
}

func (m *MockEthereumWrapper) SetClient(client *ethclient.Client) {
	m.Called(client)
}

func (m *MockEthereumWrapper) GetClient() *ethclient.Client {
	args := m.Called()
	return args.Get(0).(*ethclient.Client)
}

// MockBoundContract is a mock object that implements the BoundContract interface.
type MockBoundContract struct {
	mock.Mock
}

func (m *MockBoundContract) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	var args mock.Arguments
	switch {
	case method == "diamondCut":
		args = m.Called(opts, method, (params[0]).([]FacetCut), (params[1]).(common.Address), (params[2]).([]byte))
	default:
		args = m.Called(opts, method, params)
	}
	return args.Get(0).(*types.Transaction), args.Error(1)
}

func (m *MockBoundContract) Call(opts *bind.CallOpts, results *[]interface{}, method string, params ...interface{}) error {
	var args mock.Arguments

	switch {
	case method == "facets":
		args = m.Called(opts, results, method)
	default:
		args = m.Called(opts, method, params)
	}
	return args.Error(0)
}
