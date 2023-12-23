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

func executor(s string) {
	fmt.Println(s)
}
