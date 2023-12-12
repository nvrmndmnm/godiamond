package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/c-bata/go-prompt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/nvrmndmnm/godiamond/internal/facets"
	"github.com/nvrmndmnm/godiamond/internal/diamond"
)

type DeploymentData struct {
	Contracts map[common.Address]string
	TxHash    string
}

func printDeployUsage() {
	var usage = `
Commands:
    diamond <owner>         Deploy a new diamond
    facet                   Deploy a facet to use in an existing diamond
    init                    Deploy initial set of contracts specified by the standard 
    help                    Show help
    exit                    Exit the deploy mode

Arguments:
    --owner       string    Ethereum address of the diamond owner
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
		}
	}

	return []prompt.Suggest{}
}

func (box *DiamondBox) deployExecutor(s string) {
	s = strings.TrimSpace(s)
	args := strings.Split(s, " ")

	var tx *types.Transaction
	var err error
	var deploymentData DeploymentData
	contractsMap := make(map[common.Address]string, 0)

	switch args[0] {
	case "diamond":
		fmt.Println("diamond")
	case "facet":
		fmt.Println("facet")
	case "init":
		box.diamondCutFacet, tx, _, err = facets.DeployDiamondCutFacet(box.auth, box.client)
		if err != nil {
			fmt.Println("Error: deploy diamond cut facet")
			return
		}
		
		contractsMap[box.diamondCutFacet] = "DiamondCutFacet"
		log.Printf("DiamondCutFacet address: %s\ntx: %s",
			box.diamondCutFacet.Hex(), tx.Hash().Hex())

		owner := box.config.Accounts["anvil"].Address
		box.diamond, tx, _, err = diamond.DeployDiamond(box.auth, box.client, owner, box.diamondCutFacet)
		if err != nil {
			fmt.Println("Error: deploy diamond")
			return
		}

		contractsMap[box.diamond] = "Diamond"
		log.Printf("Diamond address: %s\ntx: %s",
			box.diamond.Hex(), tx.Hash().Hex())

		box.diamondInit, tx, _, err = facets.DeployDiamondInit(box.auth, box.client)
		if err != nil {
			fmt.Println("Error: deploy diamond init")
			return
		}

		contractsMap[box.diamondInit] = "DiamondInit"
		log.Printf("DiamondInit address: %s\ntx: %s",
			box.diamondInit.Hex(), tx.Hash().Hex())

		loupeAddress, tx, _, err := facets.DeployDiamondLoupeFacet(box.auth, box.client)
		if err != nil {
			fmt.Println("Error: deploy diamond loupe facet")
			return
		}
		box.facets = append(box.facets, loupeAddress)

		contractsMap[loupeAddress] = "DiamondLoupeFacet"
		log.Printf("DiamondLoupeFacet address: %s\ntx: %s",
			loupeAddress.Hex(), tx.Hash().Hex())

		ownershipAddress, tx, _, err := facets.DeployOwnershipFacet(box.auth, box.client)
		if err != nil {
			fmt.Println("Error: deploy ownership facet")
			return
		}
		box.facets = append(box.facets, ownershipAddress)
		
		contractsMap[ownershipAddress] = "OwnershipFacet"
		log.Printf("OwnershipFacet address: %s\ntx: %s",
			ownershipAddress.Hex(), tx.Hash().Hex())

		deploymentData = DeploymentData{
			Contracts: contractsMap,
		}

		writeDeploymentDataToFile(&deploymentData)

	case "help":
		printDeployUsage()

	case "exit":
		fmt.Println("Exiting...")
		os.Exit(0)

	default:
		fmt.Printf("Unknown command: %s\n", s)
	}
}

func writeDeploymentDataToFile(data *DeploymentData) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		log.Fatal("Error marshaling deployment data", err)
	}

	err = os.WriteFile("assets/data.json", jsonData, 0644)
	if err != nil {
		log.Fatal("Error writing deployment data", err)
	}
}