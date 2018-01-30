package blockchain

import (
	"time"
	"github.com/dispatchlabs/samples/blockchain-grpc/proto"
)

// Blockchain keeps a sequence of Blocks
type Blockchain struct {
	Blocks []*proto.Block
}

// setHash calculates and sets block hash


// NewBlock creates and returns Block
func NewBlock(data string, prevBlockHash string) *proto.Block {
	var block *proto.Block
	block = &proto.Block{data, prevBlockHash, "", time.Now().Unix()}
	block.SetHash()

	return block
}

// NewGenesisBlock creates and returns genesis Block
func NewGenesisBlock() *proto.Block {
	return NewBlock("Genesis Block", "")
}

func IsValidNewBlock(newBlock proto.Block, previousBlock proto.Block) bool {
	return true
}

// AddBlock saves provided data as a block in the blockchain
func (bc *Blockchain) AddBlock(data string) *proto.Block {
	prevBlock := bc.Blocks[len(bc.Blocks)-1]
	newBlock := NewBlock(data, prevBlock.Hash)
	bc.Blocks = append(bc.Blocks, newBlock)

	return newBlock
}

// NewBlockchain creates a new Blockchain with genesis Block
func NewBlockchain() *Blockchain {
	return &Blockchain{[]*proto.Block{NewGenesisBlock()}}
}
