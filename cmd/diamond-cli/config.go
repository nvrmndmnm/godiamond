package main

import (
	"github.com/ethereum/go-ethereum/common"
)

type EOA struct {
	PrivateKey string         `koanf:"private_key"`
	Address    common.Address `koanf:"address"`
}

type Config struct {
	Accounts map[string]EOA `koanf:"eoa"`
	RPC map[string]string `koanf:"rpc"`
	Contracts struct {
		DiamondCutFacet common.Address `koanf:"cut_facet"`
		Diamond common.Address `koanf:"diamond"`
		DiamondInit common.Address `koanf:"diamond_init"`
	} `koanf:"contracts"`
}
