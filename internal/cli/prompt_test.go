package cli

import (
	"reflect"
	"testing"

	"github.com/c-bata/go-prompt"
)

func TestCommandCompleter(t *testing.T) {
	cmd := &Command{
		Name: "test",
		SubCommands: []*Command{
			{
				Name:        "dead",
				Description: "No way",
				SubCommands: []*Command{
					{
						Name:        "babe",
						Description: "From the cafe",
					},
				},
			},
			{
				Name:        "deaf",
				Description: "Way?",
				SubCommands: []*Command{},
			},
			{
				Name:        "beef",
				Description: "Beefety beef",
				SubCommands: []*Command{
					{
						Name:        "metabeef",
						Description: "Metabeef",
					},
					{
						Name:        "overbeef",
						Description: "Overbeef",
					},
				},
			},
		},
	}

	testCases := []struct {
		input         string
		expected      []prompt.Suggest
		expectedError error
	}{
		{input: "de", expected: []prompt.Suggest{
			{Text: "dead", Description: "No way"},
			{Text: "deaf", Description: "Way?"},
		}},
		{input: "bee", expected: []prompt.Suggest{{Text: "beef", Description: "Beefety beef"}}},
		{input: "dead", expected: []prompt.Suggest{}},
		{input: "beef ", expected: []prompt.Suggest{
			{Text: "--metabeef=", Description: "Metabeef"},
			{Text: "--overbeef=", Description: "Overbeef"},
		}},
	}

	for _, tc := range testCases {
		buffer := prompt.NewBuffer()
		buffer.InsertText(tc.input, false, true)
		suggestions := cmd.Completer(*buffer.Document())

		if !reflect.DeepEqual(suggestions, tc.expected) {
			t.Errorf("Unexpected completions for input '%s': got %v, expected %v",
				tc.input, suggestions, tc.expected)
		}

	}
}
