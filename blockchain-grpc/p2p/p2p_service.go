package p2p

import (
	"log"
	"github.com/dispatchlabs/samples/blockchain-grpc/proto"
	"github.com/hashicorp/consul/api"
	"google.golang.org/grpc"
	"time"
	"strings"
	"math/rand"
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
		time.Sleep(5 * time.Second)
		n.GreetAll()
	}
}

func (n *node) RandomTalk() {
	kvpairs, _, err := n.SDKV.List("Node", nil)
	if err != nil {
		log.Panicln(err)
		return
	}
	for {
		random := rand.Intn(len(kvpairs)-1)
		kventry := kvpairs[random]
		if strings.Compare(kventry.Key, n.Name) != 0 {
			// not ourself
			log.Printf("%s at address: %s randomly talking with %s at address: %s", n.Name, n.Addr, kventry.Key, kventry.Value)
			break
		}
	}
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
			log.Printf("%s Discovered %s", n.Name, kventry.Key)
			// connection not established previously
			n.SetupClient(kventry.Key, string(kventry.Value))
		}
	}
}

func (n *node) Request(requestor *node) {
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
			log.Printf("%s Discovered %s", n.Name, kventry.Key)
			// connection not established previously
			n.SetupClient(kventry.Key, string(kventry.Value))
		}
	}
}

func Start(nodeName string, port string) *node {
	log.Print("p2p Start")
	sdaddress := "localhost:8500"

	noden := node{Name: nodeName, Addr: port, SDAddress: sdaddress, Clients: nil} // noden is for opeartional purposes

	// start the node
	go noden.Start()
	return &noden
}
