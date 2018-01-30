package main

import (
	"fmt"
	"dispatchlabs/samples/blockchain-grpc/server"
	"dispatchlabs/samples/blockchain-grpc/client"
	"time"
)

func main() {
	fmt.Printf("In Main *****\n")
	go server.Start()
	time.Sleep(2 * time.Second)

	for i := 0; i < 3; i++ {
		go client.Start("add")
		time.Sleep(2 * time.Second)
	}
	client.Start("list")

}
