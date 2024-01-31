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
	"go.uber.org/zap"
)

func setupBox() (*DiamondBox, error) {
	mockEth := new(MockEthereumWrapper)

	mockEth.On("Dial", "http://localhost:6969").Return(&ethclient.Client{}, nil)
	mockEth.On("Dial", "http://some-other-url").Return(nil, errors.New("failed to connect"))

	testConfig, err := loadConfig("testdata/config_test.yaml")
	if err != nil {
		return nil, err
	}

	sugar := zap.NewExample().Sugar()
	rpcName := "test"

	box := &DiamondBox{
		config:    testConfig,
		sugar:     sugar,
		eth:       &EthereumWrapper{},
		contracts: map[string]ContractMetadata{},
	}

	for contractIdentifier, contractConfig := range testConfig.Contracts {
		contractMetadata, err := getContractMetadataByFile(contractConfig.MetadataFilePath)
		if err != nil {
			return nil, err
		}
		box.contracts[contractIdentifier] = contractMetadata
	}

	mockClient, _ := mockEth.Dial(box.config.RPC[rpcName])
	box.eth.client = mockClient

	chainId := big.NewInt(-1)
	box.eth.chainId = chainId

	hexkey := testConfig.Accounts["anvil"].PrivateKey[2:]
	mockEth.On("HexToECDSA", hexkey).Return(&ecdsa.PrivateKey{}, nil)
	privateKey, _ := mockEth.HexToECDSA(hexkey)

	mockEth.On("NewKeyedTransactorWithChainID", privateKey, box.eth.chainId).Return(&bind.TransactOpts{}, nil)
	auth, _ := mockEth.NewKeyedTransactorWithChainID(privateKey, box.eth.chainId)

	mockEth.On("SuggestGasPrice", context.Background()).Return(&big.Int{}, nil)
	gasPrice, _ := mockEth.SuggestGasPrice(context.Background())
	auth.GasPrice = gasPrice

	box.eth.auth = auth

	return box, nil
}

func setupMockCutContract() *MockBoundContract {
	mockContract := new(MockBoundContract)

	var functionSelectors SelectorFlag
	err := functionSelectors.Set("0xbc645d96")
	if err != nil {
		return nil
	}

	expectedCut := []FacetCut{{
		FacetAddress:      common.HexToAddress("0xFEEDBABEFEEDBABEFEEDBABEFEEDBABEFEEDBABE"),
		Action:            Add,
		FunctionSelectors: functionSelectors,
	}}

	expectedErrCut := []FacetCut{{
		Action:            Remove,
		FunctionSelectors: functionSelectors,
	}}

	diamondInitAddress := common.HexToAddress("0xB055BABEB055BABEB055BABEB055BABEB055BABE")
	expectedCalldata := []byte{225, 199, 57, 42}
	tx := types.NewTransaction(0, common.Address{}, big.NewInt(0), 0, big.NewInt(0), nil)

	mockContract.On(
		"Transact",
		mock.Anything,
		"diamondCut",
		expectedCut,
		diamondInitAddress,
		expectedCalldata).
		Return(tx, nil)

	mockContract.On(
		"Transact",
		mock.Anything,
		"diamondCut",
		expectedErrCut,
		diamondInitAddress,
		expectedCalldata).
		Return(tx, errors.New("failed test"))

	return mockContract
}

func setupMockLoupeContract() *MockBoundContract {
	mockContract := new(MockBoundContract)
	mockContract.On("Call", &bind.CallOpts{}, mock.Anything, "facets").Run(func(args mock.Arguments) {
		// args[1] is the additional argument to Call, which stores the call results
		results := args[1].(*[]interface{})
		*results = append(*results, []LoupeFacet{
			{
				FacetAddress:      common.HexToAddress("0xFEEDBABEFEEDBABEFEEDBABEFEEDBABEFEEDBABE"),
				FunctionSelectors: [][4]byte{{188, 100, 93, 150}},
			},
		})
	}).Return(nil)

	return mockContract
}
