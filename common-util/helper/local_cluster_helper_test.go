package helper

import (
	"testing"
	"github.com/dispatchlabs/samples/common-util/configTypes"
	"fmt"
	"io/ioutil"
	"log"
)

func TestLocalFullConfig(t *testing.T) {
	host := "127.0.0.1"
	nbrDelegates := 5
	var startingPort int64
	startingPort = 3501

	seedNode := &configTypes.NodeInfo{"seed", host, 1975,1973, nil, nil}

	delegateNodes := make([]*configTypes.NodeInfo, nbrDelegates)
	for i := 0; i < nbrDelegates; i++ {
		delegateName := fmt.Sprintf("delegate-%d", i)

		startingPort++
		grpcPort := startingPort
		startingPort++
		httpPort := startingPort

		delegateNodes[i] = &configTypes.NodeInfo{delegateName, host, httpPort, grpcPort, nil, nil}

	}
	configMap := CreateNewLocalConfigs(seedNode, delegateNodes, true)
	for _, v := range configMap {
		fmt.Printf("%s\n", v.ToPrettyJson())
	}

}

func TestUpdateDisgoExecutable(t *testing.T) {
	files, err := ioutil.ReadDir(GetDefaultDirectory())
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range files {
		if file.IsDir() {

			UpdateDisgoExecutable(file.Name())
		}
	}
}