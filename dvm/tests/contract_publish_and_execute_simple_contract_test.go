package tests

import (
	"encoding/hex"
	"fmt"
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

// ContractPublishAndExecuteSimpleContract -
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

	const timeout = 3

	// go func() {
	// 	time.Sleep(timeout * time.Second)
	// 	deployContract()
	// }()

	go func() {
		time.Sleep(timeout * time.Second)
		executeMethod_1()
	}()

	// go func() {
	// 	time.Sleep(timeout * time.Second)
	// 	executeMethod_getVar5()
	// }()
}

func deployContract() {
	var privateKey = "0f86ea981203b26b5b8244c8f661e30e5104555068a4bd168d3e3015db9bb25a"
	var from = "3ed25f42484d517cdfc72cafb7ebc9e8baa52c2c"
	var code = "608060405234801561001057600080fd5b50610162806100206000396000f300608060405260043610610041576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff1680634a846e0214610046575b600080fd5b34801561005257600080fd5b5061005b6100e8565b604051808481526020018060200183151515158152602001828103825284818151815260200191508051906020019080838360005b838110156100ab578082015181840152602081019050610090565b50505050905090810190601f1680156100d85780820380516001836020036101000a031916815260200191505b5094505050505060405180910390f35b6000606060006001808191506040805190810160405280600b81526020017f7465737420737472696e67000000000000000000000000000000000000000000815250909250925092509091925600a165627a7a72305820042f9a292c15c34edba465c69fa0276c96c0948259172fdef0140c8c70e9138a0029"
	var theTime = utils.ToMilliSeconds(time.Now())

	var tx, _ = types.NewDeployContractTransaction(
		privateKey,
		from,
		code,
		theTime,
	)

	// TAKEN FROM `func (this *DAPoSService) startGossiping`

	var fakeReceipt = &types.Receipt{
		Id:                  "fake1",
		Type:                "fake1",
		Status:              "fake1",
		HumanReadableStatus: "fake1",
	}
	services.GetCache().Set(fakeReceipt.Id, fakeReceipt, types.ReceiptCacheTTL)

	var fakeGossip = &types.Gossip{
		ReceiptId:   fakeReceipt.Id,
		Transaction: *tx,
	}

	dapos.GetDAPoSService().Temp_ProcessTransaction(fakeGossip)
}

func executeMethod_1() {
	// Taken from Genesis
	var privateKey = "0f86ea981203b26b5b8244c8f661e30e5104555068a4bd168d3e3015db9bb25a"
	var from = "3ed25f42484d517cdfc72cafb7ebc9e8baa52c2c"
	var to = "348cbf15a7db41303b035f862293f9818a0e8b8f"
	var abi = `[
		{
			"constant": true,
			"inputs": [],
			"name": "getMultiReturn",
			"outputs": [
				{
					"name": "",
					"type": "int256"
				},
				{
					"name": "",
					"type": "string"
				},
				{
					"name": "",
					"type": "bool"
				}
			],
			"payable": false,
			"stateMutability": "view",
			"type": "function"
		}
	]`

	var theTime = utils.ToMilliSeconds(time.Now())
	var method = "getMultiReturn"
	var params = make([]interface{}, 0)
	// params[0] = "5555"
	// params[0] = 5555

	var tx, _ = types.NewExecuteContractTransaction(
		privateKey,
		from,
		to,
		hex.EncodeToString([]byte(abi)),
		method,
		params,
		theTime,
	)

	var fakeReceipt = &types.Receipt{
		Id:                  "fake2",
		Type:                "fake2",
		Status:              "fake2",
		HumanReadableStatus: "fake2",
	}
	services.GetCache().Set(fakeReceipt.Id, fakeReceipt, types.ReceiptCacheTTL)

	var fakeGossip = &types.Gossip{
		ReceiptId:   fakeReceipt.Id,
		Transaction: *tx,
	}
	dapos.GetDAPoSService().Temp_ProcessTransaction(fakeGossip)
}

func executeMethod_getVar5() {
	// Taken from Genesis
	var privateKey = "0f86ea981203b26b5b8244c8f661e30e5104555068a4bd168d3e3015db9bb25a"
	var from = "3ed25f42484d517cdfc72cafb7ebc9e8baa52c2c"
	var to = "79af16594093f280252bce6431a352a869bc2c02"
	var abi = `[
		{
			"constant": false,
			"inputs": [],
			"name": "throwException",
			"outputs": [],
			"payable": false,
			"stateMutability": "nonpayable",
			"type": "function"
		},
		{
			"constant": true,
			"inputs": [],
			"name": "getVar5",
			"outputs": [
				{
					"name": "",
					"type": "string"
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
					"name": "value",
					"type": "string"
				}
			],
			"name": "setVar6Var4",
			"outputs": [],
			"payable": false,
			"stateMutability": "nonpayable",
			"type": "function"
		},
		{
			"constant": true,
			"inputs": [],
			"name": "var6",
			"outputs": [
				{
					"name": "var1",
					"type": "uint256"
				},
				{
					"name": "var2",
					"type": "bool"
				},
				{
					"name": "var3",
					"type": "uint8"
				},
				{
					"name": "var4",
					"type": "string"
				}
			],
			"payable": false,
			"stateMutability": "view",
			"type": "function"
		},
		{
			"constant": true,
			"inputs": [],
			"name": "getMultiReturn",
			"outputs": [
				{
					"name": "",
					"type": "string"
				},
				{
					"name": "",
					"type": "string"
				}
			],
			"payable": false,
			"stateMutability": "view",
			"type": "function"
		},
		{
			"constant": true,
			"inputs": [],
			"name": "var5",
			"outputs": [
				{
					"name": "",
					"type": "string"
				}
			],
			"payable": false,
			"stateMutability": "view",
			"type": "function"
		},
		{
			"constant": false,
			"inputs": [],
			"name": "incVar6Var1",
			"outputs": [],
			"payable": false,
			"stateMutability": "nonpayable",
			"type": "function"
		},
		{
			"constant": false,
			"inputs": [],
			"name": "logEvent",
			"outputs": [],
			"payable": false,
			"stateMutability": "nonpayable",
			"type": "function"
		},
		{
			"constant": true,
			"inputs": [],
			"name": "intVar",
			"outputs": [
				{
					"name": "",
					"type": "int256"
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
					"name": "value",
					"type": "string"
				}
			],
			"name": "setVar5",
			"outputs": [],
			"payable": false,
			"stateMutability": "nonpayable",
			"type": "function"
		},
		{
			"inputs": [],
			"payable": false,
			"stateMutability": "nonpayable",
			"type": "constructor"
		},
		{
			"anonymous": false,
			"inputs": [
				{
					"indexed": false,
					"name": "test1",
					"type": "string"
				},
				{
					"indexed": false,
					"name": "test2",
					"type": "string"
				}
			],
			"name": "testLog",
			"type": "event"
		}
	]`

	var theTime = utils.ToMilliSeconds(time.Now())
	var method = "logEvent" // "getVar5" // "logEvent"
	var params = make([]interface{}, 0)

	var tx, _ = types.NewExecuteContractTransaction(
		privateKey,
		from,
		to,
		hex.EncodeToString([]byte(abi)),
		method,
		params,
		theTime,
	)

	var fakeReceipt = &types.Receipt{
		Id:                  "fake3",
		Type:                "fake3",
		Status:              "fake3",
		HumanReadableStatus: "fake3",
	}
	services.GetCache().Set(fakeReceipt.Id, fakeReceipt, types.ReceiptCacheTTL)

	var fakeGossip = &types.Gossip{
		ReceiptId:   fakeReceipt.Id,
		Transaction: *tx,
	}
	dapos.GetDAPoSService().Temp_ProcessTransaction(fakeGossip)
}
