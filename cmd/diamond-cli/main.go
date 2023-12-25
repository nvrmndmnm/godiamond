package main

import (
	"fmt"
	"log"
	"math/big"
	"os"

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

	mode := selectMode(os.Args[1])

	if mode == nil {
		pflag.Usage()
		sugar.Fatalf("Error: command not found")
	}

	box, err := NewDiamondBox(config, mode, "local", chainId)
	if err != nil {
		sugar.Error("Error: couldn't fill the box with treasures")
	}

	defer box.Close()

	switch box.mode.Name {
	case "deploy":
		runCommand(box.mode.Name, box.deployExecutor, box.mode.completer)
	
	case "cut":
		runCommand(box.mode.Name, box.cutExecutor, box.mode.completer)
		
	case "loupe":
		runCommand(box.mode.Name, box.loupeExecutor, box.mode.completer)
	
	case "test":
		runCommand(box.mode.Name, box.deployExecutor, box.mode.completer)

	case "help":
		box.mode.printUsage()

	case "exit":
		fmt.Println("Exiting...")
		os.Exit(0)

	default:
		fmt.Printf("Unknown command: %s\n", box.mode.Name)
	}
}
