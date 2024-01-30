#!/usr/bin/env python3

# Use this script to fetch Solidity source files from official 
# Diamond implementations built with Hardhat.
# These source files can then be used with Foundry.
#
# The specific repository version can be adjusted as necessary
import os
import shutil

diamond_repo = "https://github.com/mudgen/diamond-3-hardhat.git"

# Create a 'contracts' directory and clone the repo into it
os.system(f'git clone --no-checkout --depth=1 --filter=tree:0 {diamond_repo} contracts')

# Perform a sparse-checkout for 'contracts' directory
os.chdir('contracts')
os.system('git sparse-checkout set --no-cone contracts')
os.system('git checkout')

# Create 'src' directory, move contracts into it, and clean up
os.makedirs('src', exist_ok=True)
for file in os.listdir('contracts'):
    shutil.move(os.path.join('contracts', file), 'src')
shutil.rmtree('.git')
shutil.rmtree('contracts')

# Change back to the original directory
os.chdir('..')
