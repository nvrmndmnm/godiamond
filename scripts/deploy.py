#!/usr/bin/env python3

# Use this script for development purposes to quickly deploy Diamond smart contracts.
# It extracts necessary parameters from a config file. After deploying the contracts,
# it prints their addresses to the console, allowing you to interact with them immediately.
#
# Ensure that Foundry and PyYAML are installed on your system.
import os
import subprocess
import yaml

# Customize an array of facet names that you want to deploy
facets = ['DiamondLoupeFacet', 'OwnershipFacet']

# Load the config file
with open('config.yaml', 'r') as file:
    config = yaml.safe_load(file)

# Extract necessary parameters
rpc_url = config['rpc']['local']
private_key = config['eoa']['anvil']['private_key']
contract_owner = config['eoa']['anvil']['address']

# Deploy DiamondCutFacet
output = subprocess.check_output([
    'forge', 'create', '--rpc-url', rpc_url, '--private-key', private_key,
    'contracts/src/facets/DiamondCutFacet.sol:DiamondCutFacet'
])
diamond_cut_facet = output.decode().split('Deployed to: ')[1].split()[0]
print(f'DiamondCutFacet address: {diamond_cut_facet}')

# Deploy Diamond
output = subprocess.check_output([
    'forge', 'create', '--rpc-url', rpc_url, 
    '--constructor-args', contract_owner, diamond_cut_facet,
    '--private-key', private_key, 'contracts/src/Diamond.sol:Diamond'
])
diamond = output.decode().split('Deployed to: ')[1].split()[0]
print(f'Diamond address: {diamond}')

# Deploy DiamondInit
output = subprocess.check_output([
    'forge', 'create', '--rpc-url', rpc_url, '--private-key', private_key,
    'contracts/src/upgradeInitializers/DiamondInit.sol:DiamondInit'
])
diamond_init = output.decode().split('Deployed to: ')[1].split()[0]
print(f'DiamondInit address: {diamond_init}')

# Deploy facets
for facet in facets:
    output = subprocess.check_output([
        'forge', 'create', '--rpc-url', rpc_url, '--private-key', private_key,
        f'contracts/src/facets/{facet}.sol:{facet}'
    ])
    facet_address = output.decode().split('Deployed to: ')[1].split()[0]
    print(f'{facet} address: {facet_address}')
