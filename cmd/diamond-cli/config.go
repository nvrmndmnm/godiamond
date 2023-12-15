package main

import (
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
	Contracts struct {
		Diamond           ContractMetadata `koanf:"diamond"`
		DiamondInit       ContractMetadata `koanf:"diamond_init"`
		DiamondCutFacet   ContractMetadata `koanf:"cut_facet"`
		DiamondLoupeFacet ContractMetadata `koanf:"loupe_facet"`
	} `koanf:"contracts"`
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
