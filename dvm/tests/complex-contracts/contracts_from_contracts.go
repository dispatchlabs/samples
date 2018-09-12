package main

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
	"time"

	"github.com/dispatchlabs/disgo/commons/types"
	"github.com/dispatchlabs/disgo/commons/utils"
)

var nrOfServices = 0
var delegateUrl = "http://127.0.0.1:1175/v1"

var contract_deploy_Deployed = false
var contract_deploy_CallDeployed = false
var contract_execute_setA = false
var contract_execute_getA = false
var contract_execute_setAProxy = false
var contract_execute_getAProxy = true

var smartContractAddress_Deployed = "793a2bb2d0922a26ffa230626d81cbc7e7e79010"
var smartContractAddress_CallDeployed = "163308477c15c0133bb4fd57054473164d89c7e1"

func main() {
	if contract_deploy_Deployed {
		contractDeploy_Deployed()
	} else if contract_deploy_CallDeployed {
		contractDeploy_CallDeployed()
	} else if contract_execute_setA {
		contractExecute_setA()
	} else if contract_execute_getA {
		contractExecute_getA()
	} else if contract_execute_setAProxy {
		contractExecute_setAProxy()
	} else if contract_execute_getAProxy {
		contractExecute_getAProxy()
	}
}
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

	if contract_deploy_Deployed {
		contractDeploy_Deployed()
	} else if contract_deploy_CallDeployed {
		contractDeploy_CallDeployed()
	} else if contract_execute_setA {
		contractExecute_setA()
	} else if contract_execute_getA {
		contractExecute_getA()
	} else if contract_execute_setAProxy {
		contractExecute_setAProxy()
	} else if contract_execute_getAProxy {
		contractExecute_getAProxy()
	}
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
	// time.Sleep(timeout * time.Second)
	// tx := executeMethod()

	// time.Sleep(3 * time.Second)
	// executeRcpt := getReceipt(tx.Hash, fmt.Sprintf("%s/receipts", delegateUrl))

	// fmt.Println(fmt.Sprintf("CALL-RECEIPT: %v", executeRcpt))

	// time.Sleep(120 * time.Second)
	// }()
}

func deployContract(compiledCode string, abi string) *types.Transaction {
	var privateKey = "0f86ea981203b26b5b8244c8f661e30e5104555068a4bd168d3e3015db9bb25a"
	var from = "3ed25f42484d517cdfc72cafb7ebc9e8baa52c2c"
	var theTime = utils.ToMilliSeconds(time.Now())

	var tx, _ = types.NewDeployContractTransaction(
		privateKey,
		from,
		compiledCode,
		hex.EncodeToString([]byte(abi)),
		theTime,
	)

	postTx(tx, fmt.Sprintf("%s/transactions", delegateUrl))

	return tx

	// response := dapos.GetDAPoSService().Temp_ProcessTransaction(tx)
	// fmt.Println(fmt.Sprintf("%v", response))
}

func executeMethod(smartContractAddress string, method string, params []interface{}) *types.Transaction {
	// Taken from Genesis
	var privateKey = "0f86ea981203b26b5b8244c8f661e30e5104555068a4bd168d3e3015db9bb25a"
	var from = "3ed25f42484d517cdfc72cafb7ebc9e8baa52c2c"
	var to = smartContractAddress
	var theTime = utils.ToMilliSeconds(time.Now())

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

func contractDeploy_Deployed() {
	var compiledCode = "608060405234801561001057600080fd5b5060df8061001f6000396000f3006080604052600436106049576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff168063d46300fd14604e578063ee919d50146076575b600080fd5b348015605957600080fd5b50606060a0565b6040518082815260200191505060405180910390f35b348015608157600080fd5b50609e6004803603810190808035906020019092919050505060a9565b005b60008054905090565b80600081905550505600a165627a7a7230582002b6ba44c19afe847bfc64f0ecbdeeb6c51e59c7cd1617c3ab104d5ca1a90c780029"
	var abi = "[{\"constant\":true,\"inputs\":[],\"name\":\"getA\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"a\",\"type\":\"uint256\"}],\"name\":\"setA\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

	deployContract(compiledCode, abi)
}
func contractDeploy_CallDeployed() {
	var compiledCode = "608060405234801561001057600080fd5b5061025f806100206000396000f30060806040526004361061004c576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff168063c66f9daf14610051578063f474bd9e146100a8575b600080fd5b34801561005d57600080fd5b50610092600480360381019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050506100f5565b6040518082815260200191505060405180910390f35b3480156100b457600080fd5b506100f3600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803590602001909291905050506101a2565b005b6000808290508073ffffffffffffffffffffffffffffffffffffffff1663d46300fd6040518163ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401602060405180830381600087803b15801561015f57600080fd5b505af1158015610173573d6000803e3d6000fd5b505050506040513d602081101561018957600080fd5b8101908080519060200190929190505050915050919050565b60008290508073ffffffffffffffffffffffffffffffffffffffff1663ee919d50836040518263ffffffff167c010000000000000000000000000000000000000000000000000000000002815260040180828152602001915050600060405180830381600087803b15801561021657600080fd5b505af115801561022a573d6000803e3d6000fd5b505050505050505600a165627a7a7230582023420d57fb06acb2057f52d5737c6665d31d9292a5f60781e6f4d0d7774fdd3a0029"
	var abi = "[{\"constant\":true,\"inputs\":[{\"name\":\"originalContract\",\"type\":\"address\"}],\"name\":\"getAProxy\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"originalContract\",\"type\":\"address\"},{\"name\":\"a\",\"type\":\"uint256\"}],\"name\":\"setAProxy\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

	deployContract(compiledCode, abi)
}
func contractExecute_setA() {

	var smartContractAddress = smartContractAddress_Deployed
	var method = "setA"
	var params = make([]interface{}, 1)
	params[0] = 9

	executeMethod(smartContractAddress, method, params)

}
func contractExecute_getA() {
	var smartContractAddress = smartContractAddress_Deployed
	var method = "getA"
	var params = make([]interface{}, 0)

	executeMethod(smartContractAddress, method, params)
}
func contractExecute_setAProxy() {
	var smartContractAddress = smartContractAddress_CallDeployed
	var method = "setAProxy"
	var params = make([]interface{}, 2)
	params[0] = smartContractAddress_Deployed // crypto.GetAddressBytes("047cc7359e7706260b837b70e8f1d62ee972b557")
	params[1] = 20

	executeMethod(smartContractAddress, method, params)
}
func contractExecute_getAProxy() {
	var smartContractAddress = smartContractAddress_CallDeployed
	var method = "getAProxy"
	var params = make([]interface{}, 1)
	params[0] = smartContractAddress_Deployed // crypto.GetAddressBytes("047cc7359e7706260b837b70e8f1d62ee972b557")

	executeMethod(smartContractAddress, method, params)
}

func postTx(tx *types.Transaction, endpoint string) {
	fmt.Println(fmt.Sprintf("NEW-TX: %s/%s", endpoint, tx.Hash))

	data := new(bytes.Buffer)
	data.WriteString(tx.String())

	response, err := http.Post(endpoint, "application/json; charset=utf-8", data)
	if err != nil {
		utils.Error(err)
		return
	}

	contents, _ := ioutil.ReadAll(response.Body)
	fmt.Println(fmt.Sprintf("NEW-TX-Response: %s\n", string(contents)))

	// If NOT then this happens https://stackoverflow.com/questions/37454236/net-http-server-too-many-open-files-error
	response.Body.Close()
}
