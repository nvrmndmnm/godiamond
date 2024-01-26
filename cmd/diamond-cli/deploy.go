package main

import (
	"encoding/json"
	"fmt"
	"math/big"
	"os"
	"path/filepath"
	"sort"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

type DeploymentData struct {
	Address   common.Address `json:"address"`
	Deployer  common.Address `json:"deployer"`
	Name      string         `json:"name"`
	Selectors [][]string     `json:"selectors"`
	ChainID   big.Int        `json:"chainId"`
	TxHash    string         `json:"tx"`
}

func (box *DiamondBox) deployContract(contractMetadata ContractMetadata, strParams ...string) (DeploymentData, error) {
	params, err := convertStringParamsToType(strParams, contractMetadata.ABI.Constructor.Inputs)
	if err != nil {
		return DeploymentData{}, err
	}

	address, tx, _, err := bind.DeployContract(box.eth.auth,
		contractMetadata.ABI,
		common.FromHex(contractMetadata.Bytecode.Object),
		box.eth.client, params...)
	if err != nil {
		return DeploymentData{}, err
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
		Deployer:  box.eth.auth.From,
		Name:      contractMetadata.AST.Nodes[len(contractMetadata.AST.Nodes)-1].Name,
		Selectors: facetSelectors,
		ChainID:   *box.eth.chainId,
		TxHash:    tx.Hash().Hex(),
	}

	fmt.Printf("%s address: %s\ntx: %s", deploymentData.Name, address.Hex(), tx.Hash().Hex()+"\n")

	return deploymentData, nil
}

func writeDeploymentDataToFile(data []DeploymentData) error {
	for _, entry := range data {
		jsonData, err := json.MarshalIndent(entry, "", "    ")
		if err != nil {
			return fmt.Errorf("failed to marshal deployment data: %v", err)
		}

		date := time.Now().Format("2006-01-02")
		time := time.Now().Format("15-04-05")

		wd, err := os.Getwd()
		if err != nil {
			return fmt.Errorf("failed to get working directory: %v", err)
		}

		path := filepath.Join(wd, "out/deployments", date)

		err = os.MkdirAll(path, 0755)
		if err != nil {
			return fmt.Errorf("failed to create a directory: %v", err)
		}

		fileName := entry.Name + "-" + time + ".json"
		path = filepath.Join(path, fileName)

		err = os.WriteFile(path, jsonData, 0644)
		if err != nil {
			return fmt.Errorf("failed to write file: %v", err)
		}
	}
	return nil
}
