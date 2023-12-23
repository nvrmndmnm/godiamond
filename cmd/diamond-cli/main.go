package main

import (
	"fmt"
	"log"
	"math/big"
	"os"

	"github.com/c-bata/go-prompt"
	"github.com/knadh/koanf"
	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/file"
	"github.com/spf13/pflag"
	"go.uber.org/zap"
)

type Arguments struct {
	ValueRPC     string
	ValueConfig  string
	ValueChainID int64
	FlagDebug    bool
}

func printUsage() {
	var usage = "diamond-cli" + `

Usage:
    diamond-cli deploy [options]
    diamond-cli cut [options]
    diamond-cli loupe [options]
	diamond-cli help

Options:
    --rpc <name>          string    RPC identifier
    --chain-id <id>       int       Chain ID (default: -1, will auto-detect)
    -c --config <path>    string    Load config file (default: "config.yaml")
    -d --debug                      Enable debug mode (default: disabled)

`
	fmt.Print(usage)
}

func runMode(mode string, exectutor prompt.Executor, completer prompt.Completer) {
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

func main() {
	var args Arguments

	pflag.Usage = printUsage
	pflag.StringVarP(&args.ValueConfig, "config", "c", "config.yaml", "Load config file")
	pflag.StringVar(&args.ValueRPC, "rpc", "", "RPC identifier")
	pflag.Int64Var(&args.ValueChainID, "chain-id", -1, "Chain id.")
	pflag.BoolVarP(&args.FlagDebug, "debug", "d", false, "Enable debug mode")

	pflag.Parse()

	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatalf("Error: failed to initialize logger: %v", err)
	}

	if args.FlagDebug {
		logger, err = zap.NewDevelopment()
		if err != nil {
			log.Fatalf("Error: failed to initialize debug logger: %v", err)
		}
	}

	sugar := logger.Sugar()
	defer logger.Sync()

	k := koanf.New(".")
	if err := k.Load(file.Provider(args.ValueConfig), yaml.Parser()); err != nil {
		sugar.Fatalf("Error: failed to load config: %v", err)
	}

	var config Config

	if err := k.Unmarshal("", &config); err != nil {
		sugar.Fatalf("Error: failed to unmarshal config: %v", err)
	}

	err = config.validateStandardContracts()
	if err != nil {
		sugar.Fatalf("Error: failed to validate config: %v", err)
	}

	if len(os.Args) < 2 {
		sugar.Error("Error: no arguments provided")
		pflag.Usage()
	}

	//TODO: decide if rpc is needed to be an argument
	// if args.ValueRPC == "" {
	// 	sugar.Error("Error: the rpc flag is required")
	// 	pflag.Usage()
	// }

	chainId := new(big.Int)
	chainId.SetInt64(args.ValueChainID)

	box, err := NewDiamondBox(config, "local", chainId)
	if err != nil {
		sugar.Error("Error: couldn't fill the box with treasures")
	}

	defer box.Close()

	defaultMode := &Command{
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

	switch os.Args[1] {
	case "deploy":
		mode := &Command{
			SubCommands: []*Command{
				{
					Name:        "diamond",
					Description: "Deploy a new diamond",
					SubCommands: []*Command{
						{
							Name:        "--owner=",
							Description: "Specify the Ethereum address of the owner",
						},
					},
				},
				{
					Name:        "facet",
					Description: "Deploy a facet to use in an existing diamond",
					SubCommands: []*Command{
						{
							Name:        "--metadata=",
							Description: "Path to contract metadata file",
						},
						{
							Name:        "--constructor-args=",
							Description: "Specify contract constructor arguments",
						},
					},
				},
				{
					Name:        "init",
					Description: "Deploy a facet to use in an existing diamond",
					SubCommands: []*Command{},
				},
			},
		}

		mode.SubCommands = append(mode.SubCommands, defaultMode.SubCommands...)

		runMode("deploy", box.deployExecutor, mode.completer)

	case "cut":
		mode := &Command{
			SubCommands: []*Command{
				{
					Name:        "add",
					Description: "Add a new facet with specified function selectors",
					SubCommands: []*Command{
						{
							Name:        "--address=",
							Description: "Specify the Ethereum address of a facet",
						},
						{
							Name:        "--selectors=",
							Description: "Specify the function selectors",
						},
					},
				},
				{
					Name:        "replace",
					Description: "Replace selectors of an existing facet",
					SubCommands: []*Command{
						{
							Name:        "--address=",
							Description: "Specify the Ethereum address of a facet",
						},
						{
							Name:        "--selectors=",
							Description: "Specify the function selectors",
						},
					},
				},
				{
					Name:        "remove",
					Description: "Remove selectors from the diamond",
					SubCommands: []*Command{
						{
							Name:        "--selectors=",
							Description: "Specify the function selectors",
						},
					},
				},
			},
		}

		runMode("cut", box.cutExecutor, mode.completer)

	case "loupe":
		mode := &Command{
			SubCommands: []*Command{
				{
					Name:        "facets",
					Description: "Show all facets and their selectors",
				},
				{
					Name:        "addresses",
					Description: "Show all facet addresses used by a diamond",
				},
				{
					Name:        "facet-selectors",
					Description: "Show all function selectors provided by a facet",
					SubCommands: []*Command{
						{
							Name:        "--address=",
							Description: "Specify the Ethereum address of a facet",
						},
					},
				},
				{
					Name:        "facet-address",
					Description: "Show the facet that supports the given selector",
					SubCommands: []*Command{
						{
							Name:        "--selector=",
							Description: "Specify the function selector",
						},
					},
				},
				{
					Name:        "supports-interface",
					Description: "Show if the contract implements an interface",
					SubCommands: []*Command{
						{
							Name:        "--id=",
							Description: "Specify the interface identifier",
						},
					},
				},
			},
		}

		runMode("loupe", box.loupeExecutor, mode.completer)

	case "test":
		mode := &Command{
			SubCommands: []*Command{
				{
					Name:        "diamond",
					Description: "Deploy a new diamond",
					SubCommands: []*Command{
						{
							Name:        "--owner=",
							Description: "Specify the Ethereum address of the owner",
						},
					},
				},
			},
		}

		runMode("test", executor, mode.completer)

	case "help":
		pflag.Usage()

	default:
		sugar.Error("Error: no command specified")
		pflag.Usage()
	}
}
