package main

import (
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/spf13/pflag"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestNewCutMode(t *testing.T) {
	box, err := setupBox()
	if err != nil {
		t.Fatalf("Failed to create DiamondBox: %v", err)
	}

	mode := NewCutMode(box)
	cutMode, ok := mode.(*CutMode)
	if !ok {
		t.Fatalf("Expected type CutMode, got %T", mode)
	}

	assert.NotNil(t, cutMode)
	assert.Equal(t, box, cutMode.box)
	assert.Equal(t, "cut", cutMode.commands.Name)
}

func TestCutMode_GetCommands(t *testing.T) {
	box, err := setupBox()
	if err != nil {
		t.Fatalf("Failed to create DiamondBox: %v", err)
	}

	mode := NewCutMode(box)
	cutMode, ok := mode.(*CutMode)
	if !ok {
		t.Fatalf("Expected type CutMode, got %T", mode)
	}

	assert.Equal(t, cutMode.commands, cutMode.GetCommands())
}

func TestCutMode_Execute(t *testing.T) {
	box, err := setupBox()
	if err != nil {
		t.Fatalf("Failed to create DiamondBox: %v", err)
	}

	mode := NewCutMode(box)
	cutMode, ok := mode.(*CutMode)
	if !ok {
		t.Fatalf("Expected type CutMode, got %T", mode)
	}

	mockContract := setupMockCutContract()

	cutMode.cutContract = mockContract
	cmd := &Command{Name: "add"}
	flags := pflag.NewFlagSet("test", pflag.ContinueOnError)
	flags.String("address", "0xFEEDBABEFEEDBABEFEEDBABEFEEDBABEFEEDBABE", "")
	flags.String("selectors", "0xbc645d96", "")

	var cut []FacetCut
	var facetAddress AddressFlag
	var functionSelectors SelectorFlag

	calldata, err := box.contracts["diamond_init"].ABI.Pack("init")
	assert.Nil(t, err)

	addressString, err := flags.GetString("address")
	assert.Nil(t, err)

	err = facetAddress.Set(addressString)
	assert.Nil(t, err)

	selectorString, err := flags.GetString("selectors")
	assert.Nil(t, err)

	err = functionSelectors.Set(selectorString)
	assert.Nil(t, err)

	cut = append(cut, FacetCut{
		FacetAddress:      common.Address(facetAddress),
		Action:            Add,
		FunctionSelectors: functionSelectors,
	})

	err = cutMode.Execute(cmd, flags)

	diamondInitAddress := common.HexToAddress(box.config.Contracts["diamond_init"].Address)
	mockContract.AssertCalled(t, "Transact", mock.Anything, "diamondCut",
		cut,
		diamondInitAddress,
		calldata)
	assert.Nil(t, err)
}

func TestCutMode_Execute_InvalidAddress(t *testing.T) {
	box, err := setupBox()
	if err != nil {
		t.Fatalf("Failed to create DiamondBox: %v", err)
	}

	mode := NewCutMode(box)
	cmd := &Command{Name: "add"}
	flags := pflag.NewFlagSet("test", pflag.ContinueOnError)
	flags.String("address", "invalid", "")
	flags.String("selectors", "0x00000000", "")

	err = mode.Execute(cmd, flags)

	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "invalid Ethereum address format")
}

func TestCutMode_Execute_InvalidSelectors(t *testing.T) {
	box, err := setupBox()
	if err != nil {
		t.Fatalf("Failed to create DiamondBox: %v", err)
	}

	mode := NewCutMode(box)
	cmd := &Command{Name: "add"}
	flags := pflag.NewFlagSet("test", pflag.ContinueOnError)
	flags.String("address", "0xCAFEBABECAFEBABECAFEBABECAFEBABECAFEBABE", "")
	flags.String("selectors", "invalid", "")

	err = mode.Execute(cmd, flags)

	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "invalid selector format")
}

func TestCutMode_Execute_FailedToCutDiamond(t *testing.T) {
	box, err := setupBox()
	if err != nil {
		t.Fatalf("Failed to create DiamondBox: %v", err)
	}

	mode := NewCutMode(box)
	cutMode, ok := mode.(*CutMode)
	if !ok {
		t.Fatalf("Expected type CutMode, got %T", mode)
	}

	mockContract := setupMockCutContract()

	cutMode.cutContract = mockContract
	cmd := &Command{Name: "remove"}
	flags := pflag.NewFlagSet("test", pflag.ContinueOnError)
	flags.String("address", "0xFEEDBABEFEEDBABEFEEDBABEFEEDBABEFEEDBABE", "")
	flags.String("selectors", "0xbc645d96", "")

	var cut []FacetCut
	var facetAddress AddressFlag
	var functionSelectors SelectorFlag

	calldata, err := box.contracts["diamond_init"].ABI.Pack("init")
	assert.Nil(t, err)

	addressString, err := flags.GetString("address")
	assert.Nil(t, err)

	err = facetAddress.Set(addressString)
	assert.Nil(t, err)

	selectorString, err := flags.GetString("selectors")
	assert.Nil(t, err)

	err = functionSelectors.Set(selectorString)
	assert.Nil(t, err)

	cut = append(cut, FacetCut{
		Action:            Remove,
		FunctionSelectors: functionSelectors,
	})

	err = cutMode.Execute(cmd, flags)

	diamondInitAddress := common.HexToAddress(box.config.Contracts["diamond_init"].Address)
	mockContract.AssertCalled(t, "Transact", mock.Anything, "diamondCut",
		cut,
		diamondInitAddress,
		calldata)
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "failed to cut diamond")
}
