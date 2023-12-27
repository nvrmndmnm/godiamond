package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/pflag"
)

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
			fmt.Println("error deploying the contract:", err)
			return
		}

		writeDeploymentDataToFile(deploymentData)

	case "init":
		//TODO: Add more flexible way of bulk deployments
		deploymentData, err := box.deployContract("cut_facet")
		if err != nil {
			fmt.Println("error deploying the cut_facet contract:", err)
			return
		}
		writeDeploymentDataToFile(deploymentData)

		owner := box.config.Accounts["anvil"].Address
		deploymentData, err = box.deployContract("diamond", owner, deploymentData.Address)
		if err != nil {
			fmt.Println("error deploying the contract:", err)
			return
		}
		writeDeploymentDataToFile(deploymentData)

		deploymentData, err = box.deployContract("diamond_init")
		if err != nil {
			fmt.Println("error deploying the contract:", err)
			return
		}
		writeDeploymentDataToFile(deploymentData)

		deploymentData, err = box.deployContract("loupe_facet")
		if err != nil {
			fmt.Println("error deploying the contract:", err)
			return
		}
		writeDeploymentDataToFile(deploymentData)

	case "help":
		box.mode.printUsage()

	case "exit":
		fmt.Println("Exiting...")
		os.Exit(0)

	default:
		fmt.Printf("Unknown command: %s\n", s)
	}
}
