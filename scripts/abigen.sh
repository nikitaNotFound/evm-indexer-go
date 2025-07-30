#!/bin/bash

ABI_DIR="./internal/smartcontracts/abi"
OUT_DIR="./internal/smartcontracts/abigen"
PKG_NAME="abigen"

mkdir -p "$OUT_DIR"

for abi_file in "$ABI_DIR"/*abi.json; do
    [ -e "$abi_file" ] || continue  # skip if no match

    base_name=$(basename "$abi_file")
    type_name="${base_name%%.*}" # gets "MyContract" from "MyContract.abi.json"
    out_file="$OUT_DIR/${type_name}.abigen.go"

    abigen --abi "$abi_file" --pkg "$PKG_NAME" --type "$type_name" --out "$out_file"
done