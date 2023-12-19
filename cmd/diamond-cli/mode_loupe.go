package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/c-bata/go-prompt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/nvrmndmnm/godiamond/internal/facets"
	"github.com/spf13/pflag"
)

type SelectorsMetadata map[string]string

func printLoupeUsage() {
	var usage = `
Commands:
    facets                        Show all facets and their selectors
    addresses                     Show all facet addresses used by a diamond
    facet-selectors <address>     Show all function selectors provided by a facet
    facet-address <selector>      Show the facet that supports the given selector
    help                          Show help
    exit                          Exit the loupe mode

Arguments:
    --address           string    Ethereum address of the facet
    --selector          string    4-byte function selector representation 
`
	fmt.Print(usage)
}

func loupeCompleter(d prompt.Document) []prompt.Suggest {
	s := []prompt.Suggest{
		{Text: "facets", Description: "Show all facets and their selectors"},
		{Text: "addresses", Description: "Show all facet addresses used by a diamond"},
		{Text: "facet-selectors", Description: "Show all function selectors provided by a facet"},
		{Text: "facet-address", Description: "Show the facet that supports the given selector"},
		{Text: "help", Description: "Show help message"},
		{Text: "exit", Description: "Exit the loupe interactive mode"},
	}

	args := strings.Split(d.Text, " ")

	if len(args) <= 1 {
		return prompt.FilterHasPrefix(s, d.GetWordBeforeCursor(), true)
	}

	if len(args) == 2 {
		switch args[0] {
		case "facet-selectors":
			return []prompt.Suggest{
				{Text: "--address=", Description: "Specify the Ethereum address of a facet"},
			}

		case "facet-address":
			return []prompt.Suggest{
				{Text: "--selector=", Description: "Specify the function selector"},
			}
		}
	}

	return []prompt.Suggest{}
}

func (box *DiamondBox) loupeExecutor(s string) {
	s = strings.TrimSpace(s)
	args := strings.Split(s, " ")

	loupe, err := facets.NewDiamondLoupeFacet(box.diamond, box.client)
	if err != nil {
		fmt.Println(err)
	}

	switch args[0] {
	case "facets":
		facets, err := loupe.Facets(&bind.CallOpts{})
		if err != nil {
			fmt.Println("Error: getting facets of a diamond", err)
		}

		for _, facet := range facets {
			fmt.Printf("facet address: %v\n", facet.FacetAddress)

			selectorsMetadata, err := getSelectorsMetadata(facet.FunctionSelectors)
			if err != nil {
				fmt.Println("Error: could not retrieve selector metadata", err)
				return
			}
			for selector, functionName := range selectorsMetadata {
				fmt.Printf("\t%s: %s \n", selector, functionName)
			}

			fmt.Println()
		}

	case "addresses":
		facetAddreses, err := loupe.FacetAddresses(&bind.CallOpts{})
		if err != nil {
			fmt.Println("Error: getting addresses of the facets", err)
		}

		for _, address := range facetAddreses {
			fmt.Println(address.String())
		}

	case "facet-selectors":
		var facetAddress AddressFlag
		var addressString string

		flags := pflag.NewFlagSet("facet-selectors", pflag.ContinueOnError)
		flags.StringVarP(&addressString, "address", "", "", "Facet address")
		err := flags.Parse(args[1:])

		if err != nil {
			fmt.Println("Error: invalid arguments for facet-selectors command")
			return
		}

		if err := facetAddress.Set(addressString); err != nil {
			fmt.Printf("Error: invalid Ethereum address format: %v\n", err)
			return
		}

		facetSelectors, err := loupe.FacetFunctionSelectors(&bind.CallOpts{},
			common.Address(facetAddress))
		if err != nil {
			fmt.Println("Error: getting facet selectors", err)
		}

		selectorsMetadata, err := getSelectorsMetadata(facetSelectors)
		if err != nil {
			fmt.Println("Error: could not retrieve selector metadata", err)
			return
		}
		for selector, functionName := range selectorsMetadata {
			fmt.Printf("\t%s: %s \n", selector, functionName)
		}

	case "facet-address":
		var functionSelector SelectorFlag
		var selectorString string

		flags := pflag.NewFlagSet("facet-address", pflag.ContinueOnError)
		flags.StringVarP(&selectorString, "selector", "", "", "Function selector")
		err := flags.Parse(args[1:])

		if err != nil {
			fmt.Println("Error: invalid arguments for facet-address command")
			return
		}

		if err := functionSelector.Set(selectorString); err != nil {
			fmt.Printf("Error: invalid selector format: %v\n", err)
			return
		}

		if len(functionSelector) > 1 {
			fmt.Println("Error: provide a single selector")
			return
		}

		selector := [4]byte(functionSelector[0])
		facetAddress, err := loupe.FacetAddress(&bind.CallOpts{}, selector)
		if err != nil {
			fmt.Println("Error: getting facet address", err)
		}
		fmt.Println("facet address: ", facetAddress.String())

	case "help":
		printLoupeUsage()

	case "exit":
		fmt.Println("Exiting...")
		os.Exit(0)

	default:
		fmt.Printf("Unknown command: %s\n", s)
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
