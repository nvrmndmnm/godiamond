package config

import (
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
)

func TestLoadConfig(t *testing.T) {
	config, err := LoadConfig("../../testdata/config_test.yaml")
	assert.NoError(t, err, "Failed to load config")

	assert.Len(t, config.Accounts, 1, "Invalid number of EOA accounts")
	assert.Equal(t, "0xcafebabecafebabecafebabecafebabecafebabecafebabecafebabecafebabe",
		config.Accounts["anvil"].PrivateKey, "Invalid private key for anvil")
	assert.Equal(t, common.HexToAddress("0xCAFEBABECAFEBABECAFEBABECAFEBABECAFEBABE"),
		config.Accounts["anvil"].Address, "Invalid address for anvil")
	assert.Len(t, config.RPC, 1, "Invalid number of RPC urls")
	assert.Equal(t, "http://localhost:6969", config.RPC["test"], "Invalid RPC url")
	assert.Len(t, config.Contracts, 5, "Invalid number of contracts")
	assert.Equal(t, common.HexToAddress("0xABADBABEABADBABEABADBABEABADBABEABADBABE"),
		common.HexToAddress(config.Contracts["diamond"].Address), "Invalid address for diamond contract")
	assert.Equal(t, "../../testdata/TestDiamond.json",
		config.Contracts["diamond"].MetadataFilePath, "Invalid metadata file path for diamond contract")
}

func TestValidateStandardContracts(t *testing.T) {
	config, err := LoadConfig("../../testdata/config_test.yaml")
	assert.NoError(t, err, "Failed to load config")

	err = config.ValidateStandardContracts()
	assert.NoError(t, err, "Failed to validate config")

	delete(config.Contracts, "diamond")
	err = config.ValidateStandardContracts()
	assert.Error(t, err)
}
