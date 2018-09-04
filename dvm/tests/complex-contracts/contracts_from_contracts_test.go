package complex_contracts

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/dispatchlabs/disgo/bootstrap"
	"github.com/dispatchlabs/disgo/commons/services"
	"github.com/dispatchlabs/disgo/commons/types"
	"github.com/dispatchlabs/disgo/commons/utils"
	"github.com/dispatchlabs/disgo/dapos"
	"github.com/dispatchlabs/disgo/disgover"
	"github.com/dispatchlabs/disgo/dvm"
)

var nrOfServices = 6
var delegateUrl = "http://127.0.0.1:1175/v1"

func Test_ContractPublishAndExecuteSimpleContract(t *testing.T) {
	utils.InitMainPackagePath()
	utils.InitializeLogger()

	utils.Events().On(services.Events.DbServiceInitFinished, allServicesInitFinished)
	utils.Events().On(services.Events.GrpcServiceInitFinished, allServicesInitFinished)
	utils.Events().On(services.Events.HttpServiceInitFinished, allServicesInitFinished)

	utils.Events().On(disgover.Events.DisGoverServiceInitFinished, allServicesInitFinished)
	utils.Events().On(dapos.Events.DAPoSServiceInitFinished, allServicesInitFinished)
	utils.Events().On(dvm.Events.DVMServiceInitFinished, allServicesInitFinished)

	utils.Info(fmt.Sprintf("NR of services left to be started: %d", nrOfServices))

	server := bootstrap.NewServer()
	server.Go()
}

func allServicesInitFinished() {
	nrOfServices--
	utils.Info(fmt.Sprintf("NR of services left to be started: %d", nrOfServices))

	if nrOfServices > 0 {
		return
	}

	const timeout = 1

	go func() {
		time.Sleep(timeout * time.Second)
		tx := deployContract()

		time.Sleep(3 * time.Second)
		deployRcpt := getReceipt(tx.Hash, fmt.Sprintf("%s/receipts", delegateUrl))

		fmt.Println(fmt.Sprintf("CONTRACT-ADDRESSP: %s", deployRcpt.ContractAddress))
	}()

	//go func() {
	//	time.Sleep(timeout * time.Second)
	//	executeMethod_set()
	//}()

	// go func() {
	// 	time.Sleep(timeout * time.Second)
	// 	executeMethod_get()
	// }()
}

func deployContract() *types.Transaction {
	var privateKey = "0f86ea981203b26b5b8244c8f661e30e5104555068a4bd168d3e3015db9bb25a"
	var from = "3ed25f42484d517cdfc72cafb7ebc9e8baa52c2c"
	var code = "608060405234801561001057600080fd5b5060df8061001f6000396000f3006080604052600436106049576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff168063d46300fd14604e578063ee919d50146076575b600080fd5b348015605957600080fd5b50606060a0565b6040518082815260200191505060405180910390f35b348015608157600080fd5b50609e6004803603810190808035906020019092919050505060a9565b005b60008054905090565b80600081905550505600a165627a7a723058205906547745a52855a1b22685e079cbdec04bad5d24c4c243d60837b39fb845890029"
	var abi = getAbi()
	var theTime = utils.ToMilliSeconds(time.Now())

	var tx, _ = types.NewDeployContractTransaction(
		privateKey,
		from,
		code,
		abi,
		theTime,
	)

	postTx(tx, fmt.Sprintf("%s/transactions", delegateUrl))

	return tx

	// response := dapos.GetDAPoSService().Temp_ProcessTransaction(tx)
	// fmt.Println(fmt.Sprintf("%v", response))
}

func executeMethod() {
	// Taken from Genesis
	var privateKey = "0f86ea981203b26b5b8244c8f661e30e5104555068a4bd168d3e3015db9bb25a"
	var from = "3ed25f42484d517cdfc72cafb7ebc9e8baa52c2c"
	var to = "78337c25f0c003344c1b16e5f4b5ebda07a08cf5"

	var theTime = utils.ToMilliSeconds(time.Now())
	var method = "getVar5"
	var params = make([]interface{}, 0)
	// var params = make([]interface{}, 1)
	// params[0] = "5555"

	var tx, _ = types.NewExecuteContractTransaction(
		privateKey,
		from,
		to,
		// hex.EncodeToString([]byte(abi)),
		method,
		params,
		theTime,
	)

	dapos.GetDAPoSService().Temp_ProcessTransaction(tx)
}

func getAbi() string {
	return `[
		{
			"constant": true,
			"inputs": [],
			"name": "getA",
			"outputs": [
				{
					"name": "",
					"type": "uint256"
				}
			],
			"payable": false,
			"stateMutability": "view",
			"type": "function"
		},
		{
			"constant": false,
			"inputs": [
				{
					"name": "a",
					"type": "uint256"
				}
			],
			"name": "setA",
			"outputs": [],
			"payable": false,
			"stateMutability": "nonpayable",
			"type": "function"
		}
	]`
}

func postTx(tx *types.Transaction, endpoint string) {
	fmt.Printf("Executing contract json: \n%s\n", tx.ToPrettyJson())
	fmt.Printf("Sending tx : %s with timestamp: %v\n", tx.Hash, tx.ToTime())
	data := new(bytes.Buffer)
	data.WriteString(tx.String())

	response, err := http.Post(
		endpoint,
		"application/json; charset=utf-8",
		data,
	)
	if err != nil {
		utils.Error(err)
		return
	}
	contents, _ := ioutil.ReadAll(response.Body)
	// If NOT then this happens https://stackoverflow.com/questions/37454236/net-http-server-too-many-open-files-error
	fmt.Printf("Response: %v\n", string(contents))
	response.Body.Close()
}

func getReceipt(hash string, endpoint string) *types.Receipt {
	response, err := http.Get(fmt.Sprintf("%s/%s", endpoint, hash))
	var receipt *types.Receipt
	if err != nil {
		fmt.Printf("%s", err)
		os.Exit(1)
	} else {
		defer response.Body.Close()
		contents, _ := ioutil.ReadAll(response.Body)
		receipt, err = unmarshalReceipt(contents)
		if err != nil {
			fmt.Printf("%s", err)
			os.Exit(1)
		}
		//fmt.Printf("%s\n", receipt.ToPrettyJson())
	}
	return receipt
}

func unmarshalReceipt(bytes []byte) (*types.Receipt, error) {
	receipt := types.Receipt{}
	var jsonMap map[string]interface{}
	err := json.Unmarshal(bytes, &jsonMap)
	if err != nil {
		utils.Fatal(err)
	}
	if jsonMap["data"] != nil {

		value := jsonMap["data"].(map[string]interface{})
		if value != nil {
			if value["transactionHash"] != nil {
				receipt.TransactionHash = value["transactionHash"].(string)
			}
			if value["status"] != nil {
				receipt.Status = value["status"].(string)
			}
			if value["humanReadableStatus"] != nil {
				receipt.HumanReadableStatus = value["humanReadableStatus"].(string)
			}
			if value["contractAddress"] != nil && value["contractAddress"] != "" {
				receipt.ContractAddress = value["contractAddress"].(string)
			}
			if value["contractResult"] != nil {
				var contractResult = value["contractResult"]
				receipt.ContractResult = contractResult.([]interface{})
			}
			if value["created"] != nil {
				created, err := time.Parse(time.RFC3339, value["created"].(string))
				if err != nil {
					return nil, err
				}
				receipt.Created = created
			}
		}
	}
	return &receipt, nil
}
