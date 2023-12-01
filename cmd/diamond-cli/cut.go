package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/nvrmndmnm/godiamond/internal/contracts"
)

type Action uint8

const (
	Add     Action = 0
	Replace Action = 1
	Remove  Action = 2
)

func (box *DiamondBox) cut(facetAddress common.Address, action Action, selectors [][4]byte) error {
	var cut []contracts.IDiamondCutFacetFacetCut

	cut = append(cut, contracts.IDiamondCutFacetFacetCut{
		FacetAddress:      facetAddress,
		Action:            uint8(action),
		FunctionSelectors: selectors,
	})

	diamondCut, err := contracts.NewDiamondCutFacet(box.diamond, box.client)
	if err != nil {
		log.Fatal(err)
	}

	abiInstance, err := abi.JSON(strings.NewReader(contracts.DiamondInitABI))
	if err != nil {
		log.Fatal(err)
	}

	calldata, err := abiInstance.Pack("init")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(calldata)

	tx, err := diamondCut.DiamondCut(box.auth, cut, box.diamondInit, calldata)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(tx)

	return nil
}
