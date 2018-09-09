#!/bin/sh
solc --version

rm -rf CallDeployed
solc -o CallDeployed --abi --bin --pretty-json CallDeployed.sol

rm -rf Deployed
solc -o Deployed     --abi --bin --pretty-json Deployed.sol