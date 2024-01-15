package main

import (
	"testing"

	"github.com/spf13/pflag"
	"github.com/stretchr/testify/assert"
)

func TestNewCutMode(t *testing.T) {
	box, err := setupBox()
	if err != nil {
		t.Fatalf("Failed to create DiamondBox: %v", err)
	}

	mode := NewCutMode(box)
	cutMode, ok := mode.(*CutMode)
	if !ok {
		t.Fatalf("Expected type CutMode, got %T", mode)
	}

	assert.NotNil(t, cutMode)
	assert.Equal(t, box, cutMode.box)
	assert.Equal(t, "cut", cutMode.commands.Name)
}

func TestCutMode_GetCommands(t *testing.T) {
	box, err := setupBox()
	if err != nil {
		t.Fatalf("Failed to create DiamondBox: %v", err)
	}

	mode := NewCutMode(box)
	cutMode, ok := mode.(*CutMode)
	if !ok {
		t.Fatalf("Expected type CutMode, got %T", mode)
	}

	assert.Equal(t, cutMode.commands, cutMode.GetCommands())
}

func TestCutMode_Execute(t *testing.T) {
	box, err := setupBox()
	if err != nil {
		t.Fatalf("Failed to create DiamondBox: %v", err)
	}

	mode := NewCutMode(box)
	cutMode, ok := mode.(*CutMode)
	if !ok {
		t.Fatalf("Expected type CutMode, got %T", mode)
	}

	cutMode.cutContract = &MockBoundContract{}

	cmd := &Command{Name: "add"}
	flags := pflag.NewFlagSet("test", pflag.ContinueOnError)
	flags.String("address", "0xCAFEBABECAFEBABECAFEBABECAFEBABECAFEBABE", "")
	flags.String("selectors", "0xdeadbeef", "")

	err = cutMode.Execute(cmd, flags)

	assert.Nil(t, err)
}
