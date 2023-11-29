#!/bin/bash

fork_url=$(yq -r .rpc.alchemy <config.yaml)

anvil --fork-url $fork_url \
    --port 8545 --host 0.0.0.0 \
    --steps-tracing --no-rate-limit \
    --compute-units-per-second 10000
