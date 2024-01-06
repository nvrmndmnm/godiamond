package main

import (
	"math/big"
	"os"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
)

func TestDeployContract(t *testing.T) {
	box, err := setupBox()
	if err != nil {
		t.Fatalf("Failed to create DiamondBox: %v", err)
	}

	data, err := box.deployContract("test")

	assert.Nil(t, err)
	assert.NotNil(t, data)
}

func TestWriteDeploymentDataToFile(t *testing.T) {
	data := &DeploymentData{
		Address:   common.HexToAddress("0x123"),
		Deployer:  common.HexToAddress("0x123"),
		Name:      "test",
		Selectors: [][]string{{"test", "test"}},
		ChainID:   *big.NewInt(1),
		TxHash:    "test",
	}

	err := writeDeploymentDataToFile(data)

	assert.Nil(t, err)

	filePath := "out/deployments/" + time.Now().Format("2006-01-02") + "/" + data.Name + "-" + time.Now().Format("15-04-05") + ".json"
	_, err = os.Stat(filePath)
	assert.False(t, os.IsNotExist(err))

	err = os.Remove(filePath)
	assert.Nil(t, err)

	_, err = os.Stat(filePath)
	assert.True(t, os.IsNotExist(err))
}