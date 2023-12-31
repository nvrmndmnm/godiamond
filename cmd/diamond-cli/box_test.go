package main

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"math/big"
	"os"
	"reflect"
	"strings"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.uber.org/zap"
)

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

func TestNewDiamondBox(t *testing.T) {
	mockEth := new(MockEthereumWrapper)
	mockClient := &ethclient.Client{}
	mockEth.On("Dial", "http://localhost:8545").Return(mockClient, nil)
	mockEth.On("Dial", "http://some-other-url").Return(nil, errors.New("failed to connect"))

	tmpFile, err := os.CreateTemp("", "")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmpFile.Name())

	jsonData := `{
		"abi":[
			{
				"type":"constructor",
				"inputs":[
					{
					"name":"_contractOwner",
					"type":"address",
					"internalType":"address"
					},
					{
					"name":"_diamondCutFacet",
					"type":"address",
					"internalType":"address"
					}
				],
				"stateMutability":"payable"
			},
			{
				"type":"fallback",
				"stateMutability":"payable"
			},
			{
				"type":"receive",
				"stateMutability":"payable"
			}
		],
		"bytecode":{
			"object":"0x60806040526040516110696e2066"
		},
		"methodIdentifiers":{
			"diamondCut((address,uint8,bytes4[])[],address,bytes)":"1f931c1c"
		},
		"ast":{
			"id":43437,
			"nodes":[{"id":43321,"name":"Diamond"}]
		}
    }`

	if _, err := tmpFile.Write([]byte(jsonData)); err != nil {
		t.Fatal(err)
	}
	tmpFile.Close()

	config := Config{
		RPC: map[string]string{
			"test": "http://localhost:8545",
		},
		Accounts: map[string]EOA{
			"anvil": {
				PrivateKey: "0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80",
			},
		},
		Contracts: map[string]ContractConfig{
			"test": {
				MetadataFilePath: tmpFile.Name(),
			},
		},
	}
	
	sugar := zap.NewExample().Sugar()
	modeName := "cut"
	rpcName := "test"
	chainId := big.NewInt(-1)

	box, err := NewDiamondBox(config, sugar, modeName, rpcName, chainId)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	assert.Equal(t, config, box.config)
	assert.Equal(t, sugar, box.sugar)
	assert.Equal(t, modeName, box.mode.GetCommands().Name)

	//box Eth: client, auth, chainid
	assert.Equal(t, big.NewInt(1), box.eth.chainId)
	// assert.Equal(t, mockClient, box.eth.client)

	expectedABIStr := `[
		{
			"type":"constructor",
			"inputs":[
				{
					"name":"_contractOwner",
					"type":"address",
					"internalType":"address"
				},
				{
					"name":"_diamondCutFacet",
					"type":"address",
					"internalType":"address"
				}
			],
			"stateMutability":"payable"
		},
		{
			"type":"fallback",
			"stateMutability":"payable"
		},
		{
			"type":"receive",
			"stateMutability":"payable"
		}
	]`

	expectedABI, err := abi.JSON(strings.NewReader(expectedABIStr))
	if err != nil {
		t.Fatalf("failed to parse expected ABI: %v", err)
	}

	if !reflect.DeepEqual(expectedABI, box.contracts["test"].ABI) {
		t.Errorf("wrong ABI; got %v, want %v", box.contracts["test"].ABI, expectedABI)
	}

	expectedBytecode := "0x60806040526040516110696e2066"
	if box.contracts["test"].Bytecode.Object != expectedBytecode {
		t.Errorf("wrong bytecode; got %v, want %v", box.contracts["test"].Bytecode.Object, expectedBytecode)
	}

	expectedSelectors := SelectorsMetadata{
		"diamondCut((address,uint8,bytes4[])[],address,bytes)":"1f931c1c",
	}
	if !reflect.DeepEqual(expectedSelectors, box.contracts["test"].MethodIdentifiers) {
		t.Errorf("wrong selectors; got %v, want %v", box.contracts["test"].MethodIdentifiers, expectedSelectors)
	}

	expectedName := "Diamond"
	if box.contracts["test"].AST.Nodes[0].Name != expectedName {
		t.Errorf("wrong selectors; got %v, want %v", box.contracts["test"].AST.Nodes[0].Name, expectedName)
	}
}
