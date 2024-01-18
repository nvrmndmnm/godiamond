package main

import (
	"errors"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/stretchr/testify/mock"
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

func setupMockCutContract() *MockBoundContract {
	mockContract := new(MockBoundContract)

	var functionSelectors SelectorFlag
	err := functionSelectors.Set("0xdeadbeef")
	if err != nil {
		return nil
	}

	expectedCut := []FacetCut{{
		FacetAddress:      common.HexToAddress("0xCAFEBABECAFEBABECAFEBABECAFEBABECAFEBABE"),
		Action:            Add,
		FunctionSelectors: functionSelectors,
	}}
	expectedAddress := common.HexToAddress("0xdeadbeefdeadbeefdeadbeefdeadbeefdeadbeef")
	expectedCalldata := []byte{225, 199, 57, 42}

	tx := types.NewTransaction(0, common.Address{}, big.NewInt(0), 0, big.NewInt(0), nil)

	mockContract.On(
		"Transact",
		mock.Anything,
		"diamondCut",
		expectedCut,
		expectedAddress,
		expectedCalldata).
		Return(tx, nil)

	return mockContract
}
