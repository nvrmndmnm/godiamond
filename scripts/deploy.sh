#!/bin/bash

rpc_url=$(yq -r .rpc.local <config.yaml)
private_key=$(yq -r .eoa.anvil.private_key <config.yaml)
contract_owner=$(yq -r .eoa.anvil.address <config.yaml)

facets=('DiamondLoupeFacet' 'OwnershipFacet')

diamond_cut_facet=$(forge create --rpc-url $rpc_url \
    --private-key $private_key \
    contracts/src/facets/DiamondCutFacet.sol:DiamondCutFacet \
    | grep "Deployed to" | cut -d ' ' -f 3)
echo "DiamondCutFacet address: " $diamond_cut_facet 

forge create --rpc-url $rpc_url \
    --constructor-args "$contract_owner" "$diamond_cut_facet" \
    --private-key $private_key \
    contracts/src/Diamond.sol:Diamond

forge create --rpc-url $rpc_url \
    --private-key $private_key \
    contracts/src/upgradeInitializers/DiamondInit.sol:DiamondInit

for facet in "${facets[@]}"
do
    facet_address=$(forge create --rpc-url $rpc_url \
        --private-key $private_key \
        contracts/src/facets/${facet}.sol:${facet} \
        | grep "Deployed to" | cut -d ' ' -f 3)
    echo $facet "address:" $facet_address
done