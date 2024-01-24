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
	RPC       string         `json:"rpc"`
	ChainID   big.Int        `json:"chainId"`
	TxHash    string         `json:"tx"`
}

func (box *DiamondBox) deployContractById(contractIdentifier string, params ...any) (*DeploymentData, error) {
	contractMetadata := box.contracts[contractIdentifier]

	address, tx, _, err := bind.DeployContract(box.eth.auth,
		contractMetadata.ABI,
		common.FromHex(contractMetadata.Bytecode.Object),
		box.eth.client, params...)
	if err != nil {
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
		Deployer:  box.eth.auth.From,
		Name:      contractMetadata.AST.Nodes[len(contractMetadata.AST.Nodes)-1].Name,
		Selectors: facetSelectors,
		ChainID:   *box.eth.chainId,
		TxHash:    tx.Hash().Hex(),
	}

	fmt.Printf("%s address: %s\ntx: %s", contractIdentifier, address.Hex(), tx.Hash().Hex()+"\n")

	return &deploymentData, nil
}

func writeDeploymentDataToFile(data *DeploymentData) error {
	jsonData, err := json.MarshalIndent(data, "", "    ")
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

	fileName := data.Name + "-" + time + ".json"
	path = filepath.Join(path, fileName)

	err = os.WriteFile(path, jsonData, 0644)
	if err != nil {
		return fmt.Errorf("failed to write deployment data: %v", err)
	}

	return nil
}
