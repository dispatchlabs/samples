package main

import (
	"log"
	"com.brightapps/src/grpc-blockchain/proto"
	"github.com/hashicorp/consul/api"
	"google.golang.org/grpc"
	"time"
	"strings"
	"fmt"
)


type node struct {
	// Self information
	Name string
	Addr string

	// Consul related variables
	SDAddress string
	SDKV      api.KV

	// used to make requests
	Clients map[string]proto.BlockchainClient
}

func (n *node) registerService() {
	config := api.DefaultConfig()
	config.Address = n.SDAddress
	consul, err := api.NewClient(config)
	if err != nil {
		log.Panicln("Unable to contact Service Discovery.")
	}

	kv := consul.KV()
	p := &api.KVPair{Key: n.Name, Value: []byte(n.Addr)}
	_, err = kv.Put(p, nil)
	if err != nil {
		log.Panicln("Unable to register with Service Discovery. %v", err)
	}

	// store the kv for future use
	n.SDKV = *kv

	log.Println("Successfully registered with Consul.")
}

func (n *node) Start() {
	// init required variables
	n.Clients = make(map[string]proto.BlockchainClient)

	// register with the service discovery unit
	n.registerService()

	// start the main loop here
	// in our case, simply time out for 1 minute and greet all

	// wait for other nodes to come up
	for {
		time.Sleep(20 * time.Second)
		n.GreetAll()
	}
}

func test() {

}

func (n *node) SetupClient(name string, addr string) {

	// setup connection with other node
	conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("cannot dial server: %v", err)
	}

	defer conn.Close()
	n.Clients[name] = proto.NewBlockchainClient(conn)

	//r, err := n.Clients[name].SayHello(context.Background(), &hs.HelloRequest{Name: n.Name})
	//if err != nil {
	//	log.Fatalf("could not greet: %v", err)
	//}
	//log.Printf("Greeting from the other node: %s", r.Message)

}

// Busy Work module, greet every new member you find
func (n *node) GreetAll() {
	// get all nodes -- inefficient, but this is just an example
	kvpairs, _, err := n.SDKV.List("Node", nil)
	if err != nil {
		log.Panicln(err)
		return
	}

	// fmt.Println("Found nodes: ")
	for _, kventry := range kvpairs {
		if strings.Compare(kventry.Key, n.Name) == 0 {
			// ourself
			continue
		}
		if n.Clients[kventry.Key] == nil {
			fmt.Println("New member: ", kventry.Key)
			// connection not established previously
			n.SetupClient(kventry.Key, string(kventry.Value))
		}
	}
}


func main() {
	log.Print("in main")
	// pass the port as an argument and also the port of the other node
	//args := os.Args[1:]
	//
	//if len(args) < 3 {
	//	fmt.Println("Arguments required: <name> <listening address> <consul address>")
	//	os.Exit(1)
	//}

	// args in order
	name := "Node3"
	listenaddr := ":5002"
	sdaddress := "localhost:8500"

	noden := node{Name: name, Addr: listenaddr, SDAddress: sdaddress, Clients: nil} // noden is for opeartional purposes

	// start the node
	noden.Start()
}
