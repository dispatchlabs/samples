
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

	var theTime = utils.ToMilliSeconds(time.Now())
	var method = "logEvent" // "getVar5" // "logEvent"
	var params = make([]interface{}, 0)
	// var params = make([]interface{}, 1)
	// params[0] = "5555"

	var tx, _ = types.NewExecuteContractTransaction(
		privateKey,
		from,
		to,
		hex.EncodeToString([]byte(abi)),
		method,
		params,
		theTime,
	)

	// fakeGossip := types.NewGossip(*tx)
	// rumor := types.NewRumor(types.GetAccount().PrivateKey, types.GetAccount().Address, tx.Hash)
	// fakeGossip.Rumors = append(fakeGossip.Rumors, *rumor)
	// fakeGossip.Cache(services.GetCache())

	// dapos.GetDAPoSService().Temp_ProcessTransaction(fakeGossip)

	dapos.GetDAPoSService().Temp_ProcessTransaction(tx)
}

	//go func() {
	//	time.Sleep(timeout * time.Second)
	//	executeMethod_setVar5()
	//}()

	// go func() {
	// 	time.Sleep(timeout * time.Second)
	var code = "608060405234801561001057600080fd5b506040805190810160405280600d81526020017f61616161616161616161616161000000000000000000000000000000000000008152506000908051906020019061005c9291906100f8565b5060006002600001819055506000600260010160006101000a81548160ff0219169083151502179055506001600260010160016101000a81548160ff021916908360ff1602179055506040805190810160405280600b81526020017f62626262626262626262620000000000000000000000000000000000000000008152506002800190805190602001906100f29291906100f8565b5061019d565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061013957805160ff1916838001178555610167565b82800160010185558215610167579182015b8281111561016657825182559160200191906001019061014b565b5b5090506101749190610178565b5090565b61019a91905b8082111561019657600081600090555060010161017e565b5090565b90565b610e7f80620001ad6000396000f300608060405260043610610107576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff16806304c6f56a1461010c5780631af35da3146101b25780631e358c3e146101e1578063216a52e51461025b578063222e04121461030a57806330bc6db21461032157806333e538e91461034c57806334e45f53146103dc5780633a458b1f146104455780634a846e02146104f45780634aea8b14146105f057806378d8866e1461069657806379af647314610726578063943640c31461073d578063a5fe087214610794578063af445500146107ab578063cb69e300146107d6578063d13f25ad1461083f578063fd213d0c1461086a575b600080fd5b34801561011857600080fd5b5061013760048036038101908080359060200190929190505050610912565b6040518080602001828103825283818151815260200191508051906020019080838360005b8381101561017757808201518184015260208101905061015c565b50505050905090810190601f1680156101a45780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b3480156101be57600080fd5b506101c7610951565b604051808215151515815260200191505060405180910390f35b3480156101ed57600080fd5b5061024560048036038101908080359060200190820180359060200190808060200260200160405190810160405280939291908181526020018383602002808284378201915050505050509192919290505050610959565b6040518082815260200191505060405180910390f35b34801561026757600080fd5b50610308600480360381019080803590602001908201803590602001908080601f0160208091040260200160405190810160405280939291908181526020018383808284378201915050505050509192919290803590602001908201803590602001908080601f016020809104026020016040519081016040528093929190818152602001838380828437820191505050505050919291929050505061097b565b005b34801561031657600080fd5b5061031f61097f565b005b34801561032d57600080fd5b50610336610993565b6040518082815260200191505060405180910390f35b34801561035857600080fd5b5061036161099c565b6040518080602001828103825283818151815260200191508051906020019080838360005b838110156103a1578082015181840152602081019050610386565b50505050905090810190601f1680156103ce5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b3480156103e857600080fd5b50610443600480360381019080803590602001908201803590602001908080601f0160208091040260200160405190810160405280939291908181526020018383808284378201915050505050509192919290505050610a3e565b005b34801561045157600080fd5b5061045a610a5a565b60405180858152602001841515151581526020018360ff1660ff16815260200180602001828103825283818151815260200191508051906020019080838360005b838110156104b657808201518184015260208101905061049b565b50505050905090810190601f1680156104e35780820380516001836020036101000a031916815260200191505b509550505050505060405180910390f35b34801561050057600080fd5b50610509610b2a565b604051808060200180602001838103835285818151815260200191508051906020019080838360005b8381101561054d578082015181840152602081019050610532565b50505050905090810190601f16801561057a5780820380516001836020036101000a031916815260200191505b50838103825284818151815260200191508051906020019080838360005b838110156105b3578082015181840152602081019050610598565b50505050905090810190601f1680156105e05780820380516001836020036101000a031916815260200191505b5094505050505060405180910390f35b3480156105fc57600080fd5b5061061b60048036038101908080359060200190929190505050610ba1565b6040518080602001828103825283818151815260200191508051906020019080838360005b8381101561065b578082015181840152602081019050610640565b50505050905090810190601f1680156106885780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b3480156106a257600080fd5b506106ab610be0565b6040518080602001828103825283818151815260200191508051906020019080838360005b838110156106eb5780820151818401526020810190506106d0565b50505050905090810190601f1680156107185780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34801561073257600080fd5b5061073b610c7e565b005b34801561074957600080fd5b50610752610c95565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b3480156107a057600080fd5b506107a9610c9d565b005b3480156107b757600080fd5b506107c0610d40565b6040518082815260200191505060405180910390f35b3480156107e257600080fd5b5061083d600480360381019080803590602001908201803590602001908080601f0160208091040260200160405190810160405280939291908181526020018383808284378201915050505050509192919290505050610d46565b005b34801561084b57600080fd5b50610854610d60565b6040518082815260200191505060405180910390f35b34801561087657600080fd5b50610897600480360381019080803515159060200190929190505050610d69565b6040518080602001828103825283818151815260200191508051906020019080838360005b838110156108d75780820151818401526020810190506108bc565b50505050905090810190601f1680156109045780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b60606040805190810160405280600481526020017f74657374000000000000000000000000000000000000000000000000000000008152509050919050565b600080905090565b600081600081518110151561096a57fe5b906020019060200201519050919050565b5050565b600080600581151561098d57fe5b04905050565b60006014905090565b606060008054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015610a345780601f10610a0957610100808354040283529160200191610a34565b820191906000526020600020905b815481529060010190602001808311610a1757829003601f168201915b5050505050905090565b80600280019080519060200190610a56929190610dae565b5050565b60028060000154908060010160009054906101000a900460ff16908060010160019054906101000a900460ff1690806002018054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015610b205780601f10610af557610100808354040283529160200191610b20565b820191906000526020600020905b815481529060010190602001808311610b0357829003601f168201915b5050505050905084565b6060806040805190810160405280600581526020017f74657374310000000000000000000000000000000000000000000000000000008152506040805190810160405280600581526020017f7465737432000000000000000000000000000000000000000000000000000000815250915091509091565b60606040805190810160405280600481526020017f74657374000000000000000000000000000000000000000000000000000000008152509050919050565b60008054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015610c765780601f10610c4b57610100808354040283529160200191610c76565b820191906000526020600020905b815481529060010190602001808311610c5957829003601f168201915b505050505081565b600260000160008154809291906001019190505550565b600030905090565b7fb20aa8922321b2e5be1e9784294eda54d640a58038ceede50492f7d7ffc8ad62604051808060200180602001838103835260058152602001807f7465737431000000000000000000000000000000000000000000000000000000815250602001838103825260058152602001807f74657374320000000000000000000000000000000000000000000000000000008152506020019250505060405180910390a1565b60015481565b8060009080519060200190610d5c929190610dae565b5050565b60006014905090565b606060008290506040805190810160405280600481526020017f7465737400000000000000000000000000000000000000000000000000000000815250915050919050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10610def57805160ff1916838001178555610e1d565b82800160010185558215610e1d579182015b82811115610e1c578251825591602001919060010190610e01565b5b509050610e2a9190610e2e565b5090565b610e5091905b80821115610e4c576000816000905550600101610e34565b5090565b905600a165627a7a72305820710a757c70e06dfc554138895b2ca1b2d6ca5661c2d09f00de04339eed1be7250029"
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
	var to = "78337c25f0c003344c1b16e5f4b5ebda07a08cf5"
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

	var theTime = utils.ToMilliSeconds(time.Now())
	var method = "logEvent" // "getVar5" // "logEvent"
	var params = make([]interface{}, 0)
	// var params = make([]interface{}, 1)
	// params[0] = "5555"

	var tx, _ = types.NewExecuteContractTransaction(
		privateKey,
		from,
		to,
		hex.EncodeToString([]byte(abi)),
		method,
		params,
		theTime,
	)

	// fakeGossip := types.NewGossip(*tx)
	// rumor := types.NewRumor(types.GetAccount().PrivateKey, types.GetAccount().Address, tx.Hash)
	// fakeGossip.Rumors = append(fakeGossip.Rumors, *rumor)
	// fakeGossip.Cache(services.GetCache())

	// dapos.GetDAPoSService().Temp_ProcessTransaction(fakeGossip)

	dapos.GetDAPoSService().Temp_ProcessTransaction(tx)
}
