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
	args := strings.Split(d.Text, " ")

	commands := mode.SubCommands
	if cmd, ok := commands[args[0]]; ok {
		commands = cmd.SubCommands
	}

	suggestions := make([]prompt.Suggest, 0, len(commands))

outer:
	for text, cmd := range commands {
		for _, arg := range args {
			if strings.Contains(arg, text) {
				continue outer
			}
		}

		suggestions = append(suggestions, prompt.Suggest{
			Text:        text,
			Description: cmd.Description,
		})
	}

	sort.Slice(suggestions, func(i, j int) bool {
		return suggestions[i].Text < suggestions[j].Text
	})

	return suggestions
}

func executor(s string) {
	fmt.Println(s)
}
