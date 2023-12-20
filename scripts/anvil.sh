#!/bin/bash

# Use this script to run the Anvil RPC from Foundry for development and testing purposes.
# It extracts the RPC URL from a config file and uses it to run the simulator as a fork.
#
# Ensure that Foundry and yq are installed on your system.
fork_url=$(yq -r .rpc.alchemy <config.yaml)

anvil --fork-url $fork_url \
    --port 8545 \
     --host 0.0.0.0 \
    --steps-tracing \
    --no-rate-limit \
    --compute-units-per-second 10000
