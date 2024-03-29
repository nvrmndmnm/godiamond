package diamond

import (
	"fmt"
	"os"
	"strings"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/nvrmndmnm/godiamond/internal/cli"
	"github.com/spf13/pflag"
)

type DeployMode struct {
	commands *cli.Command
	box      *DiamondBox
}

func NewDeployMode(box *DiamondBox) Mode {
	commands := &cli.Command{
		Name: "deploy",
		SubCommands: []*cli.Command{
			{
				Name:        "init",
				Description: "Deploy and initialize mandatory Diamond standard contracts",
				SubCommands: []*cli.Command{},
			},
			{
				Name:        "by-config-id",
				Description: "Deploy a contract by ID specified in config file",
				SubCommands: []*cli.Command{
					{
						Name:        "id",
						Description: "ID of the contract in config file",
					},
					{
						Name:        "constructor-args",
						Description: "Comma-separated constructor arguments",
					},
				},
			},
			{
				Name:        "by-file",
				Description: "Deploy a contract by specified metadata file path",
				SubCommands: []*cli.Command{
					{
						Name:        "path",
						Description: "Path to the contract metadata file",
					},
					{
						Name:        "constructor-args",
						Description: "Comma-separated constructor arguments",
					},
				},
			},
		},
	}

	commands.SubCommands = append(commands.SubCommands, defaultCommands.SubCommands...)

	return &DeployMode{commands: commands, box: box}
}

func (d *DeployMode) GetCommands() *cli.Command {
	return d.commands
}

func (d *DeployMode) PrintUsage() {
	PrintUsage(os.Stdout, d.commands)
}

func (d *DeployMode) Execute(cmd *cli.Command, flags *pflag.FlagSet, modeParams ...interface{}) error {
	var deployments []DeploymentData

	switch cmd.Name {
	case "init":
		cutFacet, err := d.box.deployContract(d.box.Contracts["cut_facet"])
		if err != nil {
			return fmt.Errorf("failed to deploy the 'cut_facet' contract: %v", err)
		}
		deployments = append(deployments, cutFacet)

		privateKey, err := d.box.Eth.HexToECDSA(d.box.Config.PrivateKey[2:])
		if err != nil {
			return fmt.Errorf("failed to convert private key string: %v", err)
		}
		
		owner := crypto.PubkeyToAddress(privateKey.PublicKey).Hex()
		diamond, err := d.box.deployContract(d.box.Contracts["diamond"], owner, cutFacet.Address.Hex())
		if err != nil {
			return fmt.Errorf("failed to deploy the 'diamond' contract: %v", err)
		}
		deployments = append(deployments, diamond)

		diamondInit, err := d.box.deployContract(d.box.Contracts["diamond_init"])
		if err != nil {
			return fmt.Errorf("failed to deploy the 'diamond_init' contract: %v", err)
		}
		deployments = append(deployments, diamondInit)

		loupeFacet, err := d.box.deployContract(d.box.Contracts["loupe_facet"])
		if err != nil {
			return fmt.Errorf("failed to deploy the 'loupe_facet' contract: %v", err)
		}
		deployments = append(deployments, loupeFacet)

		if err = d.box.initCutLoupeFacet(diamond.Address, loupeFacet.Address); err != nil {
			return fmt.Errorf("failed to cut loupe facet: %v", err)
		}

	case "by-config-id":
		contractIdentifier, err := flags.GetString("id")
		if err != nil {
			return fmt.Errorf("invalid identifier flag: %v", err)
		}

		if contractIdentifier == "" {
			return fmt.Errorf("identifier is required")
		}

		constructorArgsStr, err := flags.GetString("constructor-args")
		if err != nil {
			return fmt.Errorf("invalid constructor flag: %v", err)
		}

		argsList := strings.Split(constructorArgsStr, ",")

		deploymentData, err := d.box.deployContract(d.box.Contracts[contractIdentifier], argsList...)
		if err != nil {
			return fmt.Errorf("failed to deploy the contract: %v", err)
		}

		deployments = append(deployments, deploymentData)

	case "by-file":
		metadataFilePath, err := flags.GetString("path")
		if err != nil {
			return fmt.Errorf("invalid path flag: %v", err)
		}

		if metadataFilePath == "" {
			return fmt.Errorf("path is required")
		}

		constructorArgsStr, err := flags.GetString("constructor-args")
		if err != nil {
			return fmt.Errorf("invalid constructor flag: %v", err)
		}

		argsList := strings.Split(constructorArgsStr, ",")

		contractMetadata, err := GetContractMetadataByFile(metadataFilePath)
		if err != nil {
			return fmt.Errorf("failed to load metadata: %v", err)
		}

		deploymentData, err := d.box.deployContract(contractMetadata, argsList...)
		if err != nil {
			return fmt.Errorf("failed to deploy the contract: %v", err)
		}

		deployments = append(deployments, deploymentData)
	}

	if err := writeDeploymentDataToFile(deployments); err != nil {
		return fmt.Errorf("failed to write deployment data: %v", err)
	}
	return nil
}
