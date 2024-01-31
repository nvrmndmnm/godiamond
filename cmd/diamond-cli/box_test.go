package main

import (
	"math/big"
	"strings"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/stretchr/testify/assert"
)

func TestNewDiamondBox(t *testing.T) {
	box, err := setupBox()
	assert.NoError(t, err)

	assert.Equal(t, &ethclient.Client{}, box.eth.client)
	assert.Equal(t, &bind.TransactOpts{GasPrice: &big.Int{}}, box.eth.auth)
	assert.Equal(t, big.NewInt(-1), box.eth.chainId)

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
	assert.NoError(t, err)
	assert.Equal(t, expectedABI, box.contracts["test"].ABI)

	expectedBytecode := "0x60806040526040516110696e2066"
	assert.Equal(t, expectedBytecode, box.contracts["test"].Bytecode.Object)

	expectedSelectors := map[string]string{
		"test((address,uint8)address,bytes)": "bc645d96",
	}
	assert.Equal(t, expectedSelectors, box.contracts["test"].MethodIdentifiers)

	expectedName := "TestContract"
	assert.Equal(t, expectedName, box.contracts["test"].AST.Nodes[0].Name)
}
