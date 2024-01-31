package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"
	"strconv"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

type BoundContract interface {
	Call(opts *bind.CallOpts, results *[]interface{}, method string, params ...interface{}) error
	Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error)
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

func convertStringParamsToType(strParams []string, types abi.Arguments) ([]interface{}, error) {
	params := make([]interface{}, len(strParams))
	var err error

	for i, value := range strParams {
		switch types[i].Type.T {
		case abi.AddressTy:
			if !common.IsHexAddress(value) {
				return nil, fmt.Errorf("%s is not a valid Ethereum address", value)
			}
			params[i] = common.HexToAddress(value)

		case abi.BoolTy:
			params[i], err = strconv.ParseBool(value)
			if err != nil {
				return nil, err
			}

		case abi.BytesTy:
			params[i] = []byte(value)

		case abi.IntTy, abi.UintTy:
			res, ok := new(big.Int).SetString(value, 0)
			if !ok {
				return nil, fmt.Errorf("failed to convert to big.Int: %s", value)
			}
			params[i] = res

		case abi.StringTy:
			params[i] = value
		}
	}

	return params, nil
}
