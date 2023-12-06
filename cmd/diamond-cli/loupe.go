package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/c-bata/go-prompt"
	"github.com/spf13/pflag"
)

func printLoupeUsage() {
	var usage = `
The loupe commands are:
    facets                        Show all facets and their selectors
    addresses                     Show all facet addresses used by a diamond
    facet-selectors <address>     Show all function selectors provided by a facet
    facet-address <selector>      Show the facet that supports the given selector
    help                          Show help
    exit                          Exit the loupe mode

The arguments are:
    --address           string    Ethereum address of a facet
    --selector          string    4-byte function selector representation 
`
	fmt.Print(usage)
}

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
	s = strings.TrimSpace(s)
	args := strings.Split(s, " ")

	switch args[0] {
	case "facets":
		fmt.Printf("facets result:\n")

	case "addresses":
		fmt.Printf("addresses result:\n")

	case "facet-selectors":
		var facetAddress AddressFlag
		var addressString string

		flags := pflag.NewFlagSet("facet-selectors", pflag.ContinueOnError)
		flags.StringVarP(&addressString, "address", "", "", "Facet address")
		err := flags.Parse(args[1:])

		if err != nil {
			fmt.Println("Error: invalid arguments for facet-selectors command")
			return
		}

		if err := facetAddress.Set(addressString); err != nil {
			fmt.Printf("Error: invalid Ethereum address format: %v\n", err)
			return
		}

		fmt.Printf("facet-selectors result: %v\n", facetAddress)

	case "facet-address":
		var functionSelector SelectorFlag
		var selectorString string

		flags := pflag.NewFlagSet("facet-address", pflag.ContinueOnError)
		flags.StringVarP(&selectorString, "selector", "", "", "Function selector")
		err := flags.Parse(args[1:])

		if err != nil {
			fmt.Println("Error: invalid arguments for facet-address command")
			return
		}

		if err := functionSelector.Set(selectorString); err != nil {
			fmt.Printf("Error: invalid selector format: %v\n", err)
		}

		fmt.Printf("facet-address result: %v\n", functionSelector)

	case "help":
		printLoupeUsage()

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
		prompt.OptionPrefix("> "),
		prompt.OptionTitle("loupe"),
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
	return nil
}
