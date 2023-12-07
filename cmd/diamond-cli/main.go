package main

import (
	"fmt"
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

func main() {
	var args Arguments

	pflag.Usage = PrintUsage
	pflag.StringVarP(&args.ValueConfig, "config", "c", "config.yaml", "Load config file")
	pflag.StringVarP(&args.ValueRPC, "rpc", "", "", "RPC identifier")
	pflag.Int64Var(&args.ValueChainID, "chain-id", 0, "Chain id.")
	pflag.BoolVarP(&args.FlagDebug, "debug", "d", false, "Enable debug mode")

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

	diamondCutFacet := k.String("contracts.cut_facet")
	if len(diamondCutFacet) == 0 {
		k.Set("contracts.cut_facet", common.Address{})
	}
	diamond := k.String("contracts.diamond")
	if len(diamond) == 0 {
		k.Set("contracts.diamond", common.Address{})
	}
	diamondInit := k.String("contracts.diamond_init")
	if len(diamondInit) == 0 {
		k.Set("contracts.diamond_init", common.Address{})
	}

	var config Config
	if err := k.Unmarshal("", &config); err != nil {
		sugar.Fatalf("Error: failed to unmarshal config: %v", err)
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

	box, err := NewDiamondBox(config, "local", 0)
	if err != nil {
		sugar.Error("Error: couldn't fill the box with treasures")
	}

	defer box.Close()

	switch os.Args[1] {
	case "deploy":
		err = box.deploy()

	case "cut":
		err = box.cut()

	case "loupe":
		err = box.loupe()

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
	fmt.Print(usage)
}
