package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/c-bata/go-prompt"
	"github.com/spf13/pflag"
)

func printDeployUsage() {
	var usage = `
Commands:
    diamond <owner>                        Deploy a new diamond
    facet <metadata> <constructor-args>    Deploy a facet to use in an existing diamond
    init                                   Deploy initial set of contracts specified by the standard 
    help                                   Show help
    exit                                   Exit the deploy mode

Arguments:
    --owner              string    Ethereum address of the diamond owner
	--metadata           string    Path to contract metadata file
	--constructor-args   string    Comma-separated list of constructor arguments
`
	fmt.Print(usage)
}

func deployCompleter(d prompt.Document) []prompt.Suggest {
	s := []prompt.Suggest{
		{Text: "diamond", Description: "Deploy a new diamond"},
		{Text: "facet", Description: "Deploy a facet to use in an existing diamond"},
		{Text: "init", Description: "Deploy initial set of contracts specified by the standard"},
		{Text: "help", Description: "Show help message"},
		{Text: "exit", Description: "Exit the deploy mode"},
	}

	args := strings.Split(d.Text, " ")

	if len(args) <= 1 {
		return prompt.FilterHasPrefix(s, d.GetWordBeforeCursor(), true)
	}

	if len(args) == 2 {
		switch args[0] {
		case "diamond":
			return []prompt.Suggest{
				{Text: "--owner=", Description: "Specify the Ethereum address of the owner"},
			}
		case "facet":
			return []prompt.Suggest{
				{Text: "--metadata=", Description: "Path to contract metadata file"},
			}
		}
	}

	if len(args) == 3 {
		switch args[0] {
		case "facet":
			return []prompt.Suggest{
				{Text: "--constructor-args=", Description: "Specify contract construcor arguments"},
			}
		}
	}

	return []prompt.Suggest{}
}

func (box *DiamondBox) deployExecutor(s string) {
	s = strings.TrimSpace(s)
	args := strings.Split(s, " ")

	switch args[0] {
	case "diamond":
		fmt.Println("diamond")
	case "facet":
		var metadataFilePath, constructorArgsStr string

		flags := pflag.NewFlagSet("deploy-facet", pflag.ContinueOnError)
		flags.StringVarP(&metadataFilePath, "metadata", "", "", "Path to facet metadata file")
		flags.StringVarP(&constructorArgsStr, "constructor-args", "", "", "Constructor arguments")
		err := flags.Parse(args[1:])

		if err != nil {
			fmt.Println("Invalid arguments for facet deploy command")
			return
		}

		argsList := strings.Split(constructorArgsStr, ",")

		constructorArgs := make([]interface{}, len(argsList))
		for i, arg := range argsList {
			constructorArgs[i] = arg
		}

		deploymentData, err := box.deployContract(metadataFilePath, constructorArgsStr)
		if err != nil {
			fmt.Println("Error deploying the contract:", err)
			return
		}

		writeDeploymentDataToFile(deploymentData)

	case "init":
		//TODO: Add more flexible way of bulk deployments
		diamondCutFacetMetadataFilePath := box.config.Contracts.DiamondCutFacet.MetadataFilePath
		diamondMetadataFilePath := box.config.Contracts.Diamond.MetadataFilePath
		diamondInitMetadataFilePath := box.config.Contracts.DiamondInit.MetadataFilePath
		loupeFacetMetadataFilePath := box.config.Contracts.DiamondLoupeFacet.MetadataFilePath

		deploymentData, err := box.deployContract(diamondCutFacetMetadataFilePath)
		if err != nil {
			fmt.Println("Error deploying the contract:", err)
			return
		}
		writeDeploymentDataToFile(deploymentData)

		owner := box.config.Accounts["anvil"].Address
		deploymentData, err = box.deployContract(diamondMetadataFilePath, owner, deploymentData.Address)
		if err != nil {
			fmt.Println("Error deploying the contract:", err)
			return
		}
		writeDeploymentDataToFile(deploymentData)

		deploymentData, err = box.deployContract(diamondInitMetadataFilePath)
		if err != nil {
			fmt.Println("Error deploying the contract:", err)
			return
		}
		writeDeploymentDataToFile(deploymentData)

		deploymentData, err = box.deployContract(loupeFacetMetadataFilePath)
		if err != nil {
			fmt.Println("Error deploying the contract:", err)
			return
		}
		writeDeploymentDataToFile(deploymentData)

	case "help":
		printDeployUsage()

	case "exit":
		fmt.Println("Exiting...")
		os.Exit(0)

	default:
		fmt.Printf("Unknown command: %s\n", s)
	}
}
