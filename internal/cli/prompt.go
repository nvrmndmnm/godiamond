package cli

import (
	"strings"

	"github.com/c-bata/go-prompt"
)

type Command struct {
	Name        string
	Description string
	SubCommands []*Command
}

func (c *Command) Completer(d prompt.Document) []prompt.Suggest {
	args := strings.Split(d.TextBeforeCursor(), " ")
	lastArg := args[len(args)-1]

	var commands []*Command
	if len(args) > 1 {
		for _, cmd := range c.SubCommands {
			if cmd.Name == args[0] {
				commands = cmd.SubCommands
				break
			}
		}
	} else {
		commands = c.SubCommands
	}

	suggestions := make([]prompt.Suggest, 0, len(commands))
	for _, cmd := range commands {
		name := cmd.Name
		if len(args) > 1 {
			name = "--" + cmd.Name + "="
		}

		if strings.Contains(d.TextBeforeCursor(), name) {
			continue
		}

		if strings.HasPrefix(name, lastArg) {
			suggestions = append(suggestions, prompt.Suggest{
				Text:        name,
				Description: cmd.Description,
			})
		}
	}

	return suggestions
}

