package main

import (
	"encoding/json"
	"log"
	"net/http"
	"net/url"
	"os"
	"time"

	"github.com/nic0lae/JerryMouse/Servers"
)

type Block struct {
	PreviousHash string
	Timestamp    int64
	Data         string
	Hash         string
}

var blockchain = []Block{GetGenesisBlock()}

var apiServer = Servers.Api()

func handleRequest_Blocks(rw http.ResponseWriter, r *http.Request) {
	json.NewEncoder(rw).Encode(blockchain)
}

func handleRequest_MineBlock(data interface{}) Servers.JsonResponse {
	dataAsJson, ok := data.(*Block)
	if !ok {
		return Servers.JsonResponse{Error: "Invalid Params"}
	}

	newBlock := GenerateNextBlock(blockchain, dataAsJson.Data)
	blockchain = AddBlock(blockchain, newBlock)
	apiServer.SendToAllRealtimePeers(ResponseLatestMsg(blockchain))

	return Servers.JsonResponse{}
}

func handleRequest_NodeToNode(inChannel chan []byte, outChannel chan []byte) {
	go func() {
		for {
			data, ok := <-inChannel
			if !ok {
				close(outChannel)
				break
			} else {
				var newBlockchain []Block
				// var jsonData Servers.JsonData
				err := json.Unmarshal(data, &blockchain)
				if err != nil {
					log.Print("handleRequest_NodeToNode: IN -> Unmarshal", err)
				} else {
					blockchain = newBlockchain
				}
			}
		}
	}()
}

func issueRequest_NodeToNode(nodeIpPort string) {
	var realtimeClient = Servers.NewRealtimeClient()
	realtimeClient.ConnectToPeer(
		url.URL{Scheme: "ws", Host: nodeIpPort, Path: "/n2n"},
		&Servers.RealtimeHandler{
			Handler: handleRequest_NodeToNode,
		},
	)
}

func main() {
	apiServer.SetLowLevelHandlers([]Servers.LowLevelHandler{
		Servers.LowLevelHandler{
			Route:   "/blocks",
			Handler: handleRequest_Blocks,
			Verb:    "GET",
		},
	})
	apiServer.SetJsonHandlers([]Servers.JsonHandler{
		Servers.JsonHandler{
			Route:      "/mineBlock",
			Handler:    handleRequest_MineBlock,
			JsonObject: &Block{},
		},
	})
	apiServer.SetRealtimeHandlers([]Servers.RealtimeHandler{
		Servers.RealtimeHandler{
			Route:   "/n2n",
			Handler: handleRequest_NodeToNode,
		},
	})

	if len(os.Args) > 2 {
		go func() {
			time.Sleep(3 * time.Second)
			issueRequest_NodeToNode(os.Args[2])
		}()
	}

	apiServer.Run(os.Args[1])
}
