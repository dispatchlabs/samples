package main

import (
	"fmt"
	"github.com/dispatchlabs/disgo/commons/types"
	"time"
	"github.com/dispatchlabs/samples/transactions/helper"
	"github.com/dispatchlabs/disgo/commons/utils"
	"github.com/dispatchlabs/samples/transactions/cli"
	"os"
	"github.com/dispatchlabs/samples/transactions/config"
	"github.com/dispatchlabs/disgo/sdk"
	"errors"
)

var delay = time.Millisecond * 20
var txCount = 1
var queueEndpoint = "/v1/queue"
var testMap map[string]time.Time
var queueTimeout = time.Second * 5


func main() {

	arg := os.Args[1]

	addressToUse := "3ed25f42484d517cdfc72cafb7ebc9e8baa52c2c"
	switch arg {
	case "setup":
		config.SetUp(5, 3500)
	case "execute", "test":
		sendGrpcTransactions(addressToUse)
		delegates, err := sdk.GetDelegates("localhost:1975")
		if err != nil {
			utils.Error(err)
		}
		//time.Sleep(time.Second * 10)
		for _, delegate := range delegates {
			account, err := sdk.GetAccount(delegate, addressToUse)
			if err != nil {
				utils.Error(err)
			}
			if account == nil {
				fmt.Printf("Account from Delegate: %s is not found yet\n", delegate.String())
			} else {
				fmt.Printf("Account from Delegate: %s is \n%s\n", delegate.String(), account.ToPrettyJson())
			}
		}
	case "balance":

		delegates, err := sdk.GetDelegates("localhost:1975")
		if err != nil {
			utils.Error(err)
		}
		for _, delegate := range delegates {
			//if index == 1 {
			//	txs, err := sdk.GetTransactionsReceived(delegate, addressToUse)
			//	if err != nil {
			//		utils.Error(err)
			//	}
			//	for _, tx := range txs {
			//		receipt, _ := sdk.GetReceipt(delegate, tx.Hash)
			//		fmt.Println(receipt.ToPrettyJson())
			//	}
			//}
			account, err := sdk.GetAccount(delegate, addressToUse)
			if err != nil {
				utils.Error(err)
				continue
			}
			fmt.Printf("Account from Delegate: %s is \n%s\n", delegate.String(), account.ToPrettyJson())
		}
	case "deployContract":
		contractAddress := deployContract()
		fmt.Printf("\nContract Address: %s\n", contractAddress)
	case "executeContract":
		//executeContract("68500f38586234a98eaa98e2b9c5adf468494c55", "multiParams")
		executeContract("cc763fe3e864e03d5786b89ec7319974209c5d3e", "arrayParam")
	case "deployAndExecute":
		contractAddress := deployContract()
		fmt.Printf("\nContract Address: %s\n", contractAddress)
		executeContract(contractAddress, "getVar5")

	default:
		fmt.Errorf("Invalid argument %s\n", arg)
	}
	//testMap = map[string]time.Time{}

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


func sendGrpcTransactions(toAddress string) {
	var tx *types.Transaction

	for i := 0; i < txCount; i++ {
		tx = helper.GetTransaction(toAddress)
		SendGrpcTransaction(tx, getRandomDelegate().GrpcEndpoint, toAddress)
		time.Sleep(delay)
	}
}

func deployContract() string {
	var tx *types.Transaction
	tx = helper.GetNewDeployTx()
	helper.PostTx(tx, getRandomDelegateURL("transactions"))
	deployHash := tx.Hash
	time.Sleep(3 * time.Second)
	deployRcpt := getReceipt(deployHash)
	return deployRcpt.ContractAddress
}

func executeContract(contractAddress string, method string) {
	var tx *types.Transaction
	tx = helper.GetNewExecuteTx(contractAddress, method)

	helper.PostTx(tx, getRandomDelegateURL("transactions"))
	time.Sleep(queueTimeout)
	//getReceipt(tx.Hash)
}


func getReceipt(hash string) *types.Receipt {
	for {
		utils.Info("Get Reciept")
		receipt := helper.GetReceipt(hash, getRandomDelegateURL("receipts"))
		fmt.Printf("Hash: %s\n%s\n", hash, receipt.ToPrettyJson())
		if receipt.Status == "Pending" {
			time.Sleep(time.Second * 5)
		} else {
			return receipt
		}
	}
}

func getRandomDelegate() types.Node {
	delegates, err := sdk.GetDelegates("localhost:1975")
	if err != nil {
		utils.Error(err)
	}
	nbrDelegates := len(delegates)
	if nbrDelegates == 0 {
		utils.Fatal(errors.New("No Delegates were returned by the seed"))
	}
	rand := utils.Random(0, nbrDelegates)
	return delegates[rand]
}

func getRandomDelegateURL(endpoint string) (string) {
	url := fmt.Sprintf("http://localhost:%d/v1/%s", getRandomDelegate().HttpEndpoint.Port, endpoint)
	return url
}

func runTransfers(toAddress string) {
	var startTime = time.Now()
	var transactions = make([]*types.Transaction, 0)
	//offset := 1000 * (txCount+1)
	//ts := utils.ToMilliSeconds(time.Now()) - int64(offset)En

	//make the transactions first.
	var tx *types.Transaction

	for i := 0; i < txCount; i++ {
		tx = helper.GetTransaction("1501d68609a7b36238c0f9a89284b4f94560ef5e")
		//tx = transfers.GetRandomTransaction()
		transactions = append(transactions, tx)
		helper.AddTx(i+1, tx)
		time.Sleep(delay*2)
	}

	types.SortByTime(transactions, false)
	for _, tx := range transactions {
		helper.PostTx(tx, getRandomDelegateURL("transactions"))
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
		tx = helper.GetTransaction(toAddress)
		helper.PostTx(tx, getRandomDelegateURL("transactions"))
		time.Sleep(delay)
	}
}

func testBroken() {
	tx := helper.GetNewBadDeployTx()
	helper.PostTx(tx, getRandomDelegateURL("transactions"))
	getReceipt(tx.Hash)
}