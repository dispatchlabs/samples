package main

import (
	"fmt"
	"github.com/dispatchlabs/disgo/commons/types"
	"github.com/dispatchlabs/samples/transactions/transfers"
	"time"
	"github.com/dispatchlabs/samples/transactions/transfers/helper"
	"github.com/dispatchlabs/disgo/commons/utils"
)

var delay = time.Millisecond * 20
var txCount = 5000
var txEndpoint = "http://localhost:1575/v1/transactions"
var rcptEndpoint = "http://localhost:1575/v1/receipts"
var queueEndpoint = "http://localhost:1575/v1/queue"
var testMap map[string]time.Time
var queueTimeout = time.Second * 5

func main() {
	testMap = map[string]time.Time{}

	//TransferTest()
	//testBroken()
	//runTransfers()

	//contractAddress := deployContract()
	//fmt.Printf("\nContract Address: %s\n", contractAddress)
	//executeContract(contractAddress, "arrayParam")

	//executeContract("7604548fef43108f9c00e3f9b3979f92772a63b6", "arrayParam")

	sendGrpcTransactions()
}

func sendGrpcTransactions() {
	var tx *types.Transaction

	for i := 0; i < txCount; i++ {
		tx = transfers.GetTransaction()
		SendGrpcTransaction(tx)
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

func runTransfers() {
	var startTime = time.Now()
	var transactions = make([]*types.Transaction, 0)
	//offset := 1000 * (txCount+1)
	//ts := utils.ToMilliSeconds(time.Now()) - int64(offset)En

	//make the transactions first.
	var tx *types.Transaction

	for i := 0; i < txCount; i++ {
		tx = transfers.GetTransaction()
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

func TransferTest() {
	var tx *types.Transaction

	for i := 0; i < txCount; i++ {
		tx = transfers.GetTransaction()
		helper.PostTx(tx, txEndpoint)
		time.Sleep(delay)
	}
}

func testBroken() {
	tx := transfers.GetNewBadDeployTx()
	helper.PostTx(tx, txEndpoint)
	getReceipt(tx.Hash)
}