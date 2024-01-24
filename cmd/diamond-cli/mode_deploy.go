package main

import (
	"fmt"
	"os"
	"strings"

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
				Name:        "facet",
				Description: "Deploy a facet to use in an existing diamond",
				SubCommands: []*Command{
					{
						Name:        "metadata",
						Description: "Path to the contract metadata file",
					},
					{
						Name:        "constructor-args",
						Description: "Comma-separated constructor arguments",
					},
				},
			},
			{
				Name:        "init",
				Description: "Deploy a facet to use in an existing diamond",
				SubCommands: []*Command{},
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

func (d *DeployMode) Execute(cmd *Command, flags *pflag.FlagSet, params ...interface{}) error {
	switch cmd.Name {
	case "facet":
		metadataFilePath, err := flags.GetString("metadata")
		if err != nil {
			return fmt.Errorf("invalid metadata flag: %v", err)
		}

		constructorArgsStr, err := flags.GetString("constructor-args")
		if err != nil {
			return fmt.Errorf("invalid constructor flag: %v", err)
		}

		argsList := strings.Split(constructorArgsStr, ",")

		constructorArgs := make([]interface{}, len(argsList))
		for i, arg := range argsList {
			constructorArgs[i] = arg
		}

		deploymentData, err := d.box.deployContract(metadataFilePath, constructorArgsStr)
		if err != nil {
			return fmt.Errorf("failed to deploy the contract: %v", err)
		}

		writeDeploymentDataToFile(deploymentData)

	case "init":
		cutFacet, err := d.box.deployContract("cut_facet")
		if err != nil {
			return fmt.Errorf("failed to deploy the 'cut_facet' contract: %v", err)
		}
		writeDeploymentDataToFile(cutFacet)

		owner := d.box.config.Accounts["anvil"].Address
		diamond, err := d.box.deployContract("diamond", owner, cutFacet.Address)
		if err != nil {
			return fmt.Errorf("failed to deploy the 'diamond' contract: %v", err)
		}
		writeDeploymentDataToFile(diamond)

		diamondInit, err := d.box.deployContract("diamond_init")
		if err != nil {
			return fmt.Errorf("failed to deploy the 'diamond_init' contract: %v", err)
		}
		writeDeploymentDataToFile(diamondInit)

		loupeFacet, err := d.box.deployContract("loupe_facet")
		if err != nil {
			return fmt.Errorf("failed to deploy the 'loupe_facet' contract: %v", err)
		}
		writeDeploymentDataToFile(loupeFacet)
	}

	return nil
}
