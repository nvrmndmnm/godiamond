package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/c-bata/go-prompt"
	"github.com/spf13/pflag"
)

type Command struct {
	Name        string
	Description string
	SubCommands []*Command
}

func (c *Command) completer(d prompt.Document) []prompt.Suggest {
	args := strings.Split(d.Text, " ")

	var commands []*Command
	for _, cmd := range c.SubCommands {
		if cmd.Name == args[0] {
			commands = cmd.SubCommands
			break
		}
	}

	if commands == nil {
		commands = c.SubCommands
	}

	suggestions := make([]prompt.Suggest, 0, len(commands))

outer:
	for _, cmd := range commands {
		name := cmd.Name

		for i, arg := range args {
			if strings.Contains(arg, cmd.Name) {
				continue outer
			}

			if i > 0 {
				name = "--" + cmd.Name + "="
			}
		}

		suggestions = append(suggestions, prompt.Suggest{
			Text:        name,
			Description: cmd.Description,
		})
	}

	return suggestions
}

func (box *DiamondBox) executor(s string) {
	s = strings.TrimSpace(s)
	args := strings.Split(s, " ")
	var cmd *Command

	for _, c := range box.mode.SubCommands {
		if c.Name == args[0] {
			cmd = c
			break
		}
	}

	if cmd == nil {
		fmt.Printf("Unknown command: %s\n", args[0])
		return
	}

	switch cmd.Name {
	case "help":
		box.mode.printUsage()
		return
	case "exit":
		fmt.Println("Exiting...")
		os.Exit(0)
	}

	flags := pflag.NewFlagSet(cmd.Name, pflag.ContinueOnError)
	for _, subCmd := range cmd.SubCommands {
		flags.String(subCmd.Name, "", subCmd.Description)
	}

	err := flags.Parse(args[1:])
	if err != nil {
		fmt.Println("Invalid arguments for a command", err)
		return
	}

	switch box.mode.Name {
	case "deploy":
		box.modeDeploy(cmd, flags)
	case "cut":
		box.modeCut(cmd, flags)
	case "loupe":
		box.modeLoupe(cmd, flags)
	}
}

func (c *Command) printUsage() {
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

func (box *DiamondBox) run() {
	fmt.Println("Please enter a command. Type 'exit' to quit.")
	p := prompt.New(
		box.executor,
		box.mode.completer,
		prompt.OptionPrefix("> "),
		prompt.OptionTitle(box.mode.Name),
		prompt.OptionMaxSuggestion(uint16(len(box.mode.SubCommands))),
		prompt.OptionSuggestionBGColor(prompt.Black),
		prompt.OptionSuggestionTextColor(prompt.LightGray),
		prompt.OptionDescriptionBGColor(prompt.Black),
		prompt.OptionDescriptionTextColor(prompt.LightGray),
		prompt.OptionSelectedSuggestionBGColor(prompt.Black),
		prompt.OptionSelectedSuggestionTextColor(prompt.White),
		prompt.OptionSelectedDescriptionBGColor(prompt.Black),
		prompt.OptionSelectedDescriptionTextColor(prompt.White),
		prompt.OptionScrollbarBGColor(prompt.Black),
	)
	p.Run()
}
