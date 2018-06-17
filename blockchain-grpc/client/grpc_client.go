package client

import (
	"fmt"
	"time"

	"github.com/dispatchlabs/disgo/commons/utils"
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
		utils.Fatal(fmt.Sprintf("cannot dial server: %v", err))
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
		utils.Fatal(fmt.Sprintf("unable to add block: %v", addErr))
	}

	utils.Info(fmt.Sprintf("new block hash: %s\n", block.Hash))
}

func GetBlockchain() {
	blockchain, getErr := client.GetBlockchain(context.Background(), &proto.GetBlockchainRequest{})
	if getErr != nil {
		utils.Fatal(fmt.Sprintf("unable to get blockchain: %v", getErr))
	}

	utils.Info("blocks:")

	for _, b := range blockchain.Blocks {
		utils.Info(fmt.Sprintf("hash %s, prev hash: %s, data: %s, timestamp: %d\n", b.Hash, b.PrevBlockHash, b.Data, b.Timestamp))
	}
}
