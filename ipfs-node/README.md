# WARNING 
THIS IS Alpha Software, has tons of issues.
If something is not working as expected restart the daemon.

# Install
- Mac
	- `curl -O https://dist.ipfs.io/go-ipfs/v0.4.13/go-ipfs_v0.4.13_darwin-amd64.tar.gz`
	- `tar -xzvf ipfs.tar.gz`
	- `cd go-ipfs`
	- `sudo ./install.sh`

- Arch Linux
	- `sudo pacman -S go-ipfs`



# Setup
- `ipfs init`
	- dumps your ID as well
	- `cat ~/.ipfs/config | grep PeerID` to find it later
- `ipfs daemon`



# Cheat Sheet
| Run																							| Dox
|----																							|----------------------------------------------
|`cat ~/.ipfs/config | grep PeerID`																| find your ID on the IPFS netowrk
|`cat ~/.ipfs/config | grep PeerID | sed 's/"PeerID": "//g' | sed 's/",//g' | sed 's/ //g'`		| find your ID, remove the noise
|`ipfs id | grep ID | sed 's/"ID": "//g' | sed 's/",//g' | sed 's/\t//g'`						| same thing, find your ID, remove the noise
|`ipfs id | grep PublicKey | sed 's/"PublicKey": "//g' | sed 's/",//g' | sed 's/\t//g'`			| find your Public Key
|`cat ~/.ipfs/config | grep PrivKey | sed 's/"PrivKey": "//g' | sed 's/",//g' | sed 's/ //g'`	| find your Private Key
|`ipfs`																							| list of all commands
|`ipfs cat /ipfs/$MyIPFSNode`																	| list my exposed content
|`http://localhost:5001/webui`																	| Web UI for local node
|`http://localhost:8080/ipfs/QmVc6zuAneKJzicnJpfrqCH9gSy6bz54JhcypfJYhGUFQu/play#/ipfs/QmTKZgRNwDNZwHtJSjCp6r5FYefzpULfy37JvMt9DwvXse` | Play Video
|`http://localhost:8080/ipfs/QmZpc3HvfjEXvLWGQPWbHk3AjD5j8NEN4gmFN8Jmrd5g83/cs`					| List folder content
|`http://localhost:8080/ipfs/QmX7M9CiYXjVeFnkfVGf3y5ixTZ2ACeSGyL1vBJY1HvQPp/mdown`				| Markdown renderer app

| Run																							| Dox
|----																							|----------------------------------------------
|`ipfs cat /ipfs/xxxxxxxxxxxxx/readme`															| the hash here will be the one from `ipfs init`
|`ipfs cat /ipfs/xxxxxxxxxxxxx/quick-start`														| show of various ipfs features: add, view, list files/folders
|`ipfs --help`																					| every command describes itself

- My `ipfs cat /ipfs/xxxxxxxxxxxxx/readme` was like this `ipfs cat /ipfs/QmS4ustL54uo8FzR9455qaxZwuMiUhyvMcX9Ba8nUH4uVv/readme`
- After setup this local folder is shared on the network automatically. It is stored under the `ls ~/.ipfs/datastore`
- The content blocks are stored within $HOME/.ipfs, you need to use ipfs as a medium for accessing them.
  Larger files get split into 256K block, and the blockstore is optimized for performance.

##### Artifacts Operations
- `ipfs add FILE`
	- `ipfs cat FILE_HASH`
- `ipfs add -r FOLDER`
	- `ipfs ls FOLDER_HASH`
- REMOVE a file / folder as per `https://github.com/ipfs/faq/issues/9#issuecomment-140516800`
	-
	```text
	On the IPFS homepage, it is described as IPFS is The Permanent Web so deletion sounds like an anti-goal to me :)

	To be clear, though: it is possible for data to be deleted from the network. For example, if you add a file and then immediately unpin it and then garbage collect it before anyone has a chance to get it from you, it will be effectively deleted.
	```
	- `ipfs add FILE` -> spits out a HASH
	- `ipfs pin rm $HASH`
	- `ipfs repo gc`

- `ipfs ls PEER_HASH/FOLDER_HASH` - lists content, display files/folders with hashes
	- `ipfs refs PEER_HASH/FOLDER_HASH` - display hashes only for the content
- `ipfs get FILE`
- `ipfs object FILE`

| Run																								| Dox
|----																								|----------------------------------------------
|`ipfs object data <key>`																			| Output the raw bytes of an IPFS object.
|`ipfs object diff <obj_a> <obj_b>`																	| Display the diff between two ipfs objects.
|`ipfs object get <key>`																			| Get and serialize the DAG node named by <key>.
|`ipfs object links <key>`																			| Output the links pointed to by the specified object.
|`ipfs object new [<template>]`																		| Create a new object from an ipfs template.
|`ipfs object patch`																				| Create a new merkledag object based on an existing one.
|`ipfs object put <data>`																			| Store input as a DAG object, print its key.
|`ipfs object stat <key>`																			| Get stats for the DAG node named by <key>.



# Sample Sesion
- Terminal 1
	- `ipfs daemon`
- Terminal 2
	- `export MyIPFSNode=$(cat ~/.ipfs/config | grep PeerID | sed 's/"PeerID": "//g' | sed 's/",//g' | sed 's/ //g')`
	- `echo $MyIPFSNode`
	- `firefox http://localhost:5001/webui`
	- Add File
		- `echo "Test Content" > test-file.text`
		- `ipfs add test-file.text` will output the hash for the new file on the network
		- Refresh the WebUI and see it there under the `Files`
		- `ipfs cat /ipfs/NEW_HASH` will show `Test Content`
		OR
		- In the Web UI under the `DAG` type the `NEW_HASH` value and hit go, then RAW to see the file content or the video
	- Remove File and all other content (if no one except you accessed it)
		- `ipfs pin rm $HASH1`
		- `ipfs pin rm $HASH2`
		- `ipfs pin rm $HASH3`
		- etc
		- `ipfs repo gc`





# Dox IPFS -- Inter-Planetary File system
- IPFS is a global, versioned, peer-to-peer filesystem
- Combines ideas 
	- Git
	- BitTorrent
	- Kademlia
	- SFS
	- Web
- It is like a single bit-torrent swarm, exchanging git objects
- IPFS provides an interface as simple as the HTTP web, but with permanence built in
- You can also mount the world at /ipfs.

```text
IPFS is a protocol:
- defines a content-addressed file system
- coordinates content delivery
- combines Kademlia + BitTorrent + Git

IPFS is a filesystem:
- has directories and files
- mountable filesystem (via FUSE)

IPFS is a web:
- can be used to view documents like the web
- files accessible via HTTP at `http://ipfs.io/<path>`
- browsers or extensions can learn to use `ipfs://` directly
- hash-addressed content guarantees authenticity

IPFS is modular:
- connection layer over any network protocol
- routing layer
- uses a routing layer DHT (kademlia/coral)
- uses a path-based naming service
- uses bittorrent-inspired block exchange

IPFS uses crypto:
- cryptographic-hash content addressing
- block-level deduplication
- file integrity + versioning
- filesystem-level encryption + signing support

IPFS is p2p:
- worldwide peer-to-peer file transfers
- completely decentralized architecture
- __no__ central point of failure

IPFS is a cdn:
- add a file to the filesystem locally, and it's now available to the world
- caching-friendly (content-hash naming)
- bittorrent-based bandwidth distribution

IPFS has a name service:
- IPNS, an SFS inspired name system
- global namespace based on PKI
- serves to build trust chains
- compatible with other NSes
- can map DNS, .onion, .bit, etc to IPNS
```