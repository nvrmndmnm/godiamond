package main

import (
	"errors"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"go.uber.org/zap"
)

type MockBoundContract struct{}

func (m *MockBoundContract) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	tx := types.NewTransaction(0, common.Address{}, big.NewInt(0), 0, big.NewInt(0), nil)
	return tx, nil
}

func setupBox() (*DiamondBox, error) {
	mockEth := new(MockEthereumWrapper)
	mockClient := &ethclient.Client{}
	mockEth.On("Dial", "http://localhost:8545").Return(mockClient, nil)
	mockEth.On("Dial", "http://some-other-url").Return(nil, errors.New("failed to connect"))

	tmpFile, err := os.CreateTemp("", "")
	if err != nil {
		return nil, err
	}
	defer os.Remove(tmpFile.Name())

	testJsonData := `{
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
			},
			{
				"type": "function",
				"name": "init",
				"inputs": [],
				"outputs": [],
				"stateMutability": "nonpayable"
			},
			{
				"type": "function",
				"name": "diamondCut",
				"inputs": [
				  {
					"name": "_diamondCut",
					"type": "tuple[]",
					"internalType": "struct IDiamondCut.FacetCut[]",
					"components": [
					  {
						"name": "facetAddress",
						"type": "address",
						"internalType": "address"
					  },
					  {
						"name": "action",
						"type": "uint8",
						"internalType": "enum IDiamondCut.FacetCutAction"
					  },
					  {
						"name": "functionSelectors",
						"type": "bytes4[]",
						"internalType": "bytes4[]"
					  }
					]
				  },
				  {
					"name": "_init",
					"type": "address",
					"internalType": "address"
				  },
				  {
					"name": "_calldata",
					"type": "bytes",
					"internalType": "bytes"
				  }
				],
				"outputs": [],
				"stateMutability": "nonpayable"
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

	if _, err := tmpFile.Write([]byte(testJsonData)); err != nil {
		return nil, err
	}
	tmpFile.Close()

	config := Config{
		RPC: map[string]string{
			"test": "http://localhost:8545",
		},
		Accounts: map[string]EOA{
			"anvil": {
				PrivateKey: "0xcafebabecafebabecafebabecafebabecafebabecafebabecafebabecafebabe",
			},
		},
		Contracts: map[string]ContractConfig{
			"test": {
				MetadataFilePath: tmpFile.Name(),
			},
			"diamond": {
				MetadataFilePath: tmpFile.Name(),
			},
			"diamond_init": {
				MetadataFilePath: tmpFile.Name(),
			},
			"cut_facet": {
				MetadataFilePath: tmpFile.Name(),
			},
			"loupe_facet": {
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
		return nil, err
	}

	return box, nil
}
