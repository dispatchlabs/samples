package configTypes

import (
	"fmt"
	"os"
)

type ClusterStructure struct {
	DisgoDir		string
	ClusterRoot 	string
	SeedDirs		[]string
	DelegateDirs 	[]string
	AccountFileName string
	ConfigFileName  string
}

func NewClusterStructure(disgoDir, clusterRoot string, nbrSeeds, nbrDelegates int) *ClusterStructure {

	return &ClusterStructure{
		DisgoDir:		disgoDir,
		ClusterRoot:	clusterRoot,
		SeedDirs:		getSeedDirs(clusterRoot, nbrSeeds),
		DelegateDirs:   getDelegateDirs(clusterRoot, nbrDelegates),
		AccountFileName: "account.json",
		ConfigFileName: "config.json",
	}
}

func getSeedDirs(clusterRoot string, nbrSeeds int) []string {
	seedDirs := make([]string, nbrSeeds)
	for i := 0; i < nbrSeeds; i++ {
		seedName := fmt.Sprintf("seed-%d", i)
		seedDirs[i] = clusterRoot + string(os.PathSeparator) + seedName
	}
	return seedDirs
}

func getDelegateDirs(clusterRoot string, nbrDelegates int) []string {
	delegateDirs := make([]string, nbrDelegates)
	for i := 0; i < nbrDelegates; i++ {
		delegateName := fmt.Sprintf("delegate-%d", i)
		delegateDirs[i] = clusterRoot + string(os.PathSeparator) + delegateName
	}
	return delegateDirs
}

//func (this ClusterStructure) getAccountFileLocation(dirName string) {
//	this.
//}