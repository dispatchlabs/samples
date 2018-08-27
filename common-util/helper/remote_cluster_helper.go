package helper

import "github.com/dispatchlabs/samples/common-util/configTypes"

func GetNewRemoteConfigs(seedNode *configTypes.NodeInfo, delegateNodes []*configTypes.NodeInfo) map[string]*configTypes.NodeInfo {
	configMap := map[string]*configTypes.NodeInfo{}

	seedAccount := CreateSeedAccount();
	seedConfig := CreateSeedConfig(seedNode.Host, seedNode.HttpPort, seedNode.GrpcPort, seedAccount)
	seedNode.Account = seedAccount
	seedNode.Config = seedConfig
	configMap["seed"] = seedNode

	delegateAddressList := make([]string, len(delegateNodes))
	for i := 0; i < len(delegateNodes); i++ {
		delegateNodes[i].Config = CreateDelegateConfig(delegateNodes[i].Host, delegateNodes[i].HttpPort, delegateNodes[i].GrpcPort, seedConfig.Seeds)
		delegateNodes[i].Account = CreateDelegateAccount(delegateNodes[i].Name)

		delegateAddressList[i] = delegateNodes[i].Account.Address

		configMap[delegateNodes[i].Name] = delegateNodes[i]
	}
	seedConfig.DelegateAddresses = delegateAddressList
	return configMap
}
