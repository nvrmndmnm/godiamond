package main

import (
	"fmt"
	"strings"

	"github.com/spf13/pflag"
)

func (box *DiamondBox) modeDeploy(cmd *Command, flags *pflag.FlagSet) {
	fmt.Println(box.mode.SubCommands)
	fmt.Println("---------")
	fmt.Println(cmd)
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

		deploymentData, err := box.deployContract(metadataFilePath, constructorArgsStr)
		if err != nil {
			fmt.Println("error deploying the contract:", err)
			return
		}

		writeDeploymentDataToFile(deploymentData)

	case "init":
		cutFacet, err := box.deployContract("cut_facet")
		if err != nil {
			fmt.Println("error deploying the cut_facet contract:", err)
			return
		}
		writeDeploymentDataToFile(cutFacet)

		owner := box.config.Accounts["anvil"].Address
		diamond, err := box.deployContract("diamond", owner, cutFacet.Address)
		if err != nil {
			fmt.Println("error deploying the contract:", err)
			return
		}
		writeDeploymentDataToFile(diamond)

		diamondInit, err := box.deployContract("diamond_init")
		if err != nil {
			fmt.Println("error deploying the contract:", err)
			return
		}
		writeDeploymentDataToFile(diamondInit)

		loupeFacet, err := box.deployContract("loupe_facet")
		if err != nil {
			fmt.Println("error deploying the contract:", err)
			return
		}
		writeDeploymentDataToFile(loupeFacet)

	}
}
