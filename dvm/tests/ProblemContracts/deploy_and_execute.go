package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/dispatchlabs/disgo/commons/types"
	"github.com/dispatchlabs/disgo/commons/utils"
)

var nrOfServices = 0
var delegateUrl = "http://127.0.0.1:1175/v1"

var isDeploy = false

var smartContractAddress = "0a7f4c1abd1b31a00cfb1bb759709b187dd37084"

func main() {
	if isDeploy {
		deployContract()
	} else {
		executeContract()
	}
}

func deployContract() {
	var compiledCode = "6080604052348015600f57600080fd5b5060a48061001e6000396000f300608060405260043610603f576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff168063d2756dc0146044575b600080fd5b348015604f57600080fd5b506056606c565b6040518082815260200191505060405180910390f35b60004260005260206000f300a165627a7a72305820ee5cc8e4e4eb6f1cdc13bcfb018a314b1d89603f28abaf1ec47cbd2d19df08120029"
	var abi = "[{\"constant\":true,\"inputs\":[],\"name\":\"test_timestamp\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

	var privateKey = "0f86ea981203b26b5b8244c8f661e30e5104555068a4bd168d3e3015db9bb25a"
	var from = "3ed25f42484d517cdfc72cafb7ebc9e8baa52c2c"
	var theTime = utils.ToMilliSeconds(time.Now())

	var tx, _ = types.NewDeployContractTransaction(
		privateKey,
		from,
		compiledCode,
		abi,
		theTime,
	)

	postTx(tx, fmt.Sprintf("%s/transactions", delegateUrl))
}

func executeContract() {
	var method = "test_timestamp"
	var params = make([]interface{}, 0)

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
