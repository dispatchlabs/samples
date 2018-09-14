#!/bin/sh
solc --version

solc --evm-version byzantium --combined-json abi,bin,opcodes --pretty-json DG_304.sol > DG_304.json
solc --evm-version byzantium --combined-json abi,bin,opcodes --pretty-json DG_854.sol > DG_854.json
solc --evm-version byzantium --combined-json abi,bin,opcodes --pretty-json DG_822.sol > DG_822.json

ls -la
