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
	config    Config
	mode      Mode
	client    *ethclient.Client
	auth      *bind.TransactOpts
	rpcName   string
	chainId   *big.Int
	contracts map[string]ContractMetadata
}

func NewDiamondBox(config Config, modeName string, rpc string, chainId *big.Int) (*DiamondBox, error) {
	var err error

	box := &DiamondBox{
		config:    config,
		rpcName:   rpc,
		chainId:   chainId,
		contracts: make(map[string]ContractMetadata),
	}

	for contractIdentifier, contractMeta := range config.Contracts {
		var contractMetadata ContractMetadata

		metadataFile, err := os.ReadFile(contractMeta.MetadataFilePath)
		if err != nil {
			fmt.Printf("failed to read metadata file: %v\n", err)
			return nil, err
		}

		err = json.Unmarshal(metadataFile, &contractMetadata)
		if err != nil {
			fmt.Printf("failed to unmarshal metadata file: %v\n", err)
			return nil, err
		}

		box.contracts[contractIdentifier] = contractMetadata
	}

	box.client, err = ethclient.Dial(config.RPC[box.rpcName])
	if err != nil {
		return nil, err
	}

	if chainId.Cmp(big.NewInt(-1)) == 0 {
		box.chainId, err = box.client.NetworkID(context.Background())
		if err != nil {
			return nil, err
		}
	}

	privateKey, err := crypto.HexToECDSA(config.Accounts["anvil"].PrivateKey[2:])
	if err != nil {
		return nil, err
	}

	box.auth, err = bind.NewKeyedTransactorWithChainID(privateKey, box.chainId)
	if err != nil {
		return nil, err
	}

	gasPrice, err := box.client.SuggestGasPrice(context.Background())
	if err != nil {
		return nil, err
	}
	box.auth.GasPrice = gasPrice

	factory := NewModeFactory(box)

	box.mode = factory.CreateMode(modeName)
	if box.mode == nil {
		printUsage()
		return nil, fmt.Errorf("mode does not exist")
	}

	return box, nil
}

func (box *DiamondBox) Close() {
	box.client.Close()
}
