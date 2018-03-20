package client

import (
	"fmt"
	"log"
	"time"

	"github.com/dispatchlabs/samples/blockchain-grpc/proto"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

var client proto.BlockchainClient

func Start(flag string) {
	fmt.Println("Starting Client")

	//addFlag := flag.Bool("add", false, "Add new block")
	//listFlag := flag.Bool("list", false, "List all blocks")
	//flag.Parse()

	conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("cannot dial server: %v", err)
	}
	defer conn.Close()

	client = proto.NewBlockchainClient(conn)
	if flag == "add" {
		AddBlock()
	}
	if flag == "list" {
		GetBlockchain()
	}
	//if *addFlag {
	//	AddBlock()
	//}
	//
	//if *listFlag {
	//	GetBlockchain()
	//}
}

func AddBlock() {
	block, addErr := client.AddBlock(context.Background(), &proto.AddBlockRequest{
		Data: time.Now().String(),
	})
	if addErr != nil {
		log.Fatalf("unable to add block: %v", addErr)
	}
	log.Printf("new block hash: %s\n", block.Hash)
}

func GetBlockchain() {
	blockchain, getErr := client.GetBlockchain(context.Background(), &proto.GetBlockchainRequest{})
	if getErr != nil {
		log.Fatalf("unable to get blockchain: %v", getErr)
	}

	log.Println("blocks:")
	for _, b := range blockchain.Blocks {
		log.Printf("hash %s, prev hash: %s, data: %s, timestamp: %d\n", b.Hash, b.PrevBlockHash, b.Data, b.Timestamp)
	}
}
