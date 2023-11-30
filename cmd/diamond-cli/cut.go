package main

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/nvrmndmnm/godiamond/internal/contracts"
)

type Action uint8

const (
	Add     Action = 0
	Replace Action = 1
	Remove  Action = 2
)

func cut(facetAddress common.Address, action Action, selectors [][4]byte) error {

	fmt.Println("cut")
	var cut []*contracts.IDiamondCutFacetCut

	cut = append(cut, &contracts.IDiamondCutFacetCut{
		FacetAddress:      facetAddress,
		Action:            uint8(action),
		FunctionSelectors: selectors,
	})

	return nil
}
