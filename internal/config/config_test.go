package config

import (
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
)

func TestLoadConfig(t *testing.T) {
	config, err := LoadConfig("../../testdata/config_test.yaml")
	assert.NoError(t, err, "Failed to load config")

	assert.Len(t, config.RPC, 1, "Invalid number of RPC entries")
	assert.Equal(t, "http://localhost:6969", config.RPC["test"], "Invalid RPC url")
	assert.Len(t, config.Metadata, 5, "Invalid number of metadata entries")
	assert.Equal(t, common.HexToAddress("0xABADBABEABADBABEABADBABEABADBABEABADBABE"),
		config.DiamondAddress, "Invalid address for diamond contract")
	assert.Equal(t, "../../testdata/TestDiamond.json",
		config.Metadata["diamond"], "Invalid metadata file path for diamond contract")
}

func TestValidateStandardContracts(t *testing.T) {
	config, err := LoadConfig("../../testdata/config_test.yaml")
	assert.NoError(t, err, "Failed to load config")

	err = config.ValidateStandardContracts()
	assert.NoError(t, err, "Failed to validate config")

	delete(config.Metadata, "diamond")
	err = config.ValidateStandardContracts()
	assert.Error(t, err)
}
