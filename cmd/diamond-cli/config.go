package main

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common"
)

type EOA struct {
	PrivateKey string         `koanf:"private_key"`
	Address    common.Address `koanf:"address"`
}

type ContractConfig struct {
	Address          common.Address `koanf:"address"`
	MetadataFilePath string         `koanf:"metadata"`
}

type Config struct {
	Accounts  map[string]EOA              `koanf:"eoa"`
	RPC       map[string]string           `koanf:"rpc"`
	Contracts map[string]ContractConfig `koanf:"contracts"`
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
