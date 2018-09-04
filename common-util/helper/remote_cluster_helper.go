package helper

import (
	"github.com/dispatchlabs/samples/common-util/configTypes"
	"github.com/dispatchlabs/disgo/commons/types"
)

func GetNewRemoteConfigs(seedNodes, delegateNodes []*configTypes.NodeInfo) map[string]*configTypes.NodeInfo {
	configMap := map[string]*configTypes.NodeInfo{}

	seedsConfig := make([]*types.Node, len(seedNodes))
	for i := 0; i < len(seedNodes); i++ {
		seedAccount := CreateSeedAccount();
		seedConfig := CreateSeedConfig(seedNodes[i].Host, seedNodes[i].HttpPort, seedNodes[i].GrpcPort, seedAccount)
		seedNodes[i].Account = seedAccount
		seedNodes[i].Config = seedConfig
		seedsConfig[i] = seedConfig.Seeds[0]
		configMap[seedNodes[i].Name] = seedNodes[i]
	}

	delegateAddressList := make([]string, len(delegateNodes))
	for i := 0; i < len(delegateNodes); i++ {
		delegateNodes[i].Config = CreateDelegateConfig(delegateNodes[i].Host, delegateNodes[i].HttpPort, delegateNodes[i].GrpcPort, seedsConfig)
		delegateNodes[i].Account = CreateDelegateAccount(delegateNodes[i].Name)

		delegateAddressList[i] = delegateNodes[i].Account.Address

		configMap[delegateNodes[i].Name] = delegateNodes[i]
	}
	for _, seedNode := range seedNodes {
		seedNode.Config.DelegateAddresses = delegateAddressList
	}
	return configMap
}
