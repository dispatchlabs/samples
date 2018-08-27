package helper

import (
	"os/user"
	"os"
	"github.com/dispatchlabs/samples/common-util/util"
	"github.com/dispatchlabs/samples/common-util/configTypes"
	"github.com/dispatchlabs/disgo/commons/utils"
	"fmt"
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
func CreateNewLocalConfigs(seedNode *configTypes.NodeInfo, delegateNodes []*configTypes.NodeInfo, restricted bool) map[string]*configTypes.NodeInfo {
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
	if restricted {
		seedConfig.DelegateAddresses = delegateAddressList
	}
	clusterStructure := configTypes.NewClusterStructure(GetDisgoDirectory(), GetDefaultDirectory(), 1, len(delegateNodes))
	for _, v := range configMap {
		SaveFiles("", v)
	}

	return configMap
}

func SaveFiles(rootDir string, nodeInfo *configTypes.NodeInfo) {
	if rootDir == "" {
		rootDir = GetDefaultDirectory()
	}
	nodeDir := rootDir + string(os.PathSeparator) + nodeInfo.Name
	accountFileName := nodeDir + string(os.PathSeparator) + "account.json"
	configFileName := nodeDir + string(os.PathSeparator) + "config.json"

	util.WriteFile(nodeDir, accountFileName, nodeInfo.Account.ToPrettyJson())
	util.WriteFile(nodeDir, configFileName, nodeInfo.Config.ToPrettyJson())
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
