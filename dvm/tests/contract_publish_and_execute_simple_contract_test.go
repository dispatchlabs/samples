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

	deployContract()

	// go func() {
	// 	// time.Sleep(10 * time.Second)
	// 	executeMethod_setVar5()
	// }()
}

func deployContract() {
	var privateKey = "0f86ea981203b26b5b8244c8f661e30e5104555068a4bd168d3e3015db9bb25a"
	var from = "3ed25f42484d517cdfc72cafb7ebc9e8baa52c2c"
	var code = "608060405234801561001057600080fd5b506040805190810160405280600d81526020017f61616161616161616161616161000000000000000000000000000000000000008152506000908051906020019061005c9291906100f7565b50600060016000018190555060006001800160006101000a81548160ff02191690831515021790555060018060010160016101000a81548160ff021916908360ff1602179055506040805190810160405280600b81526020017f6262626262626262626262000000000000000000000000000000000000000000815250600160020190805190602001906100f19291906100f7565b5061019c565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061013857805160ff1916838001178555610166565b82800160010185558215610166579182015b8281111561016557825182559160200191906001019061014a565b5b5090506101739190610177565b5090565b61019991905b8082111561019557600081600090555060010161017d565b5090565b90565b6103a1806101ab6000396000f300608060405260043610610062576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff16806333e538e91461006757806334e45f53146100f757806379af647314610160578063cb69e30014610177575b600080fd5b34801561007357600080fd5b5061007c6101e0565b6040518080602001828103825283818151815260200191508051906020019080838360005b838110156100bc5780820151818401526020810190506100a1565b50505050905090810190601f1680156100e95780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34801561010357600080fd5b5061015e600480360381019080803590602001908201803590602001908080601f0160208091040260200160405190810160405280939291908181526020018383808284378201915050505050509192919290505050610282565b005b34801561016c57600080fd5b5061017561029f565b005b34801561018357600080fd5b506101de600480360381019080803590602001908201803590602001908080601f01602080910402602001604051908101604052809392919081815260200183838082843782019150505050505091929192905050506102b6565b005b606060008054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156102785780601f1061024d57610100808354040283529160200191610278565b820191906000526020600020905b81548152906001019060200180831161025b57829003601f168201915b5050505050905090565b806001600201908051906020019061029b9291906102d0565b5050565b600160000160008154809291906001019190505550565b80600090805190602001906102cc9291906102d0565b5050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061031157805160ff191683800117855561033f565b8280016001018555821561033f579182015b8281111561033e578251825591602001919060010190610323565b5b50905061034c9190610350565b5090565b61037291905b8082111561036e576000816000905550600101610356565b5090565b905600a165627a7a723058203111c899d1b9b1f4149c5245a496b9f11b8c0ee21feeec5260b4c9fe5bcd86470029"
	var theTime = utils.ToMilliSeconds(time.Now())

	var tx, _ = types.NewContractTransaction(
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

func executeMethod_setVar5() {
	// Taken from Genesis
	var privateKey = "0f86ea981203b26b5b8244c8f661e30e5104555068a4bd168d3e3015db9bb25a"
	var from = "3ed25f42484d517cdfc72cafb7ebc9e8baa52c2c"
	var to = "c3be1a3a5c6134cca51896fadf032c4c61bc355e" // "c3be1a3a5c6134cca51896fadf032c4c61bc355e"
	var abi = `[
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
		}
	]`

	var theTime = utils.ToMilliSeconds(time.Now())
	var method = "setVar5"
	var params = make([]interface{}, 1)
	params[0] = "5555"

	var tx, _ = types.NewContractCallTransaction(
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
