package helper

import (
	"os/user"
	"github.com/dispatchlabs/samples/common-util/configTypes"
	"github.com/dispatchlabs/disgo/commons/utils"
	"fmt"
	"github.com/dispatchlabs/disgo/commons/types"
)

/*
 * Functions specifically for creating local cluster to run
 */

// Restricted config specifies delegate list in the seed node so any nodes joining are a only type Node
// Non-Restricted config does not specify delegate list in the seed node so any nodes joining are a delegate
func CreateNewLocalConfigs(clusterStructure *configTypes.ClusterStructure, seedNodes []*configTypes.NodeInfo, delegateNodes []*configTypes.NodeInfo, restricted bool) map[string]*configTypes.NodeInfo {
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
	if restricted {
		for _, seedNode := range seedNodes {
			seedNode.Config.DelegateAddresses = delegateAddressList
		}
	}
	return configMap
}


func SetupDefaultConfig(host string, nbrSeeds, nbrDelegates, seedStartingPort, delegateStartingPort int) {
	clusterStructure := configTypes.NewClusterStructure(GetDisgoDirectory(), GetDefaultDirectory(), nbrSeeds, nbrDelegates)

	seedNodes :=  make([]*configTypes.NodeInfo, nbrSeeds)
	for i := 0; i < nbrSeeds; i++ {
		seedName := fmt.Sprintf("seed-%d", i)
		grpcPort := seedStartingPort + (i*4)
		httpPort := grpcPort + 2

		seedInfo := &configTypes.NodeInfo{seedName, host, int64(httpPort),int64(grpcPort), nil, nil}
		seedNodes[i] = seedInfo
	}

	delegateNodes := make([]*configTypes.NodeInfo, nbrDelegates)
	for i := 0; i < nbrDelegates; i++ {
		delegateName := fmt.Sprintf("delegate-%d", i)

		grpcPort := delegateStartingPort + (i*2)
		httpPort := grpcPort + 1

		delegateNodes[i] = &configTypes.NodeInfo{delegateName, host, int64(httpPort),int64(grpcPort), nil, nil}

	}
	configMap := CreateNewLocalConfigs(clusterStructure, seedNodes, delegateNodes, true)
	for _, v := range configMap {
		fmt.Printf("%s\n", v.ToPrettyJson())
	}
}

var defaultDirectory string
func GetDefaultDirectory() string {
	if defaultDirectory == "" {
		usr, err := user.Current()
		if err != nil {
			utils.Fatal( err )
		}
		defaultDirectory = usr.HomeDir + "/go/src/github.com/dispatchlabs/samples/run-nodes-locally"
	}
	return defaultDirectory
}

var disgoDir string
func GetDisgoDirectory() string {
	if disgoDir == "" {
		usr, err := user.Current()
		if err != nil {
			utils.Fatal( err )
		}
		disgoDir = usr.HomeDir + "/go/src/github.com/dispatchlabs/disgo"
	}
	return disgoDir
}
