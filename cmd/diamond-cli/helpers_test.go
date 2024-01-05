package main

import (
	"errors"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/ethclient"
	"go.uber.org/zap"
)

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
