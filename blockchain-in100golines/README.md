# ![](https://storage.googleapis.com/material-icons/external-assets/v4/icons/svg/ic_info_outline_black_24px.svg) Concepts

- `ledger`
	- a book containing accounts to which debits and credits are posted from books of original entry
	- Quote => _Transactions between accounts are recorded on online ledgers and prices posted publicly on exchanges such as Coinbase's GDAX, one of the indexes that tracks the value of bitcoin._
	- Quote => _The Bitcoin blockchain is the public, shared ledger that keeps track of payments in the Bitcoin network._
	- _Records in your ledger tell where the money is coming from and going to_

- `distributed ledger` aka `shared ledger` aka `distributed ledger technology`
	- A consensus of replicated, shared, and synchronized digital data geographically spread across multiple sites, countries, or institutions. There is no central administrator or centralised data storage.
	- A peer-to-peer network is required as well as consensus algorithms to ensure replication across nodes. One form of distributed ledger design is the blockchain, which can be either public or private. Not all distributed ledgers have to use blockchain to provide secure and valid achievement of distributed consensus. A blockchain is only one type of data structure considered to be a distributed ledger.

- [`hyperledger`](https://en.wikipedia.org/wiki/Hyperledger)
	- Umbrella project for a bunch of open source blockchains and related tools by Linux Foundation to support collaborative development of blockchain-based distributed ledgers
	- Hyperledger blockchain platforms
		- [Hyperledger Burrow](https://en.wikipedia.org/wiki/Hyperledger#Hyperledger_Burrow)
		- [Hyperledger Fabric](https://en.wikipedia.org/wiki/Hyperledger#Hyperledger_Fabric)
		- [Hyperledger Iroha](https://en.wikipedia.org/wiki/Hyperledger#Hyperledger_Iroha)
		- [Hyperledger Sawtooth](https://en.wikipedia.org/wiki/Hyperledger#Hyperledger_Sawtooth)
	- Hyperledger developer tooling
		- [Hyperledger Cello](https://en.wikipedia.org/wiki/Hyperledger#Hyperledger_Cello)
		- [Hyperledger Composer](https://en.wikipedia.org/wiki/Hyperledger#Hyperledger_Composer)
		- [Hyperledger Explorer](https://en.wikipedia.org/wiki/Hyperledger#Hyperledger_Explorer)
		- [Hyperledger Indy](https://en.wikipedia.org/wiki/Hyperledger#Hyperledger_Indy)

- `blockchain` a continuously growing list of records (blocks) which are linked and secured using cryptography
	- One block contains
		- hash pointer to previous block
		- timestamp
		- transaction data
	- Iinherently resistant to modification of the data
	- Open, distributed ledger that can record transactions between two parties efficiently and in a verifiable and permanent way
	- Once recorded data in the block cannot be altered retroactively without the alteration of all subsequent blocks. This can be done only with agreement of the network majority.
	- Secure by design and an example of distributed computing system with high Byzantine fault tolerance (aka [Byzantine Generals' Problem](https://en.wikipedia.org/wiki/Byzantine_fault_tolerance#Byzantine_Generals'_Problem) aka `every general must agree on a common decision, a halfhearted/chaotic attack by a few generals would count as huge damage and be worse than a coordinated attack or a coordinated retreat`)
	- Potentially suitable for the recording of events, medical records, identity management, transaction processing, documenting provenance, food traceability or voting.

# ![](https://storage.googleapis.com/material-icons/external-assets/v4/icons/svg/ic_code_black_24px.svg) Code

- Quote => _Basde on popular blockchain-based projects such as Bitcoin and Ethereum and problems they solve, the term `blockchain` is usually strongly tied to concepts like `transactions`, `smart contracts` or `cryptocurrencies`. For starters this can be confusing and can make understanding blockchains a muddier task than it has to be, especially source code/data structures wise._

###### Example Blockchains
- `https://github.com/lhartikk/naivechain`
- `https://github.com/lucrussell/tiny-blockchain`
- `https://github.com/Shachindra/Hack4Climate.git`

###### Sample Implmentation Go Blockhains in folders above
- `Go-NaiveChain`
