package main

import (
	"encoding/hex"
	"fmt"
	"strings"

	"github.com/ethereum/go-ethereum/common"
)

type SelectorFlag [][4]byte
type AddressFlag common.Address

func (b *SelectorFlag) String() string {
	var result []string
	
	for _, item := range *b {
		result = append(result, fmt.Sprintf("%#v", item))
	}

	return strings.Join(result, ", ")
}

func (b *SelectorFlag) Set(value string) error {
	trimmedValue := strings.TrimSpace(value)
	trimmedValue = strings.TrimPrefix(trimmedValue, "0x")
	
	values := strings.Split(trimmedValue, ",")

	for _, v := range values {
		decoded, err := hex.DecodeString(strings.TrimSpace(v))

		if err != nil {
			return fmt.Errorf("error decoding hex: %v", err)
		}

		if len(decoded) != 4 {
			return fmt.Errorf("each value must be a 4-byte array")
		}

		*b = append(*b, [4]byte(decoded))
	}

	return nil
}

func (b *SelectorFlag) Type() string {
	return fmt.Sprintf("%T", b)
}

func (f *AddressFlag) Set(value string) error {
	trimmedValue := strings.TrimSpace(value)
	address := common.HexToAddress(trimmedValue)

	if address == (common.Address{}) {
		return fmt.Errorf("invalid Ethereum address: %s", trimmedValue)
	}
	*f = AddressFlag(address)

	return nil
}
