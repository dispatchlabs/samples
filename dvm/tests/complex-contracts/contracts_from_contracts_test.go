package complex_contracts

import (
	"bytes"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/dispatchlabs/disgo/commons/types"
	"github.com/dispatchlabs/disgo/commons/utils"
)

var nrOfServices = 0
var delegateUrl = "http://127.0.0.1:1275/v1"

func Test_ContractPublishAndExecuteSimpleContract(t *testing.T) {
	// utils.InitMainPackagePath()
	// utils.InitializeLogger()

	// utils.Events().On(services.Events.DbServiceInitFinished, allServicesInitFinished)
	// utils.Events().On(services.Events.GrpcServiceInitFinished, allServicesInitFinished)
	// utils.Events().On(services.Events.HttpServiceInitFinished, allServicesInitFinished)

	// utils.Events().On(disgover.Events.DisGoverServiceInitFinished, allServicesInitFinished)
	// utils.Events().On(dapos.Events.DAPoSServiceInitFinished, allServicesInitFinished)
	// utils.Events().On(dvm.Events.DVMServiceInitFinished, allServicesInitFinished)

	// utils.Info(fmt.Sprintf("NR of services left to be started: %d", nrOfServices))

	// server := bootstrap.NewServer()
	// server.Go()

	// allServicesInitFinished()

	// deployContract()
	executeMethod()
}

func allServicesInitFinished() {
	nrOfServices--
	utils.Info(fmt.Sprintf("NR of services left to be started: %d", nrOfServices))

	if nrOfServices > 0 {
		return
	}

	const timeout = 1

	// go func() {
	// time.Sleep(timeout * time.Second)
	// tx := deployContract()

	// time.Sleep(3 * time.Second)
	// deployRcpt := getReceipt(tx.Hash, fmt.Sprintf("%s/receipts", delegateUrl))

	// fmt.Println(fmt.Sprintf("CONTRACT-ADDRESS: %s", deployRcpt.ContractAddress))
	// }()

	// ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~

	// go func() {
	time.Sleep(timeout * time.Second)
	tx := executeMethod()

	time.Sleep(3 * time.Second)
	executeRcpt := getReceipt(tx.Hash, fmt.Sprintf("%s/receipts", delegateUrl))

	fmt.Println(fmt.Sprintf("CALL-RECEIPT: %v", executeRcpt))

	time.Sleep(120 * time.Second)
	// }()
}

func deployContract() *types.Transaction {
	var privateKey = "0f86ea981203b26b5b8244c8f661e30e5104555068a4bd168d3e3015db9bb25a"
	var from = "3ed25f42484d517cdfc72cafb7ebc9e8baa52c2c"
	var code = getCompiledCode()
	var theTime = utils.ToMilliSeconds(time.Now())

	var tx, _ = types.NewDeployContractTransaction(
		privateKey,
		from,
		code,
		hex.EncodeToString([]byte(getAbi())),
		theTime,
	)

	postTx(tx, fmt.Sprintf("%s/transactions", delegateUrl))

	return tx

	// response := dapos.GetDAPoSService().Temp_ProcessTransaction(tx)
	// fmt.Println(fmt.Sprintf("%v", response))
}

func executeMethod() *types.Transaction {
	// Taken from Genesis
	var privateKey = "0f86ea981203b26b5b8244c8f661e30e5104555068a4bd168d3e3015db9bb25a"
	var from = "3ed25f42484d517cdfc72cafb7ebc9e8baa52c2c"
	var to = "da8c13fc708d85c8fd2b5d5975e7300faae90a92"

	var theTime = utils.ToMilliSeconds(time.Now())

	// var method = "getA"
	// var params = make([]interface{}, 0)
	// params[0] = 1

	// var method = "setA"
	// var params = make([]interface{}, 1)
	// params[0] = 54321

	// var method = "setAProxy"
	// var params = make([]interface{}, 2)
	// params[0] = "d914892762206c3f34a16ce6efb276026a89e820" // crypto.GetAddressBytes("047cc7359e7706260b837b70e8f1d62ee972b557")
	// params[1] = 20

	var method = "getAProxy"
	var params = make([]interface{}, 1)
	params[0] = "d914892762206c3f34a16ce6efb276026a89e820" // crypto.GetAddressBytes("047cc7359e7706260b837b70e8f1d62ee972b557")

	var tx, _ = types.NewExecuteContractTransaction(
		privateKey,
		from,
		to,
		method,
		params,
		theTime,
	)

	postTx(tx, fmt.Sprintf("%s/transactions", delegateUrl))

	return tx

	// dapos.GetDAPoSService().Temp_ProcessTransaction(tx)
}

func getCompiledCode() string {
	return "608060405234801561001057600080fd5b5061025f806100206000396000f30060806040526004361061004c576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff168063c66f9daf14610051578063f474bd9e146100a8575b600080fd5b34801561005d57600080fd5b50610092600480360381019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050506100f5565b6040518082815260200191505060405180910390f35b3480156100b457600080fd5b506100f3600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803590602001909291905050506101a2565b005b6000808290508073ffffffffffffffffffffffffffffffffffffffff1663d46300fd6040518163ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401602060405180830381600087803b15801561015f57600080fd5b505af1158015610173573d6000803e3d6000fd5b505050506040513d602081101561018957600080fd5b8101908080519060200190929190505050915050919050565b60008290508073ffffffffffffffffffffffffffffffffffffffff1663ee919d50836040518263ffffffff167c010000000000000000000000000000000000000000000000000000000002815260040180828152602001915050600060405180830381600087803b15801561021657600080fd5b505af115801561022a573d6000803e3d6000fd5b505050505050505600a165627a7a7230582095716c21ef28aca7f84be244bc34b9a6c4e292d75d3067060f6fd2a486a1bd640029"
}

func getAbi() string {
	return `[
		{
			"constant": true,
			"inputs": [
				{
					"name": "originalContract",
					"type": "address"
				}
			],
			"name": "getAProxy",
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
					"name": "originalContract",
					"type": "address"
				},
				{
					"name": "a",
					"type": "uint256"
				}
			],
			"name": "setAProxy",
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
