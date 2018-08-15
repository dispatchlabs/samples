package main

import (
	"fmt"
	"time"

	"github.com/dispatchlabs/disgo/commons/types"
	"github.com/dispatchlabs/disgo/commons/utils"
	"github.com/dispatchlabs/samples/transactions/transfers"
	"github.com/dispatchlabs/samples/transactions/transfers/helper"
)

var delay = time.Second
var txCount = 1
var txEndpoint = "http://localhost:1175/v1/transactions"
var rcptEndpoint = "http://localhost:1175/v1/receipts"
var queueEndpoint = "http://localhost:1175/v1/queue"
var testMap map[string]time.Time
var queueTimeout = time.Second * 5

func main() {
	testMap = map[string]time.Time{}
	//runTransfers()

	// contractAddress := deployContract()
	// fmt.Println("ADDRESS: " + contractAddress)

	//executeContract(contractAddress)

	executeContract("3dd4f2bd7427ebe4dafcdf4ffad2fb31db50fbd9")
}

func deployContract() string {
	var tx *types.Transaction
	tx = transfers.GetNewDeployTx()
	time.Sleep(3 * time.Second)
	helper.PostTx(tx, txEndpoint)
	deployHash := tx.Hash
	time.Sleep(3 * time.Second)
	deployRcpt := getReceipt(deployHash)
	getReceipt(deployHash)
	return deployRcpt.ContractAddress
}

func executeContract(contractAddress string) {
	var tx *types.Transaction
	tx = transfers.GetNewExecuteTx(contractAddress)

	helper.PostTx(tx, txEndpoint)
	time.Sleep(queueTimeout)
	getReceipt(tx.Hash)
}

func getReceipt(hash string) *types.Receipt {
	for {
		utils.Info("Get Reciept")
		receipt := helper.GetReceipt(hash, rcptEndpoint)
		fmt.Printf("Hash: %s\n%s\n", hash, receipt.ToPrettyJson())
		if receipt.Status == "Pending" {
			time.Sleep(time.Second * 2)
		} else {
			return receipt
		}
	}

}

func runTransfers() {
	var startTime = time.Now()
	var transactions = make([]*types.Transaction, 0)
	//offset := 1000 * (txCount+1)
	//ts := utils.ToMilliSeconds(time.Now()) - int64(offset)En

	//make the transactions first.
	var tx *types.Transaction

	for i := 0; i < txCount; i++ {
		tx = transfers.GetTransaction(utils.ToMilliSeconds(time.Now()))
		//tx = transfers.GetRandomTransaction()
		transactions = append(transactions, tx)
		helper.AddTx(i+1, tx)
		time.Sleep(delay * 2)
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
