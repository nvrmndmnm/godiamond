package main

import (
	"log"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/nvrmndmnm/godiamond/internal/contracts"
)

func (box *DiamondBox) deploy() error {
	var tx *types.Transaction
	var err error

	box.diamondCutFacet, tx, _, err = contracts.DeployDiamondCutFacet(box.auth, box.client)
	if err != nil {
		return err
	}

	log.Printf("DiamondCutFacet address: %s\ntx: %s",
		box.diamondCutFacet.Hex(), tx.Hash().Hex())

	owner := box.config.Accounts["global"].Address
	box.diamond, tx, _, err = contracts.DeployDiamond(box.auth, box.client, owner, box.diamondCutFacet)
	if err != nil {
		return err
	}
	log.Printf("Diamond address: %s\ntx: %s",
		box.diamond.Hex(), tx.Hash().Hex())

	box.diamondInit, tx, _, err = contracts.DeployDiamondInit(box.auth, box.client)
	if err != nil {
		return err
	}
	log.Printf("DiamondInit address: %s\ntx: %s",
		box.diamondInit.Hex(), tx.Hash().Hex())

	loupeAddress, tx, _, err := contracts.DeployDiamondLoupeFacet(box.auth, box.client)
	if err != nil {
		return err
	}
	box.facets = append(box.facets, loupeAddress)
	log.Printf("DeployDiamondLoupeFacet address: %s\ntx: %s",
		loupeAddress.Hex(), tx.Hash().Hex())

	ownershipAddress, tx, _, err := contracts.DeployOwnershipFacet(box.auth, box.client)
	if err != nil {
		return err
	}
	box.facets = append(box.facets, ownershipAddress)
	log.Printf("DeployOwnershipFacet address: %s\ntx: %s",
		ownershipAddress.Hex(), tx.Hash().Hex())

	return nil
}
