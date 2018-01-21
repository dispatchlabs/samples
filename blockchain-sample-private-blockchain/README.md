# ![](https://storage.googleapis.com/material-icons/external-assets/v4/icons/svg/ic_power_settings_new_black_24px.svg) Dev Env Setup

###### `Mist` + `geth` + `Solidity`
```text
https://github.com/ethereum/mist
	Mist Browser
	Ethereum wallet and Dapp browser

https://github.com/ethereum/go-ethereum/wiki/geth
	Go implementation of the Ethereum protocol.
	Runs a full Ethereum node. Provides JS console interface

https://github.com/ethereum/solidity
	Contract-Oriented Programming Language
```

- Arch Linux
	- `yaourt mist`
	- `pacman -S geth`
	- `pacman -S solidity`

# Create Private Blockchain
- `geth --port 30304 --identity "SamplePrivateBlockchainNode" --nodiscover --networkid 1999 --datadir ./data init genesis.json`
- `geth account list --datadir ./data`
- `geth account new --datadir ./data`
- `geth account new --datadir ./data`
- `geth account list --datadir ./data`

###### Init Account With $$$
- Copy one of the newly creted accounts to the `genesis.json`
- `geth removedb --datadir ./data`
- `geth --port 30304 --identity "SamplePrivateBlockchainNode" --nodiscover --networkid 1999 --datadir ./data init genesis.json`

###### Run Private Blockchain Node
- `geth --port 30304 --identity "SamplePrivateBlockchainNode" --nodiscover --networkid 1999 --datadir ./data`
- `geth attach ./data/geth.ipc`

###### Compile First Contract
- `solc --bin --abi SampleContract.sol`

# Deploy Contract From Shell
```javascript
var testAddress = "0x708bbcb8f02a7141ee61d8897d243d75302e8eb9";
miner.start()
personal.unlockAccount(web3.eth.coinbase)
web3.eth.sendTransaction({from: web3.eth.coinbase, to: testAddress, value: web3.toWei("0.1","ether")})
web3.eth.getBalance(testAddress)


personal.unlockAccount(testAddress)
web3.eth.defaultAccount = testAddress


var abi = ... // JSON value from the `solc` earlier
var myContract = web3.eth.contract(abi)


var contractData = myContract.new.getData({data: "0x..."}) // binary value from the `solc` earlier

var myContractInstace = myContract.new({data: contractData, from: testAddress, gas: 1000000})
```

###### Call Cntract Function
```javascript
var beforeBalance = web3.eth.getBalance(testAddress)
personal.unlockAccount(testAddress)
myContractInstace.set(50)
var afterBalance = web3.eth.getBalance(testAddress)
var cost = beforeBalance - afterBalance
```

# Deploy Contract From Mist
 - If any, close Mist, open terminal and run `geth`
 - Copy the `geth.ipc` path
 - `geth --mine --datadir ./data --rpc --ipcpath /home/nicu/.ethereum/geth.ipc`
 - Start Mist and see the mining on main account, then deploy contract for the second account
