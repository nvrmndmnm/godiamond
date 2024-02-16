package diamond

import (
	"fmt"
	"os"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/nvrmndmnm/godiamond/internal/cli"
	"github.com/nvrmndmnm/godiamond/internal/ethereum"
	"github.com/spf13/pflag"
)

type CutMode struct {
	commands    *cli.Command
	box         *DiamondBox
	cutContract ethereum.BoundContract
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
	commands := &cli.Command{
		Name: "cut",
		SubCommands: []*cli.Command{
			{
				Name:        "add",
				Description: "Add a new facet with specified function selectors",
				SubCommands: []*cli.Command{
					{
						Name:        "address",
						Description: "Ethereum address of a facet",
					},
					{
						Name:        "selectors",
						Description: "Comma-separated function selectors",
					},
				},
			},
			{
				Name:        "replace",
				Description: "Replace selectors of an existing facet",
				SubCommands: []*cli.Command{
					{
						Name:        "address",
						Description: "Ethereum address of a facet",
					},
					{
						Name:        "selectors",
						Description: "Comma-separated function selectors",
					},
				},
			},
			{
				Name:        "remove",
				Description: "Remove selectors from the diamond",
				SubCommands: []*cli.Command{
					{
						Name:        "selectors",
						Description: "Comma-separated function selectors",
					},
				},
			},
		},
	}

	commands.SubCommands = append(commands.SubCommands, defaultCommands.SubCommands...)

	diamondAddress := box.Config.DiamondAddress
	if diamondAddress == (common.Address{}) {
		box.Sugar.Fatal("diamond address is not set in the config")
	}

	cutContract := bind.NewBoundContract(box.Config.DiamondAddress, box.Contracts["cut_facet"].ABI,
		box.Eth.Client, box.Eth.Client, box.Eth.Client)

	return &CutMode{commands: commands, box: box, cutContract: cutContract}
}

func (c *CutMode) GetCommands() *cli.Command {
	return c.commands
}

func (c *CutMode) PrintUsage() {
	PrintUsage(os.Stdout, c.commands)
}

func (c *CutMode) Execute(cmd *cli.Command, flags *pflag.FlagSet, params ...interface{}) error {
	var cut []FacetCut
	var action uint8
	var facetAddress cli.AddressFlag
	var functionSelectors cli.SelectorFlag

	if cmd.Name == "add" || cmd.Name == "replace" {
		addressString, err := flags.GetString("address")
		if err != nil {
			return fmt.Errorf("invalid address flag: %v", err)
		}

		if err := facetAddress.Set(addressString); err != nil {
			return fmt.Errorf("invalid Ethereum address format: %v", err)
		}
	}

	selectorString, err := flags.GetString("selectors")
	if err != nil {
		return fmt.Errorf("invalid selector flag: %v", err)
	}

	if err := functionSelectors.Set(selectorString); err != nil {
		return fmt.Errorf("invalid selector format: %v", err)
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

	tx, err := c.cutContract.Transact(c.box.Eth.Auth, "diamondCut", cut,
		common.Address{}, []byte{})
	if err != nil {
		return fmt.Errorf("failed to cut diamond: %v", err)
	}

	fmt.Printf("Cut successful\ntx: %s\n", tx.Hash())

	return nil
}
