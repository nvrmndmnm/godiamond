#!/bin/bash

# Use this script to fetch Solidity source files from official 
# Diamond implementations built with Hardhat.
# These source files can then be used with Foundry.
#
# The specific repository version can be adjusted as necessary.
# Compare versions: https://github.com/mudgen/diamond
diamond_repo=https://github.com/mudgen/diamond-3-hardhat.git

# Clone, perform a sparse-checkout for 'contracts' directory, 
# clean up and move contracts to 'src' directory.
git clone --no-checkout --depth=1 --filter=tree:0 \
  $diamond_repo contracts
cd contracts
git sparse-checkout set --no-cone contracts
git checkout
mkdir src 
mv contracts/* src
rm -rf .git contracts
cd ..