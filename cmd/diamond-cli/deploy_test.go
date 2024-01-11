package main

import (
	"math/big"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type BoxMock struct {
	mock.Mock
}

func (b *BoxMock) deployContract(s string, a1, a2 common.Address) ([]byte, error) {
	args := b.Called(s, a1, a2)
	return args.Get(0).([]byte), args.Error(1)
}

func TestDeployContract(t *testing.T) {
	box := new(BoxMock)

	box.On("deployContract", "test", mock.AnythingOfType("common.Address"), mock.AnythingOfType("common.Address")).Return([]byte("data"), nil)

	data, err := box.deployContract("test",
		common.HexToAddress("0xf39fd6e51aad88f6f4ce6ab8827279cfffb92266"),
		common.HexToAddress("0xCAFEBABECAFEBABECAFEBABECAFEBABECAFEBABE"))

	assert.Nil(t, err)
	assert.NotNil(t, data)
	box.AssertExpectations(t)
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

	wd, err := os.Getwd()
	if err != nil {
		t.Fatalf("Failed to get working directory: %v", err)
	}

	outDir := filepath.Join(wd, "../../out/deployments")
	filePath := filepath.Join(outDir, time.Now().Format("2006-01-02"), data.Name+"-"+time.Now().Format("15-04-05")+".json")

	t.Log(filePath)
	_, err = os.Stat(filePath)
	assert.False(t, os.IsNotExist(err))

	err = os.Remove(filePath)
	assert.Nil(t, err)

	_, err = os.Stat(filePath)
	assert.True(t, os.IsNotExist(err))
}
