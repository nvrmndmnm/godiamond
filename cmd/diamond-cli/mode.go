package main

import (
	"fmt"

	"github.com/spf13/pflag"
)

type ModeFactory struct {
	modes map[string]func() Mode
}

type Mode interface {
	Execute(cmd *Command, flags *pflag.FlagSet)
	GetCommands() *Command
	PrintUsage()
}

var defaultCommands = &Command{
	SubCommands: []*Command{
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

func PrintUsage(c *Command) {
	fmt.Printf("\nCommands:\n")

	for _, cmd := range c.SubCommands {
		fmt.Printf("    %s\t\t%s\n", cmd.Name, cmd.Description)
	}

	fmt.Printf("\nArguments:\n")

	for _, cmd := range c.SubCommands {
		if len(cmd.SubCommands) > 0 {
			for _, subCmd := range cmd.SubCommands {
				//TODO: add line-length dependent spaces, ignore duplicates
				fmt.Printf("    %s\t\t%s\n", "--"+subCmd.Name+"=", subCmd.Description)
			}
		}
	}
}

