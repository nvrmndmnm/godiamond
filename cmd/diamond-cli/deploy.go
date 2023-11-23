package main

import (
	"context"
	"fmt"
	"math/big"

	"log"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/nvrmndmnm/godiamond/internal/contracts"
)

func deploy(config Config, rpcIdentifier string, chainId int64) (error) {
	fmt.Println("deploy")
	//init client
	client, err := ethclient.Dial(config.RPC["local"])
	if err != nil {
		return err
	}

	if chainId == 0 {
		log.Println("chainId is 0")
		networkId, err := client.NetworkID(context.Background())
		if err != nil {
			return err
		}

		chainId = networkId.Int64()
	}
	//init auth
	privateKey, err := crypto.HexToECDSA(config.Accounts["anvil"].PrivateKey[2:])
	if err != nil {
		return err
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(chainId))
	if err != nil {
		return err
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		return err
	}

	auth.GasPrice = gasPrice
	fmt.Println(chainId)

	//deploy DiamondCutFacet
	address, tx, _, err := contracts.DeployDiamondCutFacet(auth, client)
	if err != nil {
		return err
	}
	log.Printf("deployed DiamondCutFacet at address %s, tx %s",
		address.Hex(), tx.Hash().Hex())

	//deploy Diamond
	owner := config.Accounts["global"].Address
	cutfacet := address

	address, tx, _, err = contracts.DeployDiamond(auth, client, owner, cutfacet)
	if err != nil {
		return err
	}

	log.Printf("deployed Diamond at address %s, tx %s",
		address.Hex(), tx.Hash().Hex())

	//deploy DiamondInit
	address, tx, _, err = contracts.DeployDiamondInit(auth, client)
	if err != nil {
		return err
	}
	log.Printf("deployed DiamondInit at address %s, tx %s",
		address.Hex(), tx.Hash().Hex())

	//deploy facets
	address, tx, _, err = contracts.DeployDiamondLoupeFacet(auth, client)
	if err != nil {
		return err
	}
	log.Printf("deployed DeployDiamondLoupeFacet at address %s, tx %s",
		address.Hex(), tx.Hash().Hex())

	address, tx, _, err = contracts.DeployOwnershipFacet(auth, client)
	if err != nil {
		return err
	}
	log.Printf("deployed DeployOwnershipFacet at address %s, tx %s",
		address.Hex(), tx.Hash().Hex())

	return nil
}
