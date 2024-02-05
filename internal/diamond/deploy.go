package diamond

import (
	"encoding/json"
	"fmt"
	"math/big"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/nvrmndmnm/godiamond/internal/cli"
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

	address, tx, _, err := bind.DeployContract(box.Eth.Auth,
		contractMetadata.ABI,
		common.FromHex(contractMetadata.Bytecode.Object),
		box.Eth.Client, params...)
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
		Deployer:  box.Eth.Auth.From,
		Name:      contractMetadata.AST.Nodes[len(contractMetadata.AST.Nodes)-1].Name,
		Selectors: facetSelectors,
		ChainID:   *box.Eth.ChainId,
		TxHash:    tx.Hash().Hex(),
	}

	fmt.Printf("%s address: %s\ntx: %s", deploymentData.Name, address.Hex(), tx.Hash().Hex()+"\n")

	return deploymentData, nil
}

func (box *DiamondBox) initCutLoupeFacet(diamondAddress, diamondInitAddress,
	cutFacetAddress, loupeFacetAddress common.Address) error {
	cutContract := bind.NewBoundContract(diamondAddress, box.Contracts["cut_facet"].ABI,
		box.Eth.Client, box.Eth.Client, box.Eth.Client)

	calldata, err := box.Contracts["diamond_init"].ABI.Pack("init")
	if err != nil {
		return err
	}

	loupeMethodIdentifiers := box.Contracts["loupe_facet"].MethodIdentifiers
	var loupeSelectors cli.SelectorFlag

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

	_, err = cutContract.Transact(box.Eth.Auth, "diamondCut", cut, diamondInitAddress, calldata)
	if err != nil {
		return err
	}

	return nil
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

func convertStringParamsToType(strParams []string, types abi.Arguments) ([]interface{}, error) {
	params := make([]interface{}, len(strParams))
	var err error

	for i, value := range strParams {
		switch types[i].Type.T {
		case abi.AddressTy:
			if !common.IsHexAddress(value) {
				return nil, fmt.Errorf("%s is not a valid Ethereum address", value)
			}
			params[i] = common.HexToAddress(value)

		case abi.BoolTy:
			params[i], err = strconv.ParseBool(value)
			if err != nil {
				return nil, err
			}

		case abi.BytesTy:
			params[i] = []byte(value)

		case abi.IntTy, abi.UintTy:
			res, ok := new(big.Int).SetString(value, 0)
			if !ok {
				return nil, fmt.Errorf("failed to convert to big.Int: %s", value)
			}
			params[i] = res

		case abi.StringTy:
			params[i] = value
		}
	}

	return params, nil
}
