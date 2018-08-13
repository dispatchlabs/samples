package main

import (
	"fmt"
	"github.com/dispatchlabs/disgo/commons/types"
	"github.com/dispatchlabs/samples/transactions/transfers"
	"time"
	"github.com/dispatchlabs/samples/transactions/transfers/helper"
	"github.com/dispatchlabs/disgo/commons/utils"
	"github.com/dispatchlabs/samples/transactions/cli"
	"os"
	"github.com/dispatchlabs/samples/transactions/config"
	"github.com/dispatchlabs/disgo/sdk"
)

var delay = time.Millisecond * 20
var txCount = 100
var txEndpoint = "http://localhost:1575/v1/transactions"
var rcptEndpoint = "http://localhost:1575/v1/receipts"
var queueEndpoint = "http://localhost:1575/v1/queue"
var testMap map[string]time.Time
var queueTimeout = time.Second * 5

func main() {

	arg := os.Args[1]
	addressToUse := "d909e9b6e8909943a7a2783581451021584fdf11"
	switch arg {
	case "setup":
		config.SetUp()
	case "execute":
		sendGrpcTransactions(addressToUse, 1173)
		delegates, err := sdk.GetDelegates("localhost:1975")
		if err != nil {
			utils.Error(err)
		}
		for _, delegate := range delegates {
			account, err := sdk.GetAccount(delegate, addressToUse)
			if err != nil {
				utils.Error(err)
			}
			fmt.Printf("Account from Delegate: %s is \n%s\n", delegate.String(), account.ToPrettyJson())
		}
	case "balance":
		delegates, err := sdk.GetDelegates("localhost:1975")
		if err != nil {
			utils.Error(err)
		}
		for _, delegate := range delegates {
			account, err := sdk.GetAccount(delegate, addressToUse)
			if err != nil {
				utils.Error(err)
			}
			fmt.Printf("Account from Delegate: %s is \n%s\n", delegate.String(), account.ToPrettyJson())
		}
	default:
		fmt.Errorf("Invalid argument %s\n", arg)
	}
	//testMap = map[string]time.Time{}

	//TransferTest()
	//testBroken()
	//runTransfers()

	//contractAddress := deployContract()
	//fmt.Printf("\nContract Address: %s\n", contractAddress)
	//executeContract(contractAddress, "arrayParam")

	//executeContract("e6645f99688061161086cc2d442fa5ca51d9dc83", "arrayParam")
}



func Startup() {
	cli.Exec("cd /Users/Bob/go/src/github.com/dispatchlabs/samples/run-nodes-locally/seed; ls -al; ./disgo")
	//time.Sleep(time.Second * 3)
	//go cli.Exec("cd /Users/Bob/go/src/github.com/dispatchlabs/samples/run-nodes-locally/delegat-1; ls -al; ./disgo")
	//go cli.Exec("cd /Users/Bob/go/src/github.com/dispatchlabs/samples/run-nodes-locally/delegat-2; ls -al; ./disgo")
	//go cli.Exec("cd /Users/Bob/go/src/github.com/dispatchlabs/samples/run-nodes-locally/delegat-3; ls -al; ./disgo")
	//go cli.Exec("cd /Users/Bob/go/src/github.com/dispatchlabs/samples/run-nodes-locally/delegat-4; ls -al; ./disgo")
	//time.Sleep(time.Minute)
}


func sendGrpcTransactions(toAddress string, grpcPort int64) {
	var tx *types.Transaction

	for i := 0; i < txCount; i++ {
		tx = transfers.GetTransaction(toAddress)
		SendGrpcTransaction(tx, grpcPort, toAddress)
		time.Sleep(delay)
	}
}

func deployContract() string {
	var tx *types.Transaction
	tx = transfers.GetNewDeployTx()
	helper.PostTx(tx, txEndpoint)
	deployHash := tx.Hash
	time.Sleep(3 * time.Second)
	deployRcpt := getReceipt(deployHash)
	return deployRcpt.ContractAddress
}

func executeContract(contractAddress string, method string) {
	var tx *types.Transaction
	tx = transfers.GetNewExecuteTx(contractAddress, method)

	helper.PostTx(tx, txEndpoint)
	time.Sleep(queueTimeout)
	//getReceipt(tx.Hash)
}


func getReceipt(hash string) *types.Receipt {
	for {
		utils.Info("Get Reciept")
		receipt := helper.GetReceipt(hash, rcptEndpoint)
		fmt.Printf("Hash: %s\n%s\n", hash, receipt.ToPrettyJson())
		if receipt.Status == "Pending" {
			time.Sleep(time.Second * 5)
		} else {
			return receipt
		}
	}
}

func runTransfers(toAddress string) {
	var startTime = time.Now()
	var transactions = make([]*types.Transaction, 0)
	//offset := 1000 * (txCount+1)
	//ts := utils.ToMilliSeconds(time.Now()) - int64(offset)En

	//make the transactions first.
	var tx *types.Transaction

	for i := 0; i < txCount; i++ {
		tx = transfers.GetTransaction("1501d68609a7b36238c0f9a89284b4f94560ef5e")
		//tx = transfers.GetRandomTransaction()
		transactions = append(transactions, tx)
		helper.AddTx(i+1, tx)
		time.Sleep(delay*2)
	}

	types.SortByTime(transactions, false)
	for _, tx := range transactions {
		helper.PostTx(tx, txEndpoint)
		testMap[tx.Hash] = time.Now()
	}
	time.Sleep(time.Second)
	fmt.Printf("QUEUE DUMP: \n%s\n", helper.GetQueue(queueEndpoint))
	time.Sleep(queueTimeout)
	for k, _ := range testMap {
		receipt := getReceipt(k)
		helper.AddReceipt(k, receipt)
	}

	fmt.Println(fmt.Sprintf("TXes: %d, TOTAL Time: [%v] Nanoseconds\n\n", txCount, time.Since(startTime).Nanoseconds()))
	helper.PrintTiming()
}

func TransferTest(toAddress string) {
	var tx *types.Transaction

	for i := 0; i < txCount; i++ {
		tx = transfers.GetTransaction(toAddress)
		helper.PostTx(tx, txEndpoint)
		time.Sleep(delay)
	}
}

func testBroken() {
	tx := transfers.GetNewBadDeployTx()
	helper.PostTx(tx, txEndpoint)
	getReceipt(tx.Hash)
}