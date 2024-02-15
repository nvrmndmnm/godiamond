package config

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/knadh/koanf"
	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/file"
)

type Config struct {
	PrivateKey string `koanf:"private_key"`
	RPC      map[string]string    `koanf:"rpc"`
	Metadata map[string]string `koanf:"metadata"`
	DiamondAddress common.Address `koanf:"diamond_address"`
}

func LoadConfig(path string) (Config, error) {
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

func (c *Config) ValidateStandardContracts() error {
	standardContracts := []string{"diamond", "diamond_init", "cut_facet", "loupe_facet"}
	for _, contract := range standardContracts {
		if _, ok := c.Metadata[contract]; !ok {
			return fmt.Errorf("missing mandatory ERC-2535 contract: %s", contract)
		}
	}
	return nil
}
