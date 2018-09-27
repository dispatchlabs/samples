#!/bin/sh
solc --version

solc --evm-version byzantium --combined-json abi,bin,opcodes --pretty-json ./sol/DG_855.sol > ./json/DG_855.json
solc --evm-version byzantium --combined-json abi,bin,opcodes --pretty-json ./sol/DG_854.sol > ./json/DG_854.json
solc --evm-version byzantium --combined-json abi,bin,opcodes --pretty-json ./sol/DG_822.sol > ./json/DG_822.json

solc --evm-version byzantium --combined-json abi,bin,opcodes --pretty-json ./sol/DG_702_1.sol > ./json/DG_702_1.json
solc --evm-version byzantium --combined-json abi,bin,opcodes --pretty-json ./sol/DG_702_2.sol > ./json/DG_702_2.json
solc --evm-version byzantium --combined-json abi,bin,opcodes --pretty-json ./sol/DG_702_3.sol > ./json/DG_702_3.json

solc --evm-version byzantium --combined-json abi,bin,opcodes --pretty-json ./sol/DG_895.sol > ./json/DG_895.json
solc --evm-version byzantium --combined-json abi,bin,opcodes --pretty-json ./sol/DG_896.sol > ./json/DG_896.json

solc --evm-version byzantium --combined-json abi,bin,opcodes --pretty-json ./sol/DG_862.sol > ./json/DG_862.json

solc --evm-version byzantium --combined-json abi,bin,opcodes --pretty-json ./sol/DG_891.sol > ./json/DG_891.json

solc --evm-version byzantium --combined-json abi,bin,opcodes --pretty-json ./sol/DG_897.sol > ./json/DG_897.json

solc --evm-version byzantium --combined-json abi,bin,opcodes --pretty-json ./sol/DG_856.sol > ./json/DG_856.json

solc --evm-version byzantium --combined-json abi,bin,opcodes --pretty-json ./sol/DG_851.sol > ./json/DG_851.json

solc --evm-version byzantium --combined-json abi,bin,opcodes --pretty-json ./sol/DG_892.sol > ./json/DG_892.json


ls -la
