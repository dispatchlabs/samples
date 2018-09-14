package tests

import (
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

var LoopsTesting_nrOfServices = 6

// ContractPublishAndExecuteSimpleContract -
func Test_LoopsTesting(t *testing.T) {
	utils.InitMainPackagePath()
	utils.InitializeLogger()

	utils.Events().On(services.Events.DbServiceInitFinished, LoopsTesting_allServicesInitFinished)
	utils.Events().On(services.Events.GrpcServiceInitFinished, LoopsTesting_allServicesInitFinished)
	utils.Events().On(services.Events.HttpServiceInitFinished, LoopsTesting_allServicesInitFinished)

	utils.Events().On(disgover.Events.DisGoverServiceInitFinished, LoopsTesting_allServicesInitFinished)
	utils.Events().On(dapos.Events.DAPoSServiceInitFinished, LoopsTesting_allServicesInitFinished)
	utils.Events().On(dvm.Events.DVMServiceInitFinished, LoopsTesting_allServicesInitFinished)

	utils.Info(fmt.Sprintf("NR of services left to be started: %d", LoopsTesting_nrOfServices))

	server := bootstrap.NewServer()
	server.Go()
}

func LoopsTesting_allServicesInitFinished() {
	LoopsTesting_nrOfServices--
	utils.Info(fmt.Sprintf("NR of services left to be started: %d", LoopsTesting_nrOfServices))

	if LoopsTesting_nrOfServices > 0 {
		return
	}

	const timeout = 3

	// go func() {
	// 	time.Sleep(timeout * time.Second)
	// 	LoopsTesting_deployContract()
	// }()

	go func() {
		time.Sleep(timeout * time.Second)
		executeMethod_IncHundredTimes()
	}()
}

func LoopsTesting_deployContract() {
	var privateKey = "0f86ea981203b26b5b8244c8f661e30e5104555068a4bd168d3e3015db9bb25a"
	var from = "3ed25f42484d517cdfc72cafb7ebc9e8baa52c2c"
	var code = "608060405234801561001057600080fd5b5061039f806100206000396000f300608060405260043610610083576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff16806311b9742a146100885780631a434eaf146100b3578063496cfdb9146100de5780637bf05435146101095780638ab3bdff14610134578063b594e64b1461015f578063cc9da9ca1461018a575b600080fd5b34801561009457600080fd5b5061009d6101b5565b6040518082815260200191505060405180910390f35b3480156100bf57600080fd5b506100c86101d3565b6040518082815260200191505060405180910390f35b3480156100ea57600080fd5b506100f361022a565b6040518082815260200191505060405180910390f35b34801561011557600080fd5b5061011e6102a1565b6040518082815260200191505060405180910390f35b34801561014057600080fd5b506101496102d7565b6040518082815260200191505060405180910390f35b34801561016b57600080fd5b5061017461030b565b6040518082815260200191505060405180910390f35b34801561019657600080fd5b5061019f61033e565b6040518082815260200191505060405180910390f35b600080600090505b80806001019150508080600190039150506101bd565b60008060008060009250600091505b633b9aca0082101561022157600090505b633b9aca0081101561021457828060010193505080806001019150506101f3565b81806001019250506101e2565b82935050505090565b6000806000806000809350600092505b633b9aca0083101561029757600091505b633b9aca0082101561028a57600090505b633b9aca0081101561027d578380600101945050808060010191505061025c565b818060010192505061024b565b828060010193505061023a565b8394505050505090565b6000806000809150600090505b633b9aca008110156102cf57818060010192505080806001019150506102ae565b819250505090565b6000806000809150600090505b6103e881101561030357818060010192505080806001019150506102e4565b819250505090565b6000806000809150600090505b60648110156103365781806001019250508080600101915050610318565b819250505090565b6000806000809150600090505b620f424081101561036b578180600101925050808060010191505061034b565b8192505050905600a165627a7a72305820ea4b26a846f4f7f2b12313c781ef609e7d10474daa32af5f117030317f0a8ddc0029"
	var theTime = utils.ToMilliSeconds(time.Now())

	var tx, _ = types.NewDeployContractTransaction(
		privateKey,
		from,
		code,
		"",
		theTime,
	)

	// TAKEN FROM: func (this *DAPoSService) startGossiping(transaction *types.Transaction)

	// // Cache gossip with my rumor.
	// fakeGossip := types.NewGossip(*tx)
	// rumor := types.NewRumor(types.GetAccount().PrivateKey, types.GetAccount().Address, tx.Hash)
	// fakeGossip.Rumors = append(fakeGossip.Rumors, *rumor)
	// fakeGossip.Cache(services.GetCache())

	dapos.GetDAPoSService().Temp_ProcessTransaction(tx)
}

func executeMethod_IncHundredTimes() {
	// Taken from Genesis
	var privateKey = "0f86ea981203b26b5b8244c8f661e30e5104555068a4bd168d3e3015db9bb25a"
	var from = "3ed25f42484d517cdfc72cafb7ebc9e8baa52c2c"
	var to = "627e04f584776e7631be4ffc424ed44fece8418c"
	// var abi = `[
	// 	{
	// 		"constant": true,
	// 		"inputs": [],
	// 		"name": "IncInfiniteTimes",
	// 		"outputs": [
	// 			{
	// 				"name": "",
	// 				"type": "uint256"
	// 			}
	// 		],
	// 		"payable": false,
	// 		"stateMutability": "pure",
	// 		"type": "function"
	// 	},
	// 	{
	// 		"constant": true,
	// 		"inputs": [],
	// 		"name": "IncBilTimesForFor",
	// 		"outputs": [
	// 			{
	// 				"name": "",
	// 				"type": "uint256"
	// 			}
	// 		],
	// 		"payable": false,
	// 		"stateMutability": "pure",
	// 		"type": "function"
	// 	},
	// 	{
	// 		"constant": true,
	// 		"inputs": [],
	// 		"name": "IncBilTimesForForFor",
	// 		"outputs": [
	// 			{
	// 				"name": "",
	// 				"type": "uint256"
	// 			}
	// 		],
	// 		"payable": false,
	// 		"stateMutability": "pure",
	// 		"type": "function"
	// 	},
	// 	{
	// 		"constant": true,
	// 		"inputs": [],
	// 		"name": "IncBilTimes",
	// 		"outputs": [
	// 			{
	// 				"name": "",
	// 				"type": "uint256"
	// 			}
	// 		],
	// 		"payable": false,
	// 		"stateMutability": "pure",
	// 		"type": "function"
	// 	},
	// 	{
	// 		"constant": true,
	// 		"inputs": [],
	// 		"name": "IncThousandTimes",
	// 		"outputs": [
	// 			{
	// 				"name": "",
	// 				"type": "uint256"
	// 			}
	// 		],
	// 		"payable": false,
	// 		"stateMutability": "pure",
	// 		"type": "function"
	// 	},
	// 	{
	// 		"constant": true,
	// 		"inputs": [],
	// 		"name": "IncHundredTimes",
	// 		"outputs": [
	// 			{
	// 				"name": "",
	// 				"type": "uint256"
	// 			}
	// 		],
	// 		"payable": false,
	// 		"stateMutability": "pure",
	// 		"type": "function"
	// 	},
	// 	{
	// 		"constant": true,
	// 		"inputs": [],
	// 		"name": "IncMilTimes",
	// 		"outputs": [
	// 			{
	// 				"name": "",
	// 				"type": "uint256"
	// 			}
	// 		],
	// 		"payable": false,
	// 		"stateMutability": "pure",
	// 		"type": "function"
	// 	}
	// ]`

	var theTime = utils.ToMilliSeconds(time.Now())
	var method = "IncBilTimes"
	var params = make([]interface{}, 0)

	var tx, _ = types.NewExecuteContractTransaction(
		privateKey,
		from,
		to,
		// hex.EncodeToString([]byte(abi)),
		method,
		params,
		theTime,
	)

	// TAKEN FROM: func (this *DAPoSService) startGossiping(transaction *types.Transaction)

	// Cache gossip with my rumor.
	// fakeGossip := types.NewGossip(*tx)
	// rumor := types.NewRumor(types.GetAccount().PrivateKey, types.GetAccount().Address, tx.Hash)
	// fakeGossip.Rumors = append(fakeGossip.Rumors, *rumor)
	// fakeGossip.Cache(services.GetCache())

	dapos.GetDAPoSService().Temp_ProcessTransaction(tx)
}
