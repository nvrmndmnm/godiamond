#!/bin/bash

forge build --extra-output-files abi metadata --force

src=contracts/src
out=contracts/out

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

    abigen --abi $out/$name.sol/$name.abi.json \
        --bin $out/$name.sol/$name.bin \
        --pkg contracts \
        --type $name \
        --out internal/contracts/$n.go
done
