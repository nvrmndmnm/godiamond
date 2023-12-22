package main

import (
	"fmt"
	"sort"
	"strings"

	"github.com/c-bata/go-prompt"
)

type Command struct {
	Description string
	SubCommands map[string]*Command
}

var mode = &Command{
	SubCommands: map[string]*Command{
		"comm1": {
			Description: "Command 1",
			SubCommands: map[string]*Command{
				"--sub1=": {Description: "Description1"},
			},
		},
		"comm2": {
			Description: "Command 2",
			SubCommands: map[string]*Command{
				"--subcmd1=": {Description: "Description1"},
				"--subcmd2=": {Description: "Description2"},
				"--subcmd3=": {Description: "Description3"},
			},
		},
	},
}

func completer(d prompt.Document) []prompt.Suggest {
	commands := make([]prompt.Suggest, len(mode.SubCommands))

	for commandName, command := range mode.SubCommands {
		commands = append(commands, prompt.Suggest{
			Text:        commandName,
			Description: command.Description,
		})
	}
	sort.Slice(commands, func(i, j int) bool {
		return commands[i].Text < commands[j].Text
	})

	args := strings.Split(d.Text, " ")

	if len(args) <= 1 {
		return prompt.FilterHasPrefix(commands, d.TextBeforeCursor(), true)
	}

	cmd := mode.SubCommands[args[0]]
	s := make([]prompt.Suggest, 0)

	for text, subcmd := range cmd.SubCommands {
		selected := false

		for _, arg := range args {
			if strings.Contains(arg, text) {
				selected = true
			}
		}

		if !selected {
			s = append(s, prompt.Suggest{Text: text, Description: subcmd.Description})
		}
	}

	sort.Slice(s, func(i, j int) bool {
		return s[i].Text < s[j].Text
	})

	return s
}

func executor(s string) {
	fmt.Print(s)
}
