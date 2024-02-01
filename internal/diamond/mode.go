package diamond

import (
	"fmt"
	"io"

	"github.com/nvrmndmnm/godiamond/internal/cli"
	"github.com/spf13/pflag"
)

type ModeFactory struct {
	modes map[string]func() Mode
}

type Mode interface {
	Execute(cmd *cli.Command, flags *pflag.FlagSet, params ...interface{}) error
	GetCommands() *cli.Command
	PrintUsage()
}

var defaultCommands = &cli.Command{
	SubCommands: []*cli.Command{
		{
			Name:        "help",
			Description: "Show help message",
		},
		{
			Name:        "exit",
			Description: "Exit the interactive mode",
		},
	},
}

func NewModeFactory(box *DiamondBox) *ModeFactory {
	return &ModeFactory{
		modes: map[string]func() Mode{
			"deploy": func() Mode { return NewDeployMode(box) },
			"cut":    func() Mode { return NewCutMode(box) },
			"loupe":  func() Mode { return NewLoupeMode(box) },
		},
	}
}

func (f *ModeFactory) CreateMode(name string) Mode {
	if mode, ok := f.modes[name]; ok {
		return mode()
	}
	return nil
}

func PrintUsage(w io.Writer, c *cli.Command) {
	fmt.Fprintf(w, "\nCommands:\n")

	for _, cmd := range c.SubCommands {
		fmt.Fprintf(w, "    %-20s %s\n", cmd.Name, cmd.Description)
	}

	fmt.Fprintf(w, "\nArguments:\n")

	printedSubCommands := make(map[string]bool)

	for _, cmd := range c.SubCommands {
		if len(cmd.SubCommands) > 0 {
			for _, subCmd := range cmd.SubCommands {
				if _, ok := printedSubCommands[subCmd.Name]; !ok {
					fmt.Fprintf(w, "    %-20s %s\n", "--"+subCmd.Name+"=", subCmd.Description)
					printedSubCommands[subCmd.Name] = true
				}
			}
		}
	}
}
