#!/bin/bash

# Use this script to build source smart contracts and generate extra output files 
# such as ABI and metadata, which will be used later in Go Diamond.
#
# Ensure that Foundry is installed on your system.
forge build --extra-output-files abi metadata --skip script --force
