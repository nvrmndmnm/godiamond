package main

import (
	"fmt"
	"io"
	"os"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/spf13/pflag"
)

type LoupeMode struct {
	commands      *Command
	box           *DiamondBox
	loupeContract BoundContract
}

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
						Description: "Ethereum address of a facet",
					},
				},
			},
			{
				Name:        "facet-address",
				Description: "Show the facet that supports the given selector",
				SubCommands: []*Command{
					{
						Name:        "selector",
						Description: "Function selector",
					},
				},
			},
			{
				Name:        "supports-interface",
				Description: "Show if the contract implements an interface",
				SubCommands: []*Command{
					{
						Name:        "id",
						Description: "Interface identifier",
					},
				},
			},
		},
	}

	commands.SubCommands = append(commands.SubCommands, defaultCommands.SubCommands...)

	loupeContract := bind.NewBoundContract(box.config.Contracts["diamond"].Address,
		box.contracts["loupe_facet"].ABI, box.eth.client, box.eth.client, box.eth.client)

	return &LoupeMode{commands: commands, box: box, loupeContract: loupeContract}
}

func (l *LoupeMode) GetCommands() *Command {
	return l.commands
}

func (l *LoupeMode) PrintUsage() {
	PrintUsage(os.Stdout, l.commands)
}

func (l *LoupeMode) Execute(cmd *Command, flags *pflag.FlagSet, params ...interface{}) error {
	var output string
	var err error

	switch cmd.Name {
	case "facets":
		output, err = l.getFacetsOutput()
		if err != nil {
			return err
		}

	case "addresses":
		output, err = l.getFacetAddressesOutput()
		if err != nil {
			return err
		}

	case "facet-selectors":
		var facetAddress AddressFlag
		addressString, err := flags.GetString("address")
		if err != nil {
			return fmt.Errorf("invalid address flag: %v", err)
		}
		if err := facetAddress.Set(addressString); err != nil {
			return fmt.Errorf("invalid Ethereum address format: %v", err)
		}

		output, err = l.getFacetSelectorsOutput(facetAddress)
		if err != nil {
			return err
		}

	case "facet-address":
		var functionSelector SelectorFlag
		selectorString, err := flags.GetString("selector")
		if err != nil {
			return fmt.Errorf("invalid selector flag: %v", err)
		}

		if err := functionSelector.Set(selectorString); err != nil {
			return fmt.Errorf("invalid selector format: %v", err)
		}

		output, err = l.getFacetAddressOutput(functionSelector)
		if err != nil {
			return err
		}

	case "supports-interface":
		var interfaceId SelectorFlag
		interfaceIdString, err := flags.GetString("id")
		if err != nil {
			return fmt.Errorf("invalid id flag: %v", err)
		}

		if err := interfaceId.Set(interfaceIdString); err != nil {
			return fmt.Errorf("invalid id format: %v", err)
		}

		output, err = l.getSupportsInterfaceOutput(interfaceId)
		if err != nil {
			return err
		}
	}

	var out io.Writer = os.Stdout
	if len(params) > 0 {
		if writer, ok := params[0].(io.Writer); ok {
			out = writer
		}
	}

	fmt.Fprintln(out, output)

	return nil
}

func (l *LoupeMode) getFacetsOutput() (string, error) {
	var output string
	var callResult []interface{}

	err := l.loupeContract.Call(&bind.CallOpts{}, &callResult, "facets")
	if err != nil {
		return "", fmt.Errorf("failed to get facets of a diamond: %v", err)
	}

	facets := *abi.ConvertType(callResult[0], new([]LoupeFacet)).(*[]LoupeFacet)

	for _, facet := range facets {
		contractMetadata, _ := l.box.getContractMetadataByAddress(common.Address(facet.FacetAddress))
		selectorsMetadata := getFunctionIdentifiersBySelectors(facet.FunctionSelectors, contractMetadata)

		output += fmt.Sprintf("facet address: %v\n", facet.FacetAddress)

		for selector, functionName := range selectorsMetadata {
			output += fmt.Sprintf("\t%s: %s\n", selector, functionName)
		}
	}
	return output, nil
}

func (l *LoupeMode) getFacetAddressesOutput() (string, error) {
	var output string
	var callResult []interface{}

	err := l.loupeContract.Call(&bind.CallOpts{}, &callResult, "facetAddresses")
	if err != nil {
		return "", fmt.Errorf("failed to get addresses of the facets: %v", err)
	}

	facetAddresses := *abi.ConvertType(callResult[0], new([]common.Address)).(*[]common.Address)

	for _, address := range facetAddresses {
		output += fmt.Sprintf("%v\n", address)
	}
	return output, nil
}

func (l *LoupeMode) getFacetSelectorsOutput(facetAddress AddressFlag) (string, error) {
	var output string
	var callResult []interface{}

	err := l.loupeContract.Call(&bind.CallOpts{}, &callResult, "facetFunctionSelectors", common.Address(facetAddress))
	if err != nil {
		return "", fmt.Errorf("failed to get facet selectors: %v", err)
	}

	facetSelectors := *abi.ConvertType(callResult[0], new([][4]byte)).(*[][4]byte)

	contractMetadata, _ := l.box.getContractMetadataByAddress(common.Address(facetAddress))
	selectorsMetadata := getFunctionIdentifiersBySelectors(facetSelectors, contractMetadata)

	for selector, functionName := range selectorsMetadata {
		output += fmt.Sprintf("\t%s: %s\n", selector, functionName)
	}
	return output, nil
}

func (l *LoupeMode) getFacetAddressOutput(functionSelector SelectorFlag) (string, error) {
	var output string
	var callResult []interface{}

	if len(functionSelector) > 1 {
		return "", fmt.Errorf("a single selector is required")
	}

	selector := [4]byte(functionSelector[0])

	err := l.loupeContract.Call(&bind.CallOpts{}, &callResult, "facetAddress", selector)
	if err != nil {
		return "", fmt.Errorf("failed to get facet address: %v", err)
	}

	facetAddress := *abi.ConvertType(callResult[0], new(common.Address)).(*common.Address)

	output = fmt.Sprintf("%s\n", facetAddress.String())
	return output, nil
}

func (l *LoupeMode) getSupportsInterfaceOutput(interfaceId SelectorFlag) (string, error) {
	var output string
	var callResult []interface{}

	if len(interfaceId) > 1 {
		return "", fmt.Errorf("a single identifier is required")
	}

	id := [4]byte(interfaceId[0])

	err := l.loupeContract.Call(&bind.CallOpts{}, &callResult, "supportsInterface", id)
	if err != nil {
		return "", fmt.Errorf("failed to check interface support: %v", err)
	}

	status := *abi.ConvertType(callResult[0], new(bool)).(*bool)

	output = fmt.Sprintf("ERC-165 status: %v\n", status)
	return output, nil
}
