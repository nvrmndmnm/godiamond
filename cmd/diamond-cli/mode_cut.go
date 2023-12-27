package main

import (
	"fmt"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/spf13/pflag"
)

type CutMode struct {
	commands *Command
	box      *DiamondBox
}

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

func NewCutMode(box *DiamondBox) Mode {
	commands := &Command{
		Name: "cut",
		SubCommands: []*Command{
			{
				Name:        "add",
				Description: "Add a new facet with specified function selectors",
				SubCommands: []*Command{
					{
						Name:        "address",
						Description: "Specify the Ethereum address of a facet",
					},
					{
						Name:        "selectors",
						Description: "Specify the function selectors",
					},
				},
			},
			{
				Name:        "replace",
				Description: "Replace selectors of an existing facet",
				SubCommands: []*Command{
					{
						Name:        "address",
						Description: "Specify the Ethereum address of a facet",
					},
					{
						Name:        "selectors",
						Description: "Specify the function selectors",
					},
				},
			},
			{
				Name:        "remove",
				Description: "Remove selectors from the diamond",
				SubCommands: []*Command{
					{
						Name:        "selectors",
						Description: "Specify the function selectors",
					},
				},
			},
		},
	}

	commands.SubCommands = append(commands.SubCommands, defaultCommands.SubCommands...)

	return &CutMode{commands: commands, box: box}
}

func (c *CutMode) GetCommands() *Command {
	return c.commands
}

func (c *CutMode) PrintUsage() {
	PrintUsage(c.commands)
}

func (c *CutMode) Execute(cmd *Command, flags *pflag.FlagSet) {
	diamondCut := bind.NewBoundContract(c.box.config.Contracts["diamond"].Address,
		c.box.contracts["cut_facet"].ABI, c.box.client, c.box.client, c.box.client)

	calldata, err := c.box.contracts["diamond_init"].ABI.Pack("init")
	if err != nil {
		fmt.Println(err)
	}

	var cut []FacetCut

	switch cmd.Name {
	case "add", "replace", "remove":
		var action uint8
		var facetAddress AddressFlag
		var functionSelectors SelectorFlag

		if cmd.Name == "add" || cmd.Name == "replace" {
			addressString, err := flags.GetString("address")
			if err != nil {
				fmt.Println("invalid address flag")
				return
			}

			if err := facetAddress.Set(addressString); err != nil {
				fmt.Printf("invalid Ethereum address format: %v\n", err)
				return
			}
		}

		selectorString, err := flags.GetString("selectors")
		if err != nil {
			fmt.Println("invalid selector flag")
			return
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

		tx, err := diamondCut.Transact(c.box.auth, "diamondCut", cut,
			c.box.config.Contracts["diamond_init"].Address, calldata)
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println(tx.Hash())
	}
}
