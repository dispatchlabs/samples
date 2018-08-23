package helper

import (
	"fmt"
	"time"
	"encoding/hex"

	"github.com/dispatchlabs/disgo/commons/types"
	"github.com/dispatchlabs/disgo/commons/utils"
)

var deployOccurred = false

func GetRandomTransaction(toAddress string) *types.Transaction {
	value := utils.Random(1, 3)
	var contractAddress string
	switch value {
	case 1:
		return GetTransaction(toAddress)
	case 2:
		//deployOccurred = true
		return GetNewDeployTx()
	case 3:
		if deployOccurred {
			return GetNewExecuteTx(contractAddress, "getVar5")
		} else {
			return GetNewDeployTx()
		}
	}
	return nil
}

func GetTransaction(toAddress string) *types.Transaction {
	utils.Info("GetTransaction")
	var privateKey = "0f86ea981203b26b5b8244c8f661e30e5104555068a4bd168d3e3015db9bb25a"
	var from = "3ed25f42484d517cdfc72cafb7ebc9e8baa52c2c"

	var tx, _ = types.NewTransferTokensTransaction(
		privateKey,
		from,
		toAddress,
		1,
		1,
		utils.ToMilliSeconds(time.Now()),
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


func GetNewDaveDeployTx() *types.Transaction {
	utils.Info("GetNewBadDeployTx")

	var privateKey = "0f86ea981203b26b5b8244c8f661e30e5104555068a4bd168d3e3015db9bb25a"
	var from = "3ed25f42484d517cdfc72cafb7ebc9e8baa52c2c"

	tx, err := types.NewDeployContractTransaction(
		privateKey,
		from,
		getDaveCode(),
		getDaveAbi(),
		utils.ToMilliSeconds(time.Now()),
	)
	if err != nil {
		utils.Error(err)
	}
	//
	fmt.Printf("DEPLOY: %s\n", tx.ToPrettyJson())
	return tx
}


func GetNewExecuteTx(toAddress string, method string) *types.Transaction {
	utils.Info("GetNewExecuteTx")
	// Taken from Genesis
	var privateKey = "0f86ea981203b26b5b8244c8f661e30e5104555068a4bd168d3e3015db9bb25a"
	var from = "3ed25f42484d517cdfc72cafb7ebc9e8baa52c2c"

	var tx, _ = types.NewExecuteContractTransaction(
		privateKey,
		from,
		toAddress,
		hex.EncodeToString([]byte(getDaveAbi())),
		//hex.EncodeToString([]byte(getAbi())),
		method,
		getParamsForMethod(method),
		utils.ToMilliSeconds(time.Now()),
	)

	return tx

}

func getParamsForMethod(method string) []interface{} {
	var params = make([]interface{}, 0)

	switch method {
	case "setVar5":
		params = append(params, "Abcdefg")
		break
	case "setVar6Var4":
		params = append(params, "test value for var4")
		break;
	case "intParam":
		params = append(params, 20)
		break
	case "plusOne":
		params = append(params, 1)
		break
	case "uintParam":
		params = append(params, uint(30))
		break
	case "boolParamType":
		params = append(params, true)
		break
	case "multiParams":
		params = append(params, "test1")
		params = append(params, "test2")
		break
	case "arrayParam":
		var array = make([]interface{}, 1)
		array[0] = uint(20)
		params = append(params, array)
		break
		//all fall through below
	case "getVar5":
	case "returnBool":
	case "testLog":
	case "returnInt":
	case "returnUint":
	case "throwException":

		break

	}
	return params
}

func getCode() string {
	 return "608060405234801561001057600080fd5b5060408051808201909152600d8082527f61616161616161616161616161000000000000000000000000000000000000006020909201918252610055916000916100b4565b5060006002556003805461ffff191661010017905560408051808201909152600b8082527f626262626262626262626200000000000000000000000000000000000000000060209092019182526100ae916004916100b4565b5061014f565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106100f557805160ff1916838001178555610122565b82800160010185558215610122579182015b82811115610122578251825591602001919060010190610107565b5061012e929150610132565b5090565b61014c91905b8082111561012e5760008155600101610138565b90565b610a318061015e6000396000f3006080604052600436106101065763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166304c6f56a811461010b5780631af35da3146101985780631e358c3e146101c1578063216a52e514610228578063222e0412146102c157806330bc6db2146102d657806333e538e9146102eb57806334e45f53146103005780633a458b1f146103595780634a846e02146104055780634aea8b141461010b57806378d8866e146104f857806379af64731461050d578063943640c314610522578063a5fe087214610560578063af44550014610575578063cb69e3001461058a578063d13f25ad146102d6578063fd213d0c146105e3575b600080fd5b34801561011757600080fd5b506101236004356105f9565b6040805160208082528351818301528351919283929083019185019080838360005b8381101561015d578181015183820152602001610145565b50505050905090810190601f16801561018a5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b3480156101a457600080fd5b506101ad610631565b604080519115158252519081900360200190f35b3480156101cd57600080fd5b5060408051602060048035808201358381028086018501909652808552610216953695939460249493850192918291850190849080828437509497506106379650505050505050565b60408051918252519081900360200190f35b34801561023457600080fd5b506040805160206004803580820135601f81018490048402850184019095528484526102bf94369492936024939284019190819084018382808284375050604080516020601f89358b018035918201839004830284018301909452808352979a9998810197919650918201945092508291508401838280828437509497506106599650505050505050565b005b3480156102cd57600080fd5b506102bf61065d565b3480156102e257600080fd5b50610216610664565b3480156102f757600080fd5b50610123610669565b34801561030c57600080fd5b506040805160206004803580820135601f81018490048402850184019095528484526102bf9436949293602493928401919081908401838280828437509497506106ff9650505050505050565b34801561036557600080fd5b5061036e610712565b60405180858152602001841515151581526020018360ff1660ff16815260200180602001828103825283818151815260200191508051906020019080838360005b838110156103c75781810151838201526020016103af565b50505050905090810190601f1680156103f45780820380516001836020036101000a031916815260200191505b509550505050505060405180910390f35b34801561041157600080fd5b5061041a6107bb565b604051808060200180602001838103835285818151815260200191508051906020019080838360005b8381101561045b578181015183820152602001610443565b50505050905090810190601f1680156104885780820380516001836020036101000a031916815260200191505b50838103825284518152845160209182019186019080838360005b838110156104bb5781810151838201526020016104a3565b50505050905090810190601f1680156104e85780820380516001836020036101000a031916815260200191505b5094505050505060405180910390f35b34801561050457600080fd5b50610123610826565b34801561051957600080fd5b506102bf6108b4565b34801561052e57600080fd5b506105376108bf565b6040805173ffffffffffffffffffffffffffffffffffffffff9092168252519081900360200190f35b34801561056c57600080fd5b506102bf6108c3565b34801561058157600080fd5b50610216610958565b34801561059657600080fd5b506040805160206004803580820135601f81018490048402850184019095528484526102bf94369492936024939284019190819084018382808284375094975061095e9650505050505050565b3480156105ef57600080fd5b5061012360043515155b5060408051808201909152600481527f7465737400000000000000000000000000000000000000000000000000000000602082015290565b60005b90565b600081600081518110151561064857fe5b906020019060200201519050919050565b5050565b6000806005fe5b601490565b60008054604080516020601f60026000196101006001881615020190951694909404938401819004810282018101909252828152606093909290918301828280156106f55780601f106106ca576101008083540402835291602001916106f5565b820191906000526020600020905b8154815290600101906020018083116106d857829003601f168201915b5050505050905090565b805161065990600490602084019061096d565b60028054600354600480546040805160206101006001851615810260001901909416889004601f8101829004820283018201909352828252959660ff8087169794909604909516949390929091908301828280156107b15780601f10610786576101008083540402835291602001916107b1565b820191906000526020600020905b81548152906001019060200180831161079457829003601f168201915b5050505050905084565b60408051808201825260058082527f746573743100000000000000000000000000000000000000000000000000000060208084019190915283518085019094529083527f74657374320000000000000000000000000000000000000000000000000000009083015291565b6000805460408051602060026001851615610100026000190190941693909304601f810184900484028201840190925281815292918301828280156108ac5780601f10610881576101008083540402835291602001916108ac565b820191906000526020600020905b81548152906001019060200180831161088f57829003601f168201915b505050505081565b600280546001019055565b3090565b6040805181815260058183018190527f746573743100000000000000000000000000000000000000000000000000000060608301526080602083018190528201527f746573743200000000000000000000000000000000000000000000000000000060a082015290517fb20aa8922321b2e5be1e9784294eda54d640a58038ceede50492f7d7ffc8ad629181900360c00190a1565b60015481565b80516106599060009060208401905b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106109ae57805160ff19168380011785556109db565b828001600101855582156109db579182015b828111156109db5782518255916020019190600101906109c0565b506109e79291506109eb565b5090565b61063491905b808211156109e757600081556001016109f15600a165627a7a723058203eab72d39ee65aa562fd534bb59b01a7449a3dd657923f38fa03290e1f3f8a1e0029"
}

func getDaveCode() string {
	return "6080604052348015600f57600080fd5b50609c8061001e6000396000f300608060405260043610603e5763ffffffff7c0100000000000000000000000000000000000000000000000000000000600035041663f5a6259f81146043575b600080fd5b348015604e57600080fd5b506058600435606a565b60408051918252519081900360200190f35b600101905600a165627a7a7230582052a887255bee69b86c68b80729c39cb6d8c2651404d12f5b12ce002ebf8f1b0b0029"
}

func getDaveAbi() string {
	return `[{"constant":true,"inputs":[{"name":"y","type":"uint256"}],"name":"plusOne","outputs":[{"name":"x","type":"uint256"}],"payable":false,"stateMutability":"pure","type":"function"}]`	
}
func getAbi() string {

	return `[
	{
		"constant": true,
		"inputs": [
			{
				"name": "param",
				"type": "int256"
			}
		],
		"name": "intParam",
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
		"name": "returnBool",
		"outputs": [
			{
				"name": "",
				"type": "bool"
			}
		],
		"payable": false,
		"stateMutability": "nonpayable",
		"type": "function"
	},
	{
		"constant": true,
		"inputs": [
			{
				"name": "param",
				"type": "uint256[]"
			}
		],
		"name": "arrayParam",
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
				"name": "value",
				"type": "string"
			},
			{
				"name": "value2",
				"type": "string"
			}
		],
		"name": "multiParams",
		"outputs": [],
		"payable": false,
		"stateMutability": "nonpayable",
		"type": "function"
	},
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
		"name": "returnUint",
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
		"inputs": [
			{
				"name": "param",
				"type": "uint256"
			}
		],
		"name": "uintParam",
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
		"constant": true,
		"inputs": [],
		"name": "returnAddress",
		"outputs": [
			{
				"name": "",
				"type": "address"
			}
		],
		"payable": false,
		"stateMutability": "view",
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
		"constant": true,
		"inputs": [],
		"name": "returnInt",
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
		"constant": true,
		"inputs": [
			{
				"name": "value",
				"type": "bool"
			}
		],
		"name": "boolParamType",
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
}