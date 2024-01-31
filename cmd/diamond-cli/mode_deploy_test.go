package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewDeployMode(t *testing.T) {
	box, err := setupBox()
	if err != nil {
		t.Fatalf("Failed to create DiamondBox: %v", err)
	}

	mode := NewDeployMode(box)
	deployMode, ok := mode.(*DeployMode)
	if !ok {
		t.Fatalf("Expected type DeployMode, got %T", mode)
	}

	assert.NotNil(t, deployMode)
	assert.Equal(t, box, deployMode.box)
	assert.Equal(t, "deploy", deployMode.commands.Name)
}
func TestDeployMode_GetCommands(t *testing.T) {
	box, err := setupBox()
	if err != nil {
		t.Fatalf("Failed to create DiamondBox: %v", err)
	}

	mode := NewDeployMode(box)
	deployMode, ok := mode.(*DeployMode)
	if !ok {
		t.Fatalf("Expected type DeployMode, got %T", mode)
	}

	assert.Equal(t, deployMode.commands, deployMode.GetCommands())
}

func TestDeployMode_Execute(t *testing.T) {
	box, err := setupBox()
	if err != nil {
		t.Fatalf("Failed to create DiamondBox: %v", err)
	}

	mode := NewDeployMode(box)
	deployMode, ok := mode.(*DeployMode)
	if !ok {
		t.Fatalf("Expected type DeployMode, got %T", mode)
	}

	assert.NotNil(t, deployMode)
	assert.Equal(t, box, deployMode.box)
	assert.Equal(t, "deploy", deployMode.commands.Name)
}
