package main

import (
	"errors"
	"fmt"
	"github.com/dispatchlabs/disgo/commons/types"
	"github.com/dispatchlabs/disgo/commons/utils"
	"github.com/dispatchlabs/disgo/sdk"
	"github.com/dispatchlabs/samples/common-util/config"
	"github.com/dispatchlabs/samples/transactions/helper"
	"os"
	"time"
	"io/ioutil"
)

var privateKey 	= "0f86ea981203b26b5b8244c8f661e30e5104555068a4bd168d3e3015db9bb25a"
var from 		= "3ed25f42484d517cdfc72cafb7ebc9e8baa52c2c"

var delay = time.Millisecond * 2
var txCount = 1
var queueEndpoint = "/v1/queue"
var testMap map[string]time.Time
var queueTimeout = time.Second * 5

func main() {

	arg := os.Args[1]

	addressToUse := "d7a6acf5f89cf2ca4d618b3a5aeeb3d3ef4e0574"
	switch arg {
	case "setup":
		//config.SetUp(5, 3500)
	case "update":
		config.RefreshDisgoExecutable("")
	case "execute", "test":
		//transaction := sendGrpcTransactions(addressToUse)
		hashes := sendHttpTransactions(addressToUse)

		delegates, err := sdk.GetDelegates("localhost:1975")
		if err != nil {
			utils.Fatal(err)
		}
		time.Sleep(time.Second * 1)
		for _, delegate := range delegates {
			tx, err := sdk.GetTransaction(delegate, hashes[0])
			if err != nil {
				utils.Error("Error getting Transaction ", err)
			}
			if tx == nil {
				fmt.Printf("Transaction from Delegate: %s is not found yet\n", delegate.String())
			} else {
				fmt.Printf("Transaction from Delegate: %s is \n%s\n", delegate.String(), tx.ToPrettyJson())
			}
		}
	case "executeHttp":
		sendHttpTransactions(addressToUse)
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
	case "deployContractFromFile":
		contractAddress := deployContractFromFile(os.Args[2:])
		fmt.Printf("\nContract Address: %s\n", contractAddress)
	case "executeContract":
		//executeContract("68500f38586234a98eaa98e2b9c5adf468494c55", "multiParams")
		//executeContract("f8e84ac2f4d70fbb84d9d33bac70e4da809ae29c", "hi")
		executeContract("9b4a6a6d26916bae7bfde23ac0fc6cf6e1273d4f", "intParams")
	case "executeVarArgContract":
		if len(os.Args) < 4 {
			fmt.Println("executeVarArgContract must have at least 3 arguments\n")
			break
		}
		executeVarArgContract(os.Args[2], os.Args[3], os.Args[4:])
	case "deployAndExecute":
		contractAddress := deployContract()
		fmt.Printf("\nContract Address: %s\n", contractAddress)
		executeContract(contractAddress, "intParam")
		//NewManualExecuteTx(contractAddress, "intParam")
	default:
		fmt.Errorf("Invalid argument %s\n", arg)
	}
	//testMap = map[string]time.Time{}

}

//func sendGrpcTransactions(toAddress string) *types.Transaction {
//	var tx *types.Transaction
//
//	for i := 0; i < txCount; i++ {
//		tx = helper.GetTransaction(toAddress)
//		gossipResponse, err := SendGrpcTransaction(tx, getRandomDelegate().GrpcEndpoint, toAddress)
//		if err != nil {
//			utils.Error(err)
//		} else {
//			//fmt.Printf("grpc response: %v\n", gossipResponse)
//			fmt.Printf("Transaction Hash: %v\n", gossipResponse.Transaction.Hash)
//		}
//		time.Sleep(delay)
//	}
//	return tx
//}

func sendHttpTransactions(toAddress string) []string {
	hashes := make([]string, txCount)
	var err error
	for i := 0; i < txCount; i++ {
		hashes[i], err = sdk.TransferTokens(getRandomDelegate(), privateKey, from, toAddress, 1)
		if err != nil {
			utils.Error(err)
		}
	}
	return hashes;
}

func deployContract() string {
	deployHash, err := sdk.DeploySmartContract(
		getRandomDelegate(),
		privateKey,
		from,
		helper.GetCode(),
		helper.GetAbi())
	if err != nil {
		utils.Error(err)
	}
	time.Sleep(3 * time.Second)
	deployRcpt := getReceipt(deployHash)
	return deployRcpt.ContractAddress
}

func deployContractFromFile(args []string) string {
	if len(args) != 2 {
		fmt.Println("deployContractFromFile needs a binary file (arg 1) and abi file (arg 2)")
		return ""
	}
	binary, err := ioutil.ReadFile(args[0])
	if err != nil {
		panic(err)
	}

	abi, err := ioutil.ReadFile(args[1])
	if err != nil {
		panic(err)
	}
	deployHash, err := sdk.DeploySmartContract(getRandomDelegate(), privateKey, from, string(binary), string(abi))
	if err != nil {
		utils.Error(err)
	}
	time.Sleep(3 * time.Second)
	deployRcpt := getReceipt(deployHash)
	return deployRcpt.ContractAddress
}

func executeContract(contractAddress string, method string) string {
	hash, err := sdk.ExecuteSmartContractTransaction(getRandomDelegate(), privateKey, from, contractAddress, method, helper.GetParamsForMethod(method))
	if err != nil {
		utils.Error(err)
	}
	fmt.Printf("Hash: %s", hash)
	//time.Sleep(3 * time.Second)
	//getReceipt(hash)
	return hash
}

func NewManualExecuteTx(toAddress string, method string) *types.Transaction {
	utils.Info("GetNewExecuteTx")
	// Taken from Genesis
	var privateKey = "0f86ea981203b26b5b8244c8f661e30e5104555068a4bd168d3e3015db9bb25a"
	var from = "3ed25f42484d517cdfc72cafb7ebc9e8baa52c2c"

	var tx, _ = types.NewExecuteContractTransaction(
		privateKey,
		from,
		toAddress,
		method,
		helper.GetParamsForMethod(method),
		utils.ToMilliSeconds(time.Now()),
	)
	//tx.Abi = helper.GetAbi()
	tx.Value = 0
	tx.Code = ""
	tx.FromName = ""
	tx.ToName = ""
	helper.PostTx(tx, "http://localhost:1575/v1/transactions")

	time.Sleep(queueTimeout)
	return tx

}

func executeVarArgContract(contractAddress string, method string, args []string) {
	fmt.Println(contractAddress)
	fmt.Println(method)
	fmt.Println(args)

	var tx *types.Transaction
	tx = helper.GetNewExecuteTxWithVarableParams(contractAddress, method, args)

	helper.PostTx(tx, getRandomDelegateURL("transactions"))
	time.Sleep(queueTimeout)
	getReceipt(tx.Hash)
}

func getReceipt(hash string) types.Receipt {
	for {
		utils.Info("Get Reciept")
		tx, err := sdk.GetTransaction(getRandomDelegate(), hash)
		if err != nil {
			utils.Error(err)
		}
		fmt.Printf(tx.String())
		receipt := tx.Receipt
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


func getRandomDelegateURL(endpoint string) string {
	url := fmt.Sprintf("http://localhost:%d/v1/%s", getRandomDelegate().HttpEndpoint.Port, endpoint)
	//url := fmt.Sprintf("http://35.203.143.69:%d/v1/%s", getRandomDelegate().HttpEndpoint.Port, endpoint)
	// url := fmt.Sprintf("http://35.233.231.3:%d/v1/%s", 1975, endpoint)
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
		tx = helper.GetTransaction(toAddress)
		transactions = append(transactions, tx)
		helper.AddTx(i+1, tx)
		time.Sleep(delay * 2)
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
		helper.AddReceipt(k, &receipt)
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

//func testBroken() {
//	tx := helper.GetNewBadDeployTx()
//	helper.PostTx(tx, getRandomDelegateURL("transactions"))
//	getReceipt(tx.Hash)
//}
