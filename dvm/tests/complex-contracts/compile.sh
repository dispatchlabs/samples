#!/bin/sh
solc --version

solc --evm-version byzantium --combined-json abi,bin,opcodes --pretty-json CallDeployed.sol > CallDeployed.json
solc --evm-version byzantium --combined-json abi,bin,opcodes --pretty-json CallDeployed-v2.sol > CallDeployed-v2.json
solc --evm-version byzantium --combined-json abi,bin,opcodes --pretty-json Deployed.sol > Deployed.json

ls -la
