package main

import (
	"context"
	"encoding/json"
	"fmt"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
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

type DiamondBox struct {
	config Config
	client *ethclient.Client
	auth   *bind.TransactOpts

	contracts map[string]ContractMetadata
}

func NewDiamondBox(config Config, rpc string, chainId int64) (*DiamondBox, error) {
	var err error

	box := &DiamondBox{
		config:    config,
		contracts: make(map[string]ContractMetadata),
	}

	for contractIdentifier, contractMeta := range config.Contracts {
		var contractMetadata ContractMetadata
		
		metadataFile, err := os.ReadFile(contractMeta.MetadataFilePath)
		if err != nil {
			fmt.Printf("Error: Failed to read metadata file: %v\n", err)
			return nil, err
		}

		err = json.Unmarshal(metadataFile, &contractMetadata)
		if err != nil {
			fmt.Printf("Error: Failed to unmarshal metadata file: %v\n", err)
			return nil, err
		}

		box.contracts[contractIdentifier] = contractMetadata
	}

	box.client, err = ethclient.Dial(config.RPC[rpc])
	if err != nil {
		return nil, err
	}

	if chainId == 0 {
		networkId, err := box.client.NetworkID(context.Background())
		if err != nil {
			return nil, err
		}

		chainId = networkId.Int64()
	}

	privateKey, err := crypto.HexToECDSA(config.Accounts["anvil"].PrivateKey[2:])
	if err != nil {
		return nil, err
	}

	box.auth, err = bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(chainId))
	if err != nil {
		return nil, err
	}

	gasPrice, err := box.client.SuggestGasPrice(context.Background())
	if err != nil {
		return nil, err
	}
	box.auth.GasPrice = gasPrice

	return box, nil
}

func (box *DiamondBox) Close() {
	box.client.Close()
}
