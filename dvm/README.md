# Ethereum Code Merge @ Sep 6 2018
- Step 1
	- `go get github.com/ethereum/go-ethereum`
	- The `~/go/src/github.com/ethereum/go-ethereum` has this content
	```text
	accounts/
	appveyor.yml
	AUTHORS
	build/
	circle.yml
	cmd/
	common/
	consensus/
	console/
	containers/
	contracts/
	COPYING
	COPYING.LESSER
	core/
	crypto/
	dashboard/
	Dockerfile
	Dockerfile.alltools
	.dockerignore
	eth/
	ethclient/
	ethdb/
	ethstats/
	event/
	files.text
	.git/
	.gitattributes
	.github/
	.gitignore
	.gitmodules
	interfaces.go
	internal/
	les/
	light/
	log/
	.mailmap
	Makefile
	metrics/
	miner/
	mobile/
	node/
	p2p/
	params/
	README.md
	rlp/
	rpc/
	signer/
	swarm/
	tests/
	.travis.yml
	trie/
	vendor/
	whisper/
	```
	- Remove all top files and folders except these
	```text
	accounts/
	common/
	core/
	crypto/
	ethdb/
	interfaces.go
	log/
	params/
	rlp/
	trie/
	```
- Step 2
	- open in code editor the `~/go/src/github.com/ethereum/go-ethereum`
		- for the imports replace in all files, FOLLOW ORDER
			- `github.com/ethereum/go-ethereum` -> `github.com/dispatchlabs/disgo/dvm/ethereum`
			- `github.com/dispatchlabs/disgo/dvm/ethereum/core/types` -> `github.com/dispatchlabs/disgo/dvm/ethereum/types`
			- `github.com/dispatchlabs/disgo/dvm/ethereum/accounts/abi` -> `github.com/dispatchlabs/disgo/dvm/ethereum/abi`
	- for next steps every MERGE has to be 
		- from `~/go/src/github.com/ethereum/go-ethereum`
		- to `~/go/src/github.com/dispatchlabs/disgo`
		- folder compare THEN MERGE every step below, make sure not overriding/deleting any custom code
		- merge `accounts/abi` -> `dvm/ethereum/abi`
		- merge `common/hexutil` -> `dvm/ethereum/common/hexutil`
		- merge `common/math` -> `dvm/ethereum/common/math`
		- merge `common/big.go` -> `dvm/ethereum/common/big.go`
		- merge `common/bytes.go` -> `dvm/ethereum/common/bytes.go`
		- merge `common/size.go` -> `dvm/ethereum/common/size.go`
		- merge `common/types.go` -> `dvm/ethereum/common/types.go`
		- merge `crypto/bn256` -> `dvm/ethereum/crypto/bn256`
		- don't touch `crypto/crypto.go` -> `dvm/ethereum/crypto/crypto.go`
		- merge `crypto/signature_cgo.go` -> `dvm/ethereum/crypto/signature_cgo.go`
		- merge `crypto/sha3/sha3.go` -> `dvm/ethereum/crypto/sha3.go`
		- merge `crypto/sha3/xor_unaligned.go` -> `dvm/ethereum/crypto/xor_unaligned.go`
		- merge `ethdb` -> `dvm/ethereum/ethdb`
		- merge `log` -> `dvm/ethereum/log`
		- merge `params` -> `dvm/ethereum/params`
		- merge `rlp` -> `dvm/ethereum/rlp`
		- merge `core/state` -> `dvm/ethereum/state`
		- merge `core/vm` -> `dvm/ethereum/vm`
		- merge `core/types` -> `dvm/ethereum/types`
		- merge `core/error.go` -> `dvm/ethereum/error.go`
		- merge `core/evm.go` -> `dvm/ethereum/evm.go`
		- merge `core/gaspool.go` -> `dvm/ethereum/gaspool.go`
		- merge `core/state_transition.go` -> `dvm/ethereum/state_transition.go`
		- merge `trie` -> `dvm/ethereum/trie`
		- merge `interfaces.go` -> `dvm/ethereum/interfaces.go`
