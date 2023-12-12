#!/bin/bash

forge build --extra-output-files abi metadata --force

src=contracts/src
out=contracts/out

mkdir -p internal/diamond
mkdir -p internal/contracts 

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

    if [ "$name" != "Diamond" ]; then \
        abigen --abi $out/$name.sol/$name.abi.json \
            --bin $out/$name.sol/$name.bin \
            --pkg contracts \
            --type $name \
            --out internal/contracts/$n.go;
    else \
        abigen --abi $out/$name.sol/$name.abi.json \
                --bin $out/$name.sol/$name.bin \
                --pkg diamond \
                --type $name \
                --exc IDiamondCut \
                --out internal/diamond/$n.go;
    fi
done
