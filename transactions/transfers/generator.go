package transfers

import (
	"encoding/hex"
	"fmt"
	"time"

	"github.com/dispatchlabs/disgo/commons/types"
	"github.com/dispatchlabs/disgo/commons/utils"
)

var deployOccurred = false

func GetRandomTransaction() *types.Transaction {
	value := utils.Random(1, 3)
	timestamp := utils.ToMilliSeconds(time.Now())
	var contractAddress string
	switch value {
	case 1:
		return GetTransaction(timestamp)
	case 2:
		//deployOccurred = true
		return GetNewDeployTx()
	case 3:
		if deployOccurred {
			return GetNewExecuteTx(contractAddress)
		} else {
			return GetNewDeployTx()
		}
	}
	return nil
}

func GetTransaction(timestamp int64) *types.Transaction {
	utils.Info("GetTransaction")
	var privateKey = "0f86ea981203b26b5b8244c8f661e30e5104555068a4bd168d3e3015db9bb25a"
	var from = "3ed25f42484d517cdfc72cafb7ebc9e8baa52c2c"

	var tx, _ = types.NewTransferTokensTransaction(
		privateKey,
		from,
		"d5765c93699c96327753230ac3d78edb3b34236b",
		1,
		1,
		timestamp,
	)
	fmt.Printf("Created Tx: %s\n", tx.Hash)

	return tx
}

func GetNewDeployTx() *types.Transaction {
	utils.Info("GetNewDeployTx")

	var privateKey = "0f86ea981203b26b5b8244c8f661e30e5104555068a4bd168d3e3015db9bb25a"
	var from = "3ed25f42484d517cdfc72cafb7ebc9e8baa52c2c"

	var tx, _ = types.NewDeployContractTransaction(
		privateKey,
		from,
		getCode(),
		getAbi(),
		utils.ToMilliSeconds(time.Now()),
	)
	//
	fmt.Printf("DEPLOY: %s\n", tx.ToPrettyJson())
	return tx
}

func GetNewExecuteTx(toAddress string) *types.Transaction {
	utils.Info("GetNewExecuteTx")
	// Taken from Genesis
	var privateKey = "0f86ea981203b26b5b8244c8f661e30e5104555068a4bd168d3e3015db9bb25a"
	var from = "3ed25f42484d517cdfc72cafb7ebc9e8baa52c2c"

	var method = "plusOne"
	//var params = make([]interface{}, 0)
	// var array = make([]interface{}, 1)
	// array[0] = uint(20)

	var params = make([]interface{}, 1)
	params[0] = uint(20)

	var tx, _ = types.NewExecuteContractTransaction(
		privateKey,
		from,
		toAddress,
		hex.EncodeToString([]byte(getAbi())),
		method,
		params,
		utils.ToMilliSeconds(time.Now()),
	)

	return tx

}

func getCode() string {
	return "608060405234801561001057600080fd5b5060bb8061001f6000396000f300608060405260043610603f576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff168063f5a6259f146044575b600080fd5b348015604f57600080fd5b50606c600480360381019080803590602001909291905050506082565b6040518082815260200191505060405180910390f35b60006001820190509190505600a165627a7a723058205f49aaa50be01260d346aba22a9173715a9e7e2c0b9a75790babd87316b9dbcd0029"
}

func getAbi() string {

	return `[
		{
			"constant": true,
			"inputs": [
				{
					"name": "y",
					"type": "uint256"
				}
			],
			"name": "plusOne",
			"outputs": [
				{
					"name": "x",
					"type": "uint256"
				}
			],
			"payable": false,
			"stateMutability": "pure",
			"type": "function"
		}
	]`
}
