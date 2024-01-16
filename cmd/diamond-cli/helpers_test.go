package main

import (
	"errors"
	"math/big"

	"github.com/ethereum/go-ethereum/ethclient"
	"go.uber.org/zap"
)

func setupBox() (*DiamondBox, error) {
	mockEth := new(MockEthereumWrapper)
	mockClient := &ethclient.Client{}
	mockEth.On("Dial", "http://localhost:8545").Return(mockClient, nil)
	mockEth.On("Dial", "http://some-other-url").Return(nil, errors.New("failed to connect"))

	testConfig, err := loadConfig("testdata/config_test.yaml")
	if err != nil {
		return nil, err
	}

	sugar := zap.NewExample().Sugar()
	modeName := "cut"
	rpcName := "test"
	chainId := big.NewInt(-1)

	box, err := NewDiamondBox(testConfig, sugar, modeName, rpcName, chainId)
	if err != nil {
		return nil, err
	}

	return box, nil
}
