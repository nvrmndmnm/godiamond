package main

import (
	"fmt"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
)

type EOA struct {
	PrivateKey string         `koanf:"private_key"`
	Address    common.Address `koanf:"address"`
}

type Config struct {
	Accounts  map[string]EOA    `koanf:"eoa"`
	RPC       map[string]string `koanf:"rpc"`
	Contracts map[string]ContractMetadata `koanf:"contracts"`
}

type ContractMetadata struct {
	Address          common.Address `koanf:"address"`
	MetadataFilePath string         `koanf:"metadata"`
	ABI              abi.ABI        `json:"abi"`
	Bytecode         struct {
		Object string `json:"object"`
	} `json:"bytecode"`
	MethodIdentifiers SelectorsMetadata `json:"methodIdentifiers"`
	AST               struct {
		Nodes []struct {
			Name string `json:"name"`
		} `json:"nodes"`
	} `json:"ast"`
}


func (c *Config) validateStandardContracts() error {
	standardContracts := []string{"diamond", "diamond_init", "cut_facet", "loupe_facet"}
	for _, contract := range standardContracts {
		if _, ok := c.Contracts[contract]; !ok {
			fmt.Println(c.Contracts)
			return fmt.Errorf("missing mandatory ERC-2535 contract: %s", contract)
		}
	}
	return nil
}