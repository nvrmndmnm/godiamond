package main

import (
	"fmt"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/spf13/pflag"
)

type FacetCut struct {
	FacetAddress      common.Address
	Action            uint8
	FunctionSelectors [][4]byte
}

const (
	Add     uint8 = 0
	Replace uint8 = 1
	Remove  uint8 = 2
)

func (box *DiamondBox) modeCut(cmd *Command, flags *pflag.FlagSet) {
	diamondCut := bind.NewBoundContract(box.config.Contracts["diamond"].Address,
		box.contracts["cut_facet"].ABI, box.client, box.client, box.client)

	calldata, err := box.contracts["diamond_init"].ABI.Pack("init")
	if err != nil {
		fmt.Println(err)
	}

	var cut []FacetCut

	switch cmd.Name {
	case "add", "replace", "remove":
		var action uint8
		var facetAddress AddressFlag
		var functionSelectors SelectorFlag

		addressString, err := flags.GetString("address")
		if err != nil {
			fmt.Println("invalid address flag")
			return
		}

		selectorString, err := flags.GetString("selectors")
		if err != nil {
			fmt.Println("invalid selector flag")
			return
		}

		if cmd.Name == "add" || cmd.Name == "replace" {
			if err := facetAddress.Set(addressString); err != nil {
				fmt.Printf("invalid Ethereum address format: %v\n", err)
				return
			}
		}

		if err := functionSelectors.Set(selectorString); err != nil {
			fmt.Printf("invalid selector format: %v\n", err)
			return
		}

		switch cmd.Name {
		case "add":
			action = Add
		case "replace":
			action = Replace
		case "remove":
			action = Remove
		}

		cut = append(cut, FacetCut{
			FacetAddress:      common.Address(facetAddress),
			Action:            action,
			FunctionSelectors: functionSelectors,
		})

		tx, err := diamondCut.Transact(box.auth, "diamondCut", cut,
			box.config.Contracts["diamond_init"].Address, calldata)
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println(tx.Hash())
	}
}
