package main

import (
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
)

func TestSelectorFlag_String(t *testing.T) {
	selectorFlag := &SelectorFlag{
		[4]byte{0x01, 0x02, 0x03, 0x04},
		[4]byte{0x05, 0x06, 0x07, 0x08},
	}

	expected := "[4]uint8{0x1, 0x2, 0x3, 0x4}, [4]uint8{0x5, 0x6, 0x7, 0x8}"
	actual := selectorFlag.String()

	assert.Equal(t, expected, actual)
}

func TestSelectorFlag_Set(t *testing.T) {
	selectorFlag := &SelectorFlag{}

	err := selectorFlag.Set("0x01020304, 0x05060708")
	assert.NoError(t, err)

	expected := SelectorFlag{
		[4]byte{0x01, 0x02, 0x03, 0x04},
		[4]byte{0x05, 0x06, 0x07, 0x08},
	}

	assert.Equal(t, expected, *selectorFlag)
}

func TestAddressFlag_Set(t *testing.T) {
	addressFlag := &AddressFlag{}

	err := addressFlag.Set("0xdeadbeefdeadbeefdeadbeefdeadbeefdeadbeef")
	assert.NoError(t, err)

	expected := AddressFlag(common.HexToAddress("0xdeadbeefdeadbeefdeadbeefdeadbeefdeadbeef"))

	assert.Equal(t, expected, *addressFlag)
}
