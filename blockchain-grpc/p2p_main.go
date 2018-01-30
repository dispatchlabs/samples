package main

import (
	"fmt"
	"dispatchlabs/samples/blockchain-grpc/p2p"
	"time"
)

func main() {
	fmt.Printf("In Main *****\n")

	node1 := p2p.Start("Node1", "127.0.0.1:5001")
	node2 := p2p.Start("Node2", "127.0.0.1:5002")
	node3 := p2p.Start("Node3", "127.0.0.1:5003")

	time.Sleep(10 * time.Second)

	node1.RandomTalk()
	node2.RandomTalk()
	node3.RandomTalk()
	time.Sleep(30 * time.Second)
}
