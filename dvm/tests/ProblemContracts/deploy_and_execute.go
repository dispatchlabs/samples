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

var smartContractAddress = "ffa8efe11deea1448dd1d037ca6c7255d56698f6"

func main() {
	if isDeploy {
		deployContract()
	} else {
		executeContract()
	}
}

func deployContract() {
	var abi = "[{\"constant\":false,\"inputs\":[{\"name\":\"x\",\"type\":\"int256\"},{\"name\":\"y\",\"type\":\"int256\"}],\"name\":\"sdiv\",\"outputs\":[{\"name\":\"\",\"type\":\"int256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"
	var compiledCode = "608060405234801561001057600080fd5b5060c38061001f6000396000f300608060405260043610603f576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff168063397b3a49146044575b600080fd5b348015604f57600080fd5b5060766004803603810190808035906020019092919080359060200190929190505050608c565b6040518082815260200191505060405180910390f35b6000829050929150505600a165627a7a72305820a817b6c0b7eeae931c5d8bc17aaa632883d3353c6e355fa0f038287ff1ebdefc0029"

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
	var method = "sdiv"
	var params = make([]interface{}, 2)
	params[0] = -3
	params[1] = 4

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
