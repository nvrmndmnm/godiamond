package main

import (
	"context"
	"encoding/json"
	"fmt"
	"math/big"
	"os"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"go.uber.org/zap"
)

type ContractMetadata struct {
	ABI      abi.ABI `json:"abi"`
	Bytecode struct {
		Object string `json:"object"`
	} `json:"bytecode"`
	MethodIdentifiers map[string]string `json:"methodIdentifiers"`
	AST               struct {
		Nodes []struct {
			Name string `json:"name"`
		} `json:"nodes"`
	} `json:"ast"`
}

type DiamondBox struct {
	config    Config
	sugar     *zap.SugaredLogger
	mode      Mode
	eth       *EthereumWrapper
	contracts map[string]ContractMetadata
}

func NewDiamondBox(config Config,
	sugar *zap.SugaredLogger,
	modeName string,
	rpcId string,
	chainId *big.Int,
) (*DiamondBox, error) {
	var err error

	box := &DiamondBox{
		config:    config,
		sugar:     sugar,
		contracts: make(map[string]ContractMetadata),
	}

	for contractIdentifier, contractConfig := range config.Contracts {
		contractMetadata, err := getContractMetadataByFile(contractConfig.MetadataFilePath)
		if err != nil {
			return nil, err
		}
		box.contracts[contractIdentifier] = contractMetadata
	}

	box.eth = &EthereumWrapper{}

	box.eth.client, err = box.eth.Dial(config.RPC[rpcId])
	if err != nil {
		return nil, err
	}

	if chainId.Cmp(big.NewInt(-1)) == 0 {
		box.eth.chainId, err = box.eth.NetworkID(context.Background())
		if err != nil {
			return nil, err
		}
	}

	privateKey, err := box.eth.HexToECDSA(config.Accounts["anvil"].PrivateKey[2:])
	if err != nil {
		return nil, err
	}

	box.eth.auth, err = box.eth.NewKeyedTransactorWithChainID(privateKey, box.eth.chainId)
	if err != nil {
		return nil, err
	}

	gasPrice, err := box.eth.SuggestGasPrice(context.Background())
	if err != nil {
		return nil, err
	}
	box.eth.auth.GasPrice = gasPrice

	factory := NewModeFactory(box)

	box.mode = factory.CreateMode(modeName)
	if box.mode == nil {
		printUsage()
		return nil, fmt.Errorf("mode does not exist")
	}

	return box, nil
}

func (box *DiamondBox) Close() {
	box.eth.Close()
}

func (box *DiamondBox) getContractMetadataByAddress(address common.Address) (*ContractMetadata, error) {
	for id, contract := range box.config.Contracts {
		if contract.Address == &address {
			contractMetadata := box.contracts[id]

			return &contractMetadata, nil
		}
	}

	return nil, fmt.Errorf("contract address not found in config")
}

func getContractMetadataByFile(path string) (ContractMetadata, error) {
	var contractMetadata ContractMetadata

	path = strings.Trim(path, "\"")
	metadataFile, err := os.ReadFile(path)
	if err != nil {
		return ContractMetadata{}, fmt.Errorf("failed to read metadata file: %v", err)
	}

	err = json.Unmarshal(metadataFile, &contractMetadata)
	if err != nil {
		return ContractMetadata{}, fmt.Errorf("failed to unmarshal metadata file: %v", err)
	}
	return contractMetadata, nil
}

func getFunctionIdentifiersBySelectors(selectors [][4]byte, contractMetadata *ContractMetadata) map[string]string {
	selectorsMetadata := make(map[string]string)

	for _, selector := range selectors {
		selectorString := hexutil.Encode(selector[:])
		functionName := "Function name not specified"

		if contractMetadata != nil {
			for identifier, selectorValue := range contractMetadata.MethodIdentifiers {
				if selectorString[2:] == selectorValue {
					functionName = identifier
					break
				}
			}
		}
		selectorsMetadata[selectorString] = functionName
	}

	return selectorsMetadata
}
