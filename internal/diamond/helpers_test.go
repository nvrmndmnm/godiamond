package diamond

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/nvrmndmnm/godiamond/internal/cli"
	"github.com/nvrmndmnm/godiamond/internal/config"
	"github.com/nvrmndmnm/godiamond/internal/ethereum"
	"github.com/stretchr/testify/mock"
	"go.uber.org/zap"
)

func SetupBox() (*DiamondBox, error) {
	mockEth := new(MockEthereumWrapper)

	mockEth.On("Dial", "http://localhost:6969").Return(&ethclient.Client{}, nil)
	mockEth.On("Dial", "http://some-other-url").Return(nil, errors.New("failed to connect"))

	testConfig, err := config.LoadConfig("../../testdata/config_test.yaml")
	if err != nil {
		return nil, err
	}

	sugar := zap.NewExample().Sugar()
	rpcName := "test"

	box := &DiamondBox{
		Config:    testConfig,
		Sugar:     sugar,
		Eth:       &ethereum.EthereumWrapper{},
		Contracts: map[string]ContractMetadata{},
	}

	for contractIdentifier, metadataPath := range testConfig.Metadata {
		contractMetadata, err := GetContractMetadataByFile(metadataPath)
		if err != nil {
			return nil, err
		}
		box.Contracts[contractIdentifier] = contractMetadata
	}

	mockClient, _ := mockEth.Dial(box.Config.RPC[rpcName])
	box.Eth.Client = mockClient

	chainId := big.NewInt(-1)
	box.Eth.ChainId = chainId

	hexkey := testConfig.PrivateKey[2:]
	mockEth.On("HexToECDSA", hexkey).Return(&ecdsa.PrivateKey{}, nil)
	privateKey, _ := mockEth.HexToECDSA(hexkey)

	mockEth.On("NewKeyedTransactorWithChainID", privateKey, box.Eth.ChainId).Return(&bind.TransactOpts{}, nil)
	auth, _ := mockEth.NewKeyedTransactorWithChainID(privateKey, box.Eth.ChainId)

	mockEth.On("SuggestGasPrice", context.Background()).Return(&big.Int{}, nil)
	gasPrice, _ := mockEth.SuggestGasPrice(context.Background())
	auth.GasPrice = gasPrice

	box.Eth.Auth = auth

	return box, nil
}

func SetupMockCutContract() *MockBoundContract {
	mockContract := new(MockBoundContract)

	var functionSelectors cli.SelectorFlag
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

	tx := types.NewTransaction(0, common.Address{}, big.NewInt(0), 0, big.NewInt(0), nil)

	mockContract.On(
		"Transact",
		mock.Anything,
		"diamondCut",
		expectedCut,
		common.Address{},
		[]byte{}).
		Return(tx, nil)

	mockContract.On(
		"Transact",
		mock.Anything,
		"diamondCut",
		expectedErrCut,
		common.Address{},
		[]byte{}).
		Return(tx, errors.New("failed test"))

	return mockContract
}

func SetupMockLoupeContract() *MockBoundContract {
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
