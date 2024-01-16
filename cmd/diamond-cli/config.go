package main

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/knadh/koanf"
	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/file"
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
	Accounts  map[string]EOA            `koanf:"eoa"`
	RPC       map[string]string         `koanf:"rpc"`
	Contracts map[string]ContractConfig `koanf:"contracts"`
}

func loadConfig(path string) (Config, error) {
	var config Config

	k := koanf.New(".")
	if err := k.Load(file.Provider(path), yaml.Parser()); err != nil {
		return Config{}, err
	}

	if err := k.Unmarshal("", &config); err != nil {
		return Config{}, err
	}

	return config, nil
}

func (c *Config) validateStandardContracts() error {
	standardContracts := []string{"diamond", "diamond_init", "cut_facet", "loupe_facet"}
	for _, contract := range standardContracts {
		if _, ok := c.Contracts[contract]; !ok {
			return fmt.Errorf("missing mandatory ERC-2535 contract: %s", contract)
		}
	}
	return nil
}
