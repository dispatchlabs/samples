package main

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"time"

	"github.com/nic0lae/JerryMouse/Servers"
)

func NewBlock(previousHash string, timestamp int64, data string, hash string) Block {
	return Block{PreviousHash: previousHash, Timestamp: timestamp, Data: data, Hash: hash}
}

func GetGenesisBlock() Block {
	return NewBlock("0", 1465154705, "my genesis block!!", "816534932c2b7154836da6afc367695e6337db8a921823784c14378abed4f7d7")
}

func CalculateHash(previousHash string, timestamp int64, data string) string {
	shaThis := fmt.Sprintf("%s %d %s", previousHash, timestamp, data)
	shaThisAsBA := sha256.Sum256([]byte(shaThis))
	return string(shaThisAsBA[:])
}

func GetLatestBlock(blockchain []Block) Block {
	return blockchain[len(blockchain)-1]
}

func GenerateNextBlock(blockchain []Block, blockData string) Block {
	previousBlock := GetLatestBlock(blockchain)
	nextTimestamp := time.Now().Unix()
	nextHash := CalculateHash(previousBlock.Hash, nextTimestamp, blockData)
	return NewBlock(previousBlock.Hash, nextTimestamp, blockData, nextHash)
}

func IsValidNewBlock(newBlock Block, previousBlock Block) bool {
	return true
}
func AddBlock(blockchain []Block, newBlock Block) []Block {
	if IsValidNewBlock(newBlock, GetLatestBlock(blockchain)) {
		blockchain = append(blockchain, newBlock)
	}

	return blockchain
}

func ResponseLatestMsg(blockchain []Block) []byte {
	var response Servers.JsonResponse
	response.Data = GetLatestBlock(blockchain)

	responseAsByteArray, err := json.Marshal(response)
	if err == nil {
		return responseAsByteArray
	}

	return []byte{}
}
