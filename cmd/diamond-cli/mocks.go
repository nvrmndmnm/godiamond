package main

import (
	"context"
	"crypto/ecdsa"
	"errors"
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
	return nil, nil
}

func (m *MockEthereumWrapper) NetworkID(ctx context.Context) (*big.Int, error) {
	return big.NewInt(1), nil
}

func (m *MockEthereumWrapper) SuggestGasPrice(ctx context.Context) (*big.Int, error) {
	return big.NewInt(21000), nil
}

func (m *MockEthereumWrapper) HexToECDSA(hexkey string) (*ecdsa.PrivateKey, error) {
	return nil, errors.New("HexToECDSA failed")
}

// MockBoundContract is a mock object that implements the BoundContract interface.
type MockBoundContract struct{}

func (m *MockBoundContract) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	tx := types.NewTransaction(0, common.Address{}, big.NewInt(0), 0, big.NewInt(0), nil)
	return tx, nil
}
