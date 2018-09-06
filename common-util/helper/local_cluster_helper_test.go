package helper

import (
	"testing"
	"github.com/dispatchlabs/samples/common-util/configTypes"
	"fmt"
)

func TestLocalFullConfig(t *testing.T) {
	host := "127.0.0.1"
	nbrSeeds := 1
	nbrDelegates := 5
	delegateStartingPort := 3502
	seedStartingPort := 1973

	SetupDefaultConfig(host, nbrSeeds, nbrDelegates, seedStartingPort, delegateStartingPort)

}

func setup(host string, nbrSeeds, nbrDelegates, seedStartingPort, delegateStartingPort int) {
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
