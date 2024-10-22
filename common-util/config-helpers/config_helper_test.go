package configHelpers

import (
	"fmt"
	"testing"
)

func TestSeedConfig(t *testing.T) {

	seedAccount := CreateSeedAccount()

	fmt.Printf("%s\n", seedAccount.ToPrettyJson())
	seedConfig := GetSeedConfig("127.0.0.1", 1975, 1973, seedAccount)

	fmt.Printf("%s\n", seedConfig.ToPrettyJson())

}

func TestDelegateConfig(t *testing.T) {
	nbrDelegates := 5
	var startingPort int64
	startingPort = 3500

	seedAccount := CreateSeedAccount()
	seedConfig := GetSeedConfig("127.0.0.1", 1975, 1973, seedAccount)

	for i := 1; i <= nbrDelegates; i++ {
		delegateName := fmt.Sprintf("delegate-%d", i)
		startingPort++
		grpcPort := startingPort
		startingPort++
		httpPort := startingPort
		delegateConfig := GetDelegateConfig(delegateName, httpPort, grpcPort, seedConfig.Seeds)
		fmt.Printf("%s:\n%s\n", delegateName, delegateConfig.ToPrettyJson())
	}
}

func TestNodeInfo(t *testing.T) {
	configMap := map[string]NodeInfo{}

	nbrDelegates := 5
	var startingPort int64
	startingPort = 3500

	seedAccount := CreateSeedAccount()
	seedConfig := GetSeedConfig("127.0.0.1", 1975, 1973, seedAccount)
	seedNode := NodeInfo{"seed", "127.0.0.1", seedConfig.HttpEndpoint.Port, seedConfig.GrpcEndpoint.Port, seedAccount, seedConfig}

	configMap[seedNode.Name] = seedNode
	delegateAddressList := make([]string, nbrDelegates)
	for i := 1; i <= nbrDelegates; i++ {
		delegateName := fmt.Sprintf("delegate-%d", i)
		startingPort++
		grpcPort := startingPort
		startingPort++
		httpPort := startingPort
		delegateConfig := GetDelegateConfig("127.0.0.1", httpPort, grpcPort, seedNode.Config.Seeds)
		delegateAccount := CreateDelegateAccount(delegateName)

		delegateAddressList[i-1] = delegateAccount.Address
		configMap[delegateName] = NodeInfo{delegateName, "127.0.0.1", httpPort, grpcPort, delegateAccount, delegateConfig}
	}
	seedNode.Config.DelegateAddresses = delegateAddressList
	for _, v := range configMap {
		fmt.Printf("%s\n", v.ToPrettyJson())
	}
}

func TestRemoteFullConfig(t *testing.T) {
	seedNode := &NodeInfo{"stage-seed-0", "35.203.143.69", 1975, 1973, nil, nil}

	delegateNodes := make([]*NodeInfo, 0)

	delegateNodes = append(delegateNodes, &NodeInfo{"stage-delegate-0", "35.233.231.3", 1975, 1973, nil, nil})
	delegateNodes = append(delegateNodes, &NodeInfo{"stage-delegate-1", "35.233.241.115", 1975, 1973, nil, nil})
	delegateNodes = append(delegateNodes, &NodeInfo{"stage-delegate-2", "35.230.0.126", 1975, 1973, nil, nil})

	configMap := GetNewRemoteConfigs(seedNode, delegateNodes)
	for _, v := range configMap {
		fmt.Printf("%s\n", v.ToPrettyJson())
	}

}
