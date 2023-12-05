package main

import (
	"fmt"
	"github.com/c-bata/go-prompt"
	"os"
	"strings"
)

func completer(d prompt.Document) []prompt.Suggest {
	s := []prompt.Suggest{
		{Text: "facets", Description: "Show all facets and their selectors"},
		{Text: "addresses", Description: "Show all facet addresses used by a diamond"},
		{Text: "facet-selectors", Description: "Show all function selectors provided by a facet"},
		{Text: "facet-address", Description: "Show the facet that supports the given selector"},
		{Text: "help", Description: "Show help message"},
		{Text: "exit", Description: "Exit the loupe interactive mode"},
	}

	args := strings.Split(d.Text, " ")
	if len(args) > 1 {
		switch args[0] {
		case "facet-selectors":
			s = []prompt.Suggest{
				{Text: "--address=", Description: "Specify the Ethereum address of a facet"},
			}
		case "facet-address":
			s = []prompt.Suggest{
				{Text: "--selector=", Description: "Specify the function selector"},
			}
		}
	}

	return prompt.FilterHasPrefix(s, d.GetWordBeforeCursor(), true)
}

func executor(s string) {
	switch s {
	case "facets":
		fmt.Printf("facets result:\n")
	case "addresses":
		fmt.Printf("addresses result:\n")
	case "facet-selectors":
		fmt.Printf("facet-selectors result:\n")
	case "facet-address":
		fmt.Printf("facet-address result:\n")
	case "help":
		PrintLoupeUsage()
	case "exit":
		fmt.Println("Exiting...")
		os.Exit(0)
	default:
		fmt.Printf("Unknown command: %s\n", s)
	}
}

func loupe() error {
	fmt.Println("Please enter a command. Type 'exit' to quit.")
	p := prompt.New(
		executor,
		completer,
		prompt.OptionPrefix(">>> "),
		prompt.OptionTitle("loupe"),
		prompt.OptionMaxSuggestion(1),
	)
	p.Run()
	return nil
}
func PrintLoupeUsage() {
	var usage = "loupe mode" + `

The loupe commands are:
    facets                        Show all facets and their selectors
    addresses                     Show all facet addresses used by a diamond
	facet-selectors <--address>   Show all function selectors provided by a facet
    facet-address <--selector>    Show the facet that supports the given selector
    help                          Show help
    exit                          Exit the loupe mode

The arguments are:
    <address>           string    Ethereum address of a facet
    <selector>          string    4-byte function selector representation 
`
	fmt.Print(usage)
}
