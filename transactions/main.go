package main

import (
	"fmt"
	"github.com/dispatchlabs/disgo/commons/types"
	"github.com/dispatchlabs/samples/transactions/transfers"
	"time"
	"github.com/dispatchlabs/samples/transactions/transfers/helper"
)

var delay = time.Second
var txCount = 5
var txEndpoint = "http://localhost:1975/v1/transactions"
var rcptEndpoint = "http://localhost:1975/v1/receipts"
var testMap map[string]time.Time
var queueTimeout = time.Second * 1

func main() {
	testMap = map[string]time.Time{}
	runTransfers()
}

func runTransfers() {
	var startTime = time.Now()
	var transactions = make([]*types.Transaction, 0)
	//offset := 1000 * (txCount+1)
	//ts := utils.ToMilliSeconds(time.Now()) - int64(offset)En

	//make the transactions first.
	var tx *types.Transaction

	for i := 0; i < txCount; i++ {
		//ts = ts + 1000
		tx = transfers.GetRandomTransaction()
		transactions = append(transactions, tx)
		helper.AddTx(i+1, tx)
		time.Sleep(delay)
	}

	types.SortByTime(transactions, false)
	for _, tx := range transactions {
		helper.PostTx(tx, txEndpoint)
		testMap[tx.Hash] = time.Now()
	}

	time.Sleep(queueTimeout)
	for k, v := range testMap {
		for {
			receipt := helper.GetReceipt(k, rcptEndpoint)
			if receipt.Status == "Pending" {
				time.Sleep(time.Second * 2)
			} else {
				helper.AddReceipt(k, receipt)
				fmt.Printf("Hash: %s :: %v\n%s\n", k, v, receipt.ToPrettyJson())
				break
			}
		}
	}

	fmt.Println(fmt.Sprintf("TXes: %d, TOTAL Time: [%v] Nanoseconds\n\n", txCount, time.Since(startTime).Nanoseconds()))
	helper.PrintTiming()
}

