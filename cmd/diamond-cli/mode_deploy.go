package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/spf13/pflag"
)

type DeployMode struct {
	commands *Command
	box      *DiamondBox
}

func NewDeployMode(box *DiamondBox) Mode {
	commands := &Command{
		Name: "deploy",
		SubCommands: []*Command{
			{
				Name:        "init",
				Description: "Deploy and initialize mandatory Diamond standard contracts",
				SubCommands: []*Command{},
			},
			{
				Name:        "by-config-id",
				Description: "Deploy a contract by ID specified in config file",
				SubCommands: []*Command{
					{
						Name:        "id",
						Description: "ID of the contract in config file",
					},
					{
						Name:        "constructor-args",
						Description: "Comma-separated constructor arguments",
					},
				},
			},
			{
				Name:        "by-file",
				Description: "Deploy a contract by specified file path",
				SubCommands: []*Command{
					{
						Name:        "path",
						Description: "Path to the contract metadata file",
					},
					{
						Name:        "constructor-args",
						Description: "Comma-separated constructor arguments",
					},
				},
			},
		},
	}

	commands.SubCommands = append(commands.SubCommands, defaultCommands.SubCommands...)

	return &DeployMode{commands: commands, box: box}
}

func (d *DeployMode) GetCommands() *Command {
	return d.commands
}

func (d *DeployMode) PrintUsage() {
	PrintUsage(os.Stdout, d.commands)
}

func (d *DeployMode) Execute(cmd *Command, flags *pflag.FlagSet, modeParams ...interface{}) error {
	var deployments []*DeploymentData

	switch cmd.Name {
	case "init":
		cutFacet, err := d.box.deployContractById("cut_facet")
		if err != nil {
			return fmt.Errorf("failed to deploy the 'cut_facet' contract: %v", err)
		}
		deployments = append(deployments, cutFacet)

		owner := d.box.config.Accounts["anvil"].Address.Hex()
		diamond, err := d.box.deployContractById("diamond", owner, cutFacet.Address.Hex())
		if err != nil {
			return fmt.Errorf("failed to deploy the 'diamond' contract: %v", err)
		}
		deployments = append(deployments, diamond)

		diamondInit, err := d.box.deployContractById("diamond_init")
		if err != nil {
			return fmt.Errorf("failed to deploy the 'diamond_init' contract: %v", err)
		}
		deployments = append(deployments, diamondInit)

		loupeFacet, err := d.box.deployContractById("loupe_facet")
		if err != nil {
			return fmt.Errorf("failed to deploy the 'loupe_facet' contract: %v", err)
		}
		deployments = append(deployments, loupeFacet)

		if err = d.box.cutLoupeFacet(cutFacet.Address, loupeFacet.Address); err != nil {
			return fmt.Errorf("failed to cut loupe facet: %v", err)
		}

	case "by-config-id":
		contractIdentifier, err := flags.GetString("id")
		if err != nil {
			return fmt.Errorf("invalid identifier flag: %v", err)
		}

		if contractIdentifier == "" {
			return fmt.Errorf("identifier is required")
		}

		constructorArgsStr, err := flags.GetString("constructor-args")
		if err != nil {
			return fmt.Errorf("invalid constructor flag: %v", err)
		}

		argsList := strings.Split(constructorArgsStr, ",")

		deploymentData, err := d.box.deployContractById(contractIdentifier, argsList...)
		if err != nil {
			return fmt.Errorf("failed to deploy the contract: %v", err)
		}

		deployments = append(deployments, deploymentData)
	}

	if err := writeDeploymentDataToFile(deployments); err != nil {
		return fmt.Errorf("failed to write deployment data: %v", err)
	}
	return nil
}

func (box *DiamondBox) cutLoupeFacet(cutFacetAddress, loupeFacetAddress common.Address) error {
	cutContract := bind.NewBoundContract(box.config.Contracts["diamond"].Address,
		box.contracts["cut_facet"].ABI, box.eth.client, box.eth.client, box.eth.client)

	calldata, err := box.contracts["diamond_init"].ABI.Pack("init")
	if err != nil {
		return err
	}

	loupeMethodIdentifiers := box.contracts["loupe_facet"].MethodIdentifiers
	var loupeSelectors SelectorFlag

	for _, selector := range loupeMethodIdentifiers {
		if err := loupeSelectors.Set(selector); err != nil {
			return err
		}
	}

	var cut []FacetCut
	cut = append(cut, FacetCut{
		FacetAddress:      loupeFacetAddress,
		Action:            Add,
		FunctionSelectors: loupeSelectors,
	})

	_, err = cutContract.Transact(box.eth.auth, "diamondCut", cut,
		box.config.Contracts["diamond_init"].Address, calldata)
	if err != nil {
		return err
	}

	return nil
}
