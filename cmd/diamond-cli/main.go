package main

import (
	"fmt"
	"log"
	"math/big"
	"os"

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
    --rpc <name>          string    RPC identifier (required)
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
		log.Fatalf("failed to initialize logger: %v", err)
	}

	if args.FlagDebug {
		logger, err = zap.NewDevelopment()
		if err != nil {
			log.Fatalf("failed to initialize debug logger: %v", err)
		}
	}

	sugar := logger.Sugar()
	defer logger.Sync()

	config, err := loadConfig(args.ValueConfig)
	if err != nil {
		sugar.Fatalf("failed to load config: %v", err)
	}

	if err := config.validateStandardContracts(); err != nil {
		sugar.Fatalf("failed to validate config: %v", err)
	}

	if len(os.Args) < 2 {
		sugar.Error("no arguments provided")
		pflag.Usage()
	}

	if os.Args[1] == "help" {
		printUsage()
		return
	}

	mode := os.Args[1]

	if args.ValueRPC == "" {
		sugar.Error("the RPC flag is required")
		pflag.Usage()
	}

	_, ok := config.RPC[args.ValueRPC]
	if !ok {
		sugar.Fatal("provided RPC identifier is not in the config")
	}

	chainId := new(big.Int)
	chainId.SetInt64(args.ValueChainID)

	box, err := NewDiamondBox(config, sugar, mode, args.ValueRPC, chainId)
	if err != nil {
		sugar.Fatalf("couldn't fill the box with treasures: %v", err)
	}

	defer box.Close()

	box.run()
}
