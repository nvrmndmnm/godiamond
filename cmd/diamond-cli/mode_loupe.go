package main

import (
	"fmt"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/spf13/pflag"
)

type SelectorsMetadata map[string]string

type LoupeFacet struct {
	FacetAddress      common.Address
	FunctionSelectors [][4]byte
}

func (box *DiamondBox) modeLoupe(cmd *Command, flags *pflag.FlagSet) {

	loupe := bind.NewBoundContract(box.config.Contracts["diamond"].Address,
		box.contracts["loupe_facet"].ABI, box.client, box.client, box.client)

	switch cmd.Name {
	case "facets":
		var callResult []any
		err := loupe.Call(&bind.CallOpts{}, &callResult, "facets")
		if err != nil {
			fmt.Println("error getting facets of a diamond:", err)
			return
		}

		facets := *abi.ConvertType(callResult[0], new([]LoupeFacet)).(*[]LoupeFacet)

		for _, facet := range facets {

			fmt.Printf("facet address: %v\n", facet.FacetAddress)

			selectorsMetadata, err := getSelectorsMetadata(facet.FunctionSelectors)
			if err != nil {
				fmt.Println("could not retrieve selector metadata", err)
				return
			}
			for selector, functionName := range selectorsMetadata {
				fmt.Printf("\t%s: %s \n", selector, functionName)
			}

			fmt.Println()
		}

	case "addresses":
		var callResult []any
		err := loupe.Call(&bind.CallOpts{}, &callResult, "facetAddresses")
		if err != nil {
			fmt.Println("getting addresses of the facets", err)
		}

		facetAddreses := *abi.ConvertType(callResult[0], new([]common.Address)).(*[]common.Address)

		for _, address := range facetAddreses {
			fmt.Println(address.String())
		}

	case "facet-selectors":
		var facetAddress AddressFlag
		var callResult []any

		addressString, err := flags.GetString("address")
		if err != nil {
			fmt.Println("invalid address flag")
			return
		}

		if err := facetAddress.Set(addressString); err != nil {
			fmt.Printf("invalid Ethereum address format: %v\n", err)
			return
		}

		err = loupe.Call(&bind.CallOpts{}, &callResult, "facetFunctionSelectors", common.Address(facetAddress))
		if err != nil {
			fmt.Println("getting facet selectors", err)
		}

		facetSelectors := *abi.ConvertType(callResult[0], new([][4]byte)).(*[][4]byte)

		selectorsMetadata, err := getSelectorsMetadata(facetSelectors)
		if err != nil {
			fmt.Println("could not retrieve selector metadata", err)
			return
		}

		for selector, functionName := range selectorsMetadata {
			fmt.Printf("\t%s: %s \n", selector, functionName)
		}

	case "facet-address":
		var functionSelector SelectorFlag
		var callResult []any

		selectorString, err := flags.GetString("selector")
		if err != nil {
			fmt.Println("invalid selector flag")
			return
		}

		if err := functionSelector.Set(selectorString); err != nil {
			fmt.Printf("invalid selector format: %v\n", err)
			return
		}

		if len(functionSelector) > 1 {
			fmt.Println("provide a single selector")
			return
		}

		selector := [4]byte(functionSelector[0])

		err = loupe.Call(&bind.CallOpts{}, &callResult, "facetAddress", selector)
		if err != nil {
			fmt.Println("getting facet address", err)
		}

		facetAddress := *abi.ConvertType(callResult[0], new(common.Address)).(*common.Address)

		fmt.Println("facet address: ", facetAddress.String())

	case "supports-interface":
		var interfaceId SelectorFlag
		var callResult []any

		interfaceIdString, err := flags.GetString("id")
		if err != nil {
			fmt.Println("invalid id flag")
			return
		}
		
		if err := interfaceId.Set(interfaceIdString); err != nil {
			fmt.Printf("invalid id format: %v\n", err)
			return
		}

		if len(interfaceId) > 1 {
			fmt.Println("provide a single identifier")
			return
		}

		id := [4]byte(interfaceId[0])

		err = loupe.Call(&bind.CallOpts{}, &callResult, "supportsInterface", id)
		if err != nil {
			fmt.Println("checking interface support", err)
		}

		status := *abi.ConvertType(callResult[0], new(bool)).(*bool)

		fmt.Println("ERC-165 status:", status)
	}
}

func getSelectorsMetadata(selectors [][4]byte) (SelectorsMetadata, error) {
	selectorsMetadata := make(SelectorsMetadata)
	for _, selector := range selectors {
		selectorString := hexutil.Encode(selector[:])
		// TODO: Since there is no easy way to retrieve the function signature from the
		// selector, a persistent metadata of every contract deployment is needed
		functionName := "TBD"
		selectorsMetadata[selectorString] = functionName
	}
	return selectorsMetadata, nil
}
