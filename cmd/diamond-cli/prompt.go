package main

import (
	"fmt"
	"strings"

	"github.com/c-bata/go-prompt"
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
		for _, arg := range args {
			if strings.Contains(arg, cmd.Name) {
				continue outer
			}
		}

		suggestions = append(suggestions, prompt.Suggest{
			Text:        cmd.Name,
			Description: cmd.Description,
		})
	}

	return suggestions
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
				fmt.Printf("    %s\t\t%s\n", subCmd.Name, subCmd.Description)
			}
		}
	}
}

func runCommand(mode string, exectutor prompt.Executor, completer prompt.Completer) {
	fmt.Println("Please enter a command. Type 'exit' to quit.")
	p := prompt.New(
		exectutor,
		completer,
		prompt.OptionPrefix("> "),
		prompt.OptionTitle(mode),
		prompt.OptionMaxSuggestion(4),
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
