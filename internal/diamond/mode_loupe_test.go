package diamond

import (
	"bytes"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/nvrmndmnm/godiamond/internal/cli"
	"github.com/spf13/pflag"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestNewLoupeMode(t *testing.T) {
	box, err := SetupBox()
	if err != nil {
		t.Fatalf("Failed to create DiamondBox: %v", err)
	}

	mode := NewLoupeMode(box)
	loupeMode, ok := mode.(*LoupeMode)
	if !ok {
		t.Fatalf("Expected type LoupeMode, got %T", mode)
	}

	assert.NotNil(t, loupeMode)
	assert.Equal(t, box, loupeMode.box)
	assert.Equal(t, "loupe", loupeMode.commands.Name)
}
func TestLoupeMode_GetCommands(t *testing.T) {
	box, err := SetupBox()
	if err != nil {
		t.Fatalf("Failed to create DiamondBox: %v", err)
	}

	mode := NewLoupeMode(box)
	loupeMode, ok := mode.(*LoupeMode)
	if !ok {
		t.Fatalf("Expected type LoupeMode, got %T", mode)
	}

	assert.Equal(t, loupeMode.commands, loupeMode.GetCommands())
}

func TestLoupeMode_Execute(t *testing.T) {
	box, err := SetupBox()
	if err != nil {
		t.Fatalf("Failed to create DiamondBox: %v", err)
	}

	mode := NewLoupeMode(box)
	loupeMode, ok := mode.(*LoupeMode)
	if !ok {
		t.Fatalf("Expected type LoupeMode, got %T", mode)
	}

	mockContract := SetupMockLoupeContract()

	loupeMode.loupeContract = mockContract
	cmd := &cli.Command{Name: "facets"}
	flags := pflag.NewFlagSet("test", pflag.ContinueOnError)

	buf := new(bytes.Buffer)
	err = loupeMode.Execute(cmd, flags, buf)
	assert.NoError(t, err)
	mockContract.AssertCalled(t, "Call", &bind.CallOpts{}, mock.Anything, cmd.Name)

	expectedOutput := "facet address: 0xFEeDBAbefEedbaBefeEDbAbEfEeDBABeFEEDbaBe\n\t0xbc645d96: test((address,uint8)address,bytes)\n\n"
	assert.Equal(t, expectedOutput, buf.String(), "The output does not match the expected text")

}
