#!/bin/sh
solc --version

solc --evm-version byzantium --combined-json abi,bin,opcodes --pretty-json DG_304.sol > DG_304.json

ls -la
