package main

import (
	"log"
	"os"

	"github.com/ethereum/go-ethereum/common"
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

type DiamondBox struct {
	DiamondCutFacet common.Address
	Diamond         common.Address
	DiamondInit     common.Address
	Facets          []common.Address
}

func main() {
	var args Arguments

	pflag.Usage = PrintUsage
	pflag.StringVarP(&args.ValueConfig, "config", "c", "config.yaml", "Load config file")
	pflag.StringVarP(&args.ValueRPC, "rpc", "", "", "RPC identifier")
	pflag.Int64Var(&args.ValueChainID, "chain-id", 0, "Chain id.")
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

	if len(os.Args) < 2 {
		sugar.Error("Error: no arguments provided")
		pflag.Usage()
	}

	if args.ValueRPC == "" {
		sugar.Error("Error: the rpc flag is required")
		pflag.Usage()
	}

	switch os.Args[1] {
	case "deploy":
		err = deploy(config, args.ValueRPC, args.ValueChainID)

	case "cut":
		err = cut()

	case "loupe":
		err = loupe()

	default:
		sugar.Error("Error: no command specified")
		pflag.Usage()
	}

	if err != nil {
		sugar.Fatal(err)
	}
}

func PrintUsage() {
	var usage = "diamond-cli" + `

Usage:
    diamond-cli deploy [options]
    diamond-cli cut [options]
    diamond-cli loupe [options]

Options:
    --rpc <name>          string    RPC identifier
    --chain-id <id>       int       Chain ID (default: 0)
    -c --config <path>    string    Load config file (default: "config.yaml")
    -d --debug                      Enable debug mode (default: disabled)
    -h --help                       Show help

`
	log.Print(usage)
}