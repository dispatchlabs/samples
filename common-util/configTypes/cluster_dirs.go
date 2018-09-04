package configTypes

import (
	"fmt"
	"os"
	"github.com/dispatchlabs/samples/common-util/util"
)

type ClusterStructure struct {
	DisgoDir		string
	ClusterRoot 	string
	NodeDirs		map[string]string
	DelegateDirs 	map[string]string
	AccountFileName string
	ConfigFileName  string
}

func NewClusterStructure(disgoDir, clusterRoot string, nbrSeeds, nbrDelegates int) *ClusterStructure {

	return &ClusterStructure{
		DisgoDir:		disgoDir,
		ClusterRoot:	clusterRoot,
		NodeDirs:		getNodeDirs(clusterRoot, nbrSeeds, nbrDelegates),
		AccountFileName: "account.json",
		ConfigFileName: "config.json",
	}
}

func getNodeDirs(clusterRoot string, nbrSeeds, nbrDelegates int) map[string]string {
	nodeDirs := map[string]string{}
	for i := 0; i < nbrSeeds; i++ {
		seedName := fmt.Sprintf("seed-%d", i)
		nodeDirs[seedName] = clusterRoot + string(os.PathSeparator) + seedName
	}
	for i := 0; i < nbrDelegates; i++ {
		delegateName := fmt.Sprintf("delegate-%d", i)
		nodeDirs[delegateName] = clusterRoot + string(os.PathSeparator) + delegateName
	}
	return nodeDirs
}


func (this ClusterStructure) SaveAccountAndConfigFiles(nodeInfo *NodeInfo) {
	configDir := this.NodeDirs[nodeInfo.Name] + string(os.PathSeparator) + "config/"
	util.WriteFile(configDir, configDir + string(os.PathSeparator) + this.AccountFileName, nodeInfo.Account.ToPrettyJson())
	util.WriteFile(configDir, configDir + string(os.PathSeparator) + this.ConfigFileName, nodeInfo.Config.ToPrettyJson())
}

func (this ClusterStructure) getAccountFileLocation(nodeName string) string {
	baseDir := this.NodeDirs[nodeName]
	return baseDir  + string(os.PathSeparator) + this.AccountFileName
}

func (this ClusterStructure) getConfigFileLocation(nodeName string) string {
	baseDir := this.NodeDirs[nodeName]
	return baseDir  + string(os.PathSeparator) + this.ConfigFileName
}