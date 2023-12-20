#!/bin/bash

# This script might be useful for further growing and upgrading your Diamond contracts.
# It generates Go bindings for all Solidity contracts in the source directory.
#
# Ensure that your contracts are properly built (refer to scripts/build.sh)
# and abigen from go-ethereum is installed on your system.
src=contracts/src
out=contracts/out

mkdir -p internal/diamond
mkdir -p internal/facets 

find $src -type f -name "*.sol" \
 -not -path "*/interfaces/*" \
 -not -path "*/libraries/*" \
 -print0 | while read -d $'\0' file; do
    name=${${file##*/}%.sol}

    yq -r .bytecode.object \
    <$out/$name.sol/$name.json \
    >$out/$name.sol/$name.bin 

    n=$name
    declare -l n

    # The Diamond bindings are placed to its own package to avoid type clashes.
    if [ "$name" != "Diamond" ]; then \
        abigen --abi $out/$name.sol/$name.abi.json \
            --bin $out/$name.sol/$name.bin \
            --pkg facets \
            --type $name \
            --out internal/facets/$n.go;
    else \
        abigen --abi $out/$name.sol/$name.abi.json \
                --bin $out/$name.sol/$name.bin \
                --pkg diamond \
                --type $name \
                --out internal/diamond/$n.go;
    fi
done
