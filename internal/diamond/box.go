package diamond

import (
	"context"
	"encoding/json"
	"fmt"
	"math/big"
	"os"
	"strings"

	"github.com/c-bata/go-prompt"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/nvrmndmnm/godiamond/internal/cli"
	"github.com/nvrmndmnm/godiamond/internal/config"
	"github.com/nvrmndmnm/godiamond/internal/ethereum"
	"github.com/spf13/pflag"
	"go.uber.org/zap"
)

type ContractMetadata struct {
	ABI      abi.ABI `json:"abi"`
	Bytecode struct {
		Object string `json:"object"`
	} `json:"bytecode"`
	MethodIdentifiers map[string]string `json:"methodIdentifiers"`
	AST               struct {
		Nodes []struct {
			Name string `json:"name"`
		} `json:"nodes"`
	} `json:"ast"`
}

type DiamondBox struct {
	Config    config.Config
	Sugar     *zap.SugaredLogger
	Mode      Mode
	Eth       *ethereum.EthereumWrapper
	Contracts map[string]ContractMetadata
}

func NewDiamondBox(config config.Config,
	sugar *zap.SugaredLogger,
	modeName string,
	rpcId string,
	chainId *big.Int,
) (*DiamondBox, error) {
	var err error

	box := &DiamondBox{
		Config:    config,
		Sugar:     sugar,
		Contracts: make(map[string]ContractMetadata),
	}

	for contractIdentifier, contractConfig := range config.Contracts {
		contractMetadata, err := GetContractMetadataByFile(contractConfig.MetadataFilePath)
		if err != nil {
			return nil, err
		}
		box.Contracts[contractIdentifier] = contractMetadata
	}

	box.Eth = &ethereum.EthereumWrapper{}

	box.Eth.Client, err = box.Eth.Dial(config.RPC[rpcId])
	if err != nil {
		return nil, err
	}

	if chainId.Cmp(big.NewInt(-1)) == 0 {
		box.Eth.ChainId, err = box.Eth.NetworkID(context.Background())
		if err != nil {
			return nil, err
		}
	}

	privateKey, err := box.Eth.HexToECDSA(config.Accounts["anvil"].PrivateKey[2:])
	if err != nil {
		return nil, err
	}

	box.Eth.Auth, err = box.Eth.NewKeyedTransactorWithChainID(privateKey, box.Eth.ChainId)
	if err != nil {
		return nil, err
	}

	gasPrice, err := box.Eth.SuggestGasPrice(context.Background())
	if err != nil {
		return nil, err
	}
	box.Eth.Auth.GasPrice = gasPrice

	factory := NewModeFactory(box)

	box.Mode = factory.CreateMode(modeName)
	if box.Mode == nil {
		pflag.Usage()
		return nil, fmt.Errorf("mode does not exist")
	}

	return box, nil
}

func (box *DiamondBox) executor(s string) {
	s = strings.TrimSpace(s)
	args := strings.Split(s, " ")
	var cmd *cli.Command

	for _, c := range box.Mode.GetCommands().SubCommands {
		if c.Name == args[0] {
			cmd = c
			break
		}
	}

	if cmd == nil {
		box.Sugar.Errorf("unknown command: %s\n", args[0])
		return
	}

	switch cmd.Name {
	case "help":
		box.Mode.PrintUsage()
		return
	case "exit":
		os.Exit(0)
	}

	flags := pflag.NewFlagSet(cmd.Name, pflag.ContinueOnError)
	for _, subCmd := range cmd.SubCommands {
		flags.String(subCmd.Name, "", subCmd.Description)
	}

	err := flags.Parse(args[1:])
	if err != nil {
		box.Sugar.Errorf("invalid arguments for a command: %v\n", err)
		return
	}

	err = box.Mode.Execute(cmd, flags)
	if err != nil {
		box.Sugar.Errorf("mode execution error: %v\n", err)
		return
	}
}

func (box *DiamondBox) Run() {
	fmt.Println("Please enter a command. Type 'exit' to quit.")
	p := prompt.New(
		box.executor,
		box.Mode.GetCommands().Completer,
		prompt.OptionPrefix("> "),
		prompt.OptionTitle(box.Mode.GetCommands().Name),
		prompt.OptionMaxSuggestion(uint16(len(box.Mode.GetCommands().SubCommands))),
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

func (box *DiamondBox) Close() {
	box.Eth.Close()
}

func (box *DiamondBox) getContractMetadataByAddress(address common.Address) (*ContractMetadata, error) {
	for id, contract := range box.Config.Contracts {
		if common.HexToAddress(contract.Address) == address {
			contractMetadata := box.Contracts[id]

			return &contractMetadata, nil
		}
	}

	return nil, fmt.Errorf("contract address not found in config")
}

func GetContractMetadataByFile(path string) (ContractMetadata, error) {
	var contractMetadata ContractMetadata

	path = strings.Trim(path, "\"")
	metadataFile, err := os.ReadFile(path)
	if err != nil {
		return ContractMetadata{}, fmt.Errorf("failed to read metadata file: %v", err)
	}

	err = json.Unmarshal(metadataFile, &contractMetadata)
	if err != nil {
		return ContractMetadata{}, fmt.Errorf("failed to unmarshal metadata file: %v", err)
	}
	return contractMetadata, nil
}

func getFunctionIdentifiersBySelectors(selectors [][4]byte, contractMetadata *ContractMetadata) map[string]string {
	selectorsMetadata := make(map[string]string)

	for _, selector := range selectors {
		selectorString := hexutil.Encode(selector[:])
		functionName := "Function name not specified"

		if contractMetadata != nil {
			for identifier, selectorValue := range contractMetadata.MethodIdentifiers {
				if selectorString[2:] == selectorValue {
					functionName = identifier
					break
				}
			}
		}
		selectorsMetadata[selectorString] = functionName
	}

	return selectorsMetadata
}
