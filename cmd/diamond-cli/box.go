package main

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

type DiamondBox struct {
	config Config
	client *ethclient.Client
	auth   *bind.TransactOpts

	diamond           common.Address
	diamondInit       common.Address
	diamondCutFacet   common.Address
	diamondLoupeFacet common.Address
}

func NewDiamondBox(config Config, rpc string, chainId int64) (*DiamondBox, error) {
	var err error
	box := &DiamondBox{
		config:            config,
		diamond:           config.Contracts["diamond"].Address,
		diamondInit:       config.Contracts["diamond_init"].Address,
		diamondCutFacet:   config.Contracts["cut_facet"].Address,
		diamondLoupeFacet: config.Contracts["loupe_facet"].Address,
	}

	box.client, err = ethclient.Dial(config.RPC[rpc])
	if err != nil {
		return nil, err
	}

	if chainId == 0 {
		networkId, err := box.client.NetworkID(context.Background())
		if err != nil {
			return nil, err
		}

		chainId = networkId.Int64()
	}

	privateKey, err := crypto.HexToECDSA(config.Accounts["anvil"].PrivateKey[2:])
	if err != nil {
		return nil, err
	}

	box.auth, err = bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(chainId))
	if err != nil {
		return nil, err
	}

	gasPrice, err := box.client.SuggestGasPrice(context.Background())
	if err != nil {
		return nil, err
	}
	box.auth.GasPrice = gasPrice

	return box, nil
}

func (box *DiamondBox) Close() {
	box.client.Close()
}
