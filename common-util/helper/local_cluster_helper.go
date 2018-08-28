package helper

import (
	"os/user"
	"os"
	"github.com/dispatchlabs/samples/common-util/configTypes"
	"github.com/dispatchlabs/disgo/commons/utils"
	"fmt"
	"github.com/dispatchlabs/disgo/commons/types"
)

/*
 * Functions specifically for creating local cluster to run
 */

func BuildDisgoExecutable() {
	cmd := fmt.Sprintf("cd %s; go build", GetDisgoDirectory())
	err := Exec(cmd)
	if err != nil {
		utils.Error(err)
	}
}

func UpdateDisgoExecutable(nodeName string) {
	nodeDir := GetDefaultDirectory() + string(os.PathSeparator) + nodeName + string(os.PathSeparator)

	cmd := fmt.Sprintf("cp %s/disgo %s", GetDisgoDirectory(), nodeDir)
	utils.Debug("Command: " + cmd)
	err := Exec(cmd)
	if err != nil {
		utils.Error(err)
	}
}

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
	BuildDisgoExecutable()
	for _, v := range configMap {
		clusterStructure.SaveAccountAndConfigFiles(v)
		UpdateDisgoExecutable(v.Name)
	}

	return configMap
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
