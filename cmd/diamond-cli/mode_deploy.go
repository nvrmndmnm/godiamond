package main

import (
	"fmt"
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
				Name:        "diamond",
				Description: "Deploy a new diamond",
				SubCommands: []*Command{
					{
						Name:        "owner",
						Description: "Specify the Ethereum address of the owner",
					},
				},
			},
			{
				Name:        "facet",
				Description: "Deploy a facet to use in an existing diamond",
				SubCommands: []*Command{
					{
						Name:        "metadata",
						Description: "Path to contract metadata file",
					},
					{
						Name:        "constructor-args",
						Description: "Comma-separated list of constructor arguments",
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
	PrintUsage(d.commands)
}

func (d *DeployMode) Execute(cmd *Command, flags *pflag.FlagSet) {
	switch cmd.Name {
	case "diamond":
		owner, err := flags.GetString("owner")
		if err != nil {
			fmt.Println("invalid owner flag")
			return
		}
		fmt.Println(owner)

	case "facet":
		metadataFilePath, err := flags.GetString("metadata")
		if err != nil {
			fmt.Println("invalid metadata flag")
			return
		}

		constructorArgsStr, err := flags.GetString("constructor-args")
		if err != nil {
			fmt.Println("invalid constructor flag")
			return
		}

		argsList := strings.Split(constructorArgsStr, ",")

		constructorArgs := make([]interface{}, len(argsList))
		for i, arg := range argsList {
			constructorArgs[i] = arg
		}

		deploymentData, err := d.box.deployContract(metadataFilePath, constructorArgsStr)
		if err != nil {
			fmt.Println("error deploying the contract:", err)
			return
		}

		writeDeploymentDataToFile(deploymentData)

	case "init":
		cutFacet, err := d.box.deployContract("cut_facet")
		if err != nil {
			fmt.Println("error deploying the cut_facet contract:", err)
			return
		}
		writeDeploymentDataToFile(cutFacet)

		owner := d.box.config.Accounts["anvil"].Address
		diamond, err := d.box.deployContract("diamond", owner, cutFacet.Address)
		if err != nil {
			fmt.Println("error deploying the contract:", err)
			return
		}
		writeDeploymentDataToFile(diamond)

		diamondInit, err := d.box.deployContract("diamond_init")
		if err != nil {
			fmt.Println("error deploying the contract:", err)
			return
		}
		writeDeploymentDataToFile(diamondInit)

		loupeFacet, err := d.box.deployContract("loupe_facet")
		if err != nil {
			fmt.Println("error deploying the contract:", err)
			return
		}
		writeDeploymentDataToFile(loupeFacet)

	}
}