package diamond

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
	box, err := SetupBox()
	assert.NoError(t, err)

	assert.Equal(t, &ethclient.Client{}, box.Eth.Client)
	assert.Equal(t, &bind.TransactOpts{GasPrice: &big.Int{}}, box.Eth.Auth)
	assert.Equal(t, big.NewInt(-1), box.Eth.ChainId)

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
	assert.Equal(t, expectedABI, box.Contracts["test"].ABI)

	expectedBytecode := "0x60806040526040516110696e2066"
	assert.Equal(t, expectedBytecode, box.Contracts["test"].Bytecode.Object)

	expectedSelectors := map[string]string{
		"test((address,uint8)address,bytes)": "bc645d96",
	}
	assert.Equal(t, expectedSelectors, box.Contracts["test"].MethodIdentifiers)

	expectedName := "TestContract"
	assert.Equal(t, expectedName, box.Contracts["test"].AST.Nodes[0].Name)
}
