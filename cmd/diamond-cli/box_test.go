package main

import (
	"math/big"
	"reflect"
	"strings"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/stretchr/testify/assert"
)

func TestNewDiamondBox(t *testing.T) {
	box, err := setupBox()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

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

	expectedSelectors := map[string]string{
		"test((address,uint8)address,bytes)": "bc645d96",
	}
	if !reflect.DeepEqual(expectedSelectors, box.contracts["test"].MethodIdentifiers) {
		t.Errorf("wrong selectors; got %v, want %v", box.contracts["test"].MethodIdentifiers, expectedSelectors)
	}

	expectedName := "TestContract"
	if box.contracts["test"].AST.Nodes[0].Name != expectedName {
		t.Errorf("wrong selectors; got %v, want %v", box.contracts["test"].AST.Nodes[0].Name, expectedName)
	}
}
