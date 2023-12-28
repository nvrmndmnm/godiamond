package main

import (
	"fmt"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/spf13/pflag"
)

type LoupeMode struct {
	commands *Command
	box      *DiamondBox
}

type SelectorsMetadata map[string]string

type LoupeFacet struct {
	FacetAddress      common.Address
	FunctionSelectors [][4]byte
}

func NewLoupeMode(box *DiamondBox) Mode {
	commands := &Command{
		Name: "loupe",
		SubCommands: []*Command{
			{
				Name:        "facets",
				Description: "Show all facets and their selectors",
			},
			{
				Name:        "addresses",
				Description: "Show all facet addresses used by a diamond",
			},
			{
				Name:        "facet-selectors",
				Description: "Show all function selectors provided by a facet",
				SubCommands: []*Command{
					{
						Name:        "address",
						Description: "Specify the Ethereum address of a facet",
					},
				},
			},
			{
				Name:        "facet-address",
				Description: "Show the facet that supports the given selector",
				SubCommands: []*Command{
					{
						Name:        "selector",
						Description: "Specify the function selector",
					},
				},
			},
			{
				Name:        "supports-interface",
				Description: "Show if the contract implements an interface",
				SubCommands: []*Command{
					{
						Name:        "id",
						Description: "Specify the interface identifier",
					},
				},
			},
		},
	}

	commands.SubCommands = append(commands.SubCommands, defaultCommands.SubCommands...)

	return &LoupeMode{commands: commands, box: box}
}

func (l *LoupeMode) GetCommands() *Command {
	return l.commands
}

func (l *LoupeMode) PrintUsage() {
	PrintUsage(l.commands)
}

func (l *LoupeMode) Execute(cmd *Command, flags *pflag.FlagSet) error {
	loupe := bind.NewBoundContract(l.box.config.Contracts["diamond"].Address,
		l.box.contracts["loupe_facet"].ABI, l.box.client, l.box.client, l.box.client)

	switch cmd.Name {
	case "facets":
		var callResult []any
		err := loupe.Call(&bind.CallOpts{}, &callResult, "facets")
		if err != nil {
			return fmt.Errorf("failed to get facets of a diamond: %v", err)
		}

		facets := *abi.ConvertType(callResult[0], new([]LoupeFacet)).(*[]LoupeFacet)

		for _, facet := range facets {

			fmt.Printf("facet address: %v\n", facet.FacetAddress)

			selectorsMetadata, err := getSelectorsMetadata(facet.FunctionSelectors)
			if err != nil {
				return fmt.Errorf("failed to get selector metadata: %v", err)
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
			return fmt.Errorf("failed to get addresses of the facets: %v", err)
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
			return fmt.Errorf("invalid address flag: %v", err)
		}

		if err := facetAddress.Set(addressString); err != nil {
			return fmt.Errorf("invalid Ethereum address format: %v", err)
		}

		err = loupe.Call(&bind.CallOpts{}, &callResult, "facetFunctionSelectors", common.Address(facetAddress))
		if err != nil {
			return fmt.Errorf("failed to get facet selectors: %v", err)
		}

		facetSelectors := *abi.ConvertType(callResult[0], new([][4]byte)).(*[][4]byte)

		selectorsMetadata, err := getSelectorsMetadata(facetSelectors)
		if err != nil {
			return fmt.Errorf("failed to get selector metadata: %v", err)
		}

		for selector, functionName := range selectorsMetadata {
			fmt.Printf("\t%s: %s \n", selector, functionName)
		}

	case "facet-address":
		var functionSelector SelectorFlag
		var callResult []any

		selectorString, err := flags.GetString("selector")
		if err != nil {
			return fmt.Errorf("invalid selector flag: %v", err)
		}

		if err := functionSelector.Set(selectorString); err != nil {
			return fmt.Errorf("invalid selector format: %v", err)
		}

		if len(functionSelector) > 1 {
			return fmt.Errorf("a single selector is required")
		}

		selector := [4]byte(functionSelector[0])

		err = loupe.Call(&bind.CallOpts{}, &callResult, "facetAddress", selector)
		if err != nil {
			return fmt.Errorf("failed to get facet address: %v", err)
		}

		facetAddress := *abi.ConvertType(callResult[0], new(common.Address)).(*common.Address)

		fmt.Printf("Facet address: %s\n", facetAddress.String())

	case "supports-interface":
		var interfaceId SelectorFlag
		var callResult []any

		interfaceIdString, err := flags.GetString("id")
		if err != nil {
			return fmt.Errorf("invalid id flag: %v", err)
		}

		if err := interfaceId.Set(interfaceIdString); err != nil {
			return fmt.Errorf("invalid id format: %v", err)
		}

		if len(interfaceId) > 1 {
			return fmt.Errorf("a single identifier is required")
		}

		id := [4]byte(interfaceId[0])

		err = loupe.Call(&bind.CallOpts{}, &callResult, "supportsInterface", id)
		if err != nil {
			return fmt.Errorf("failed to check interface support: %v", err)
		}

		status := *abi.ConvertType(callResult[0], new(bool)).(*bool)

		fmt.Printf("ERC-165 status: %v\n", status)
	}

	return nil
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
