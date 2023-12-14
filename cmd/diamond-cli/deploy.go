package main

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

type ContractMetadata struct {
	ABI      abi.ABI `json:"abi"`
	Bytecode struct {
		Object string `json:"object"`
	} `json:"bytecode"`
	MethodIdentifiers SelectorsMetadata `json:"methodIdentifiers"`
	AST               struct {
		Nodes []struct {
			Name string `json:"name"`
		} `json:"nodes"`
	} `json:"ast"`
}

type DeploymentData struct {
	Address   common.Address `json:"address"`
	Name      string         `json:"name"`
	Selectors [][]string     `json:"selectors"`
	TxHash    string         `json:"tx"`
}

func (box *DiamondBox) deployContract(contractMetadata *ContractMetadata) (*DeploymentData, error) {
	address, tx, _, err := bind.DeployContract(box.auth,
		contractMetadata.ABI,
		common.FromHex(contractMetadata.Bytecode.Object),
		box.client)
	if err != nil {
		fmt.Println("Error deploying contract:", err)
		return nil, err
	}

	facetSelectors := make([][]string, 0, len(contractMetadata.MethodIdentifiers))

	functionNames := make([]string, 0, len(contractMetadata.MethodIdentifiers))
	for name := range contractMetadata.MethodIdentifiers {
		functionNames = append(functionNames, name)
	}
	sort.Strings(functionNames)

	for _, name := range functionNames {
		facetSelectors = append(facetSelectors, []string{name, contractMetadata.MethodIdentifiers[name]})
	}

	deploymentData := DeploymentData{
		Address:   address,
		Name:      contractMetadata.AST.Nodes[1].Name,
		Selectors: facetSelectors,
		TxHash:    tx.Hash().Hex(),
	}

	fmt.Printf("Facet address: %s\ntx: %s", address.Hex(), tx.Hash().Hex())

	return &deploymentData, nil
}

func writeDeploymentDataToFile(data *DeploymentData) {
	jsonData, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		fmt.Println("Error marshaling deployment data", err)
	}

	fileName := "assets/" + data.Name + ".json"
	err = os.WriteFile(fileName, jsonData, 0644)
	if err != nil {
		fmt.Println("Error writing deployment data", err)
	}
}
