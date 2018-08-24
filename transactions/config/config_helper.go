package config

import (
	"fmt"
	"github.com/dispatchlabs/disgo/commons/types"
	"github.com/dispatchlabs/disgo/commons/utils"
	"os"
	"io/ioutil"
	"github.com/dispatchlabs/disgo/commons/crypto"
	"encoding/hex"
	"math/big"
	"time"
	"os/user"
	"log"
)

var genesisTransaction = `{"hash":"a48ff2bd1fb99d9170e2bae2f4ed94ed79dbc8c1002986f8054a369655e29276","type":0,"from":"e6098cc0d5c20c6c31c4d69f0201a02975264e94","to":"3ed25f42484d517cdfc72cafb7ebc9e8baa52c2c","value":10000000,"data":"","time":0,"signature":"03c1fdb91cd10aa441e0025dd21def5ebe045762c1eeea0f6a3f7e63b27deb9c40e08b656a744f6c69c55f7cb41751eebd49c1eedfbd10b861834f0352c510b200","hertz":0,"fromName":"","toName":""}`
var host = "127.0.0.1"

func SetUp(nbrDelegates int, startingPort int64) {

	seedConfig := configSeed()

	for i := 1; i <= nbrDelegates; i++ {
		delegateName := fmt.Sprintf("delegate-%d", i)
		startingPort++
		apiPort := startingPort
		startingPort++
		grpcPort := startingPort
		startingPort++
		httpPort := startingPort


		configDelegate(delegateName, int(apiPort), httpPort, grpcPort, seedConfig.Seeds)

	}
}

func configSeed() *types.Config {
	nodeName := "seed"
	dir := GetConfigDir(nodeName)
	seedAccount := GetAccount(dir, nodeName)
	fmt.Printf("\nSeed Account: %s\n", seedAccount.Address)

	configInstance := &types.Config{
		HttpEndpoint:       newEndpoint(1975),
		GrpcEndpoint:       newEndpoint(1973),
		LocalHttpApiPort:   1971,
		DelegateAddresses:  []string{},
		GrpcTimeout:        5,
		GenesisTransaction: genesisTransaction,
	}
	seeds := []*types.Node{
		{
			GrpcEndpoint: configInstance.GrpcEndpoint,
			HttpEndpoint: configInstance.HttpEndpoint,
		},
	}
	configInstance.Seeds = seeds
	SaveConfig(nodeName, configInstance)
	configInstance.Seeds[0].Address = seedAccount.Address
	return configInstance
}


func configDelegate(delegateName string, apiPort int, httpPort, grpcPort int64, seedNodes []*types.Node) {

	configInstance := &types.Config{
		HttpEndpoint:       newEndpoint(httpPort),
		GrpcEndpoint:       newEndpoint(grpcPort),
		LocalHttpApiPort:   apiPort,
		GrpcTimeout:        5,
		GenesisTransaction: genesisTransaction,
		Seeds:				seedNodes,
		IsBookkeeper:       true,

	}
	dir := GetConfigDir(delegateName)
	GetAccount(dir, delegateName)
	SaveConfig(delegateName, configInstance)
}

func GetConfigDir(nodeName string) string {
	usr, err := user.Current()
	if err != nil {
		log.Fatal( err )
	}

	rootDir := usr.HomeDir + "/go/src/github.com/dispatchlabs/samples/run-nodes-locally"

	directoryName := rootDir + string(os.PathSeparator) + nodeName + string(os.PathSeparator) + "config"
	if !utils.Exists(directoryName) {
		err := os.MkdirAll(directoryName, 0755)
		if err != nil {
			utils.Error(fmt.Sprintf("unable to create directory %s", directoryName), err)
			panic(err)
		}
	}
	return directoryName
}

func GetAccount(dirName, nodeName string) *types.Account {
	fileName := dirName + string(os.PathSeparator) + "account.json"
	if !utils.Exists(fileName) {
		publicKey, privateKey := crypto.GenerateKeyPair()
		address := crypto.ToAddress(publicKey)

		account := &types.Account{}
		account.Address = hex.EncodeToString(address)
		account.PrivateKey = hex.EncodeToString(privateKey)
		account.Balance = big.NewInt(0)
		account.Name = nodeName
		now := time.Now()
		account.Created = now
		account.Updated = now

		// Write account.
		file, err := os.Create(fileName)
		defer file.Close()
		if err != nil {
			utils.Fatal(fmt.Sprintf("unable to write %s", fileName), err)
		}
		fmt.Fprintf(file, account.ToPrettyJson())
	}
	bytes, err := ioutil.ReadFile(fileName)
	if err != nil {
		utils.Fatal("unable to read account.json", err)
	}
	account, err := types.ToAccountFromJson(bytes)
	if err != nil {
		utils.Fatal("unable to read account.json", err)
	}
	return account
}

func SaveConfig(nodeName string, newConfig *types.Config) *types.Config {
	fileName := GetConfigDir(nodeName) + string(os.PathSeparator) + "config.json"
	if !utils.Exists(fileName) {
		// Write config.
		file, err := os.Create(fileName)
		defer file.Close()
		if err != nil {
			utils.Fatal(fmt.Sprintf("unable to write %s", fileName), err)
		}
		fmt.Printf(newConfig.ToPrettyJson())

		fmt.Fprintf(file, newConfig.ToPrettyJson())
	}
	bytes, err := ioutil.ReadFile(fileName)
	if err != nil {
		utils.Fatal("unable to read config.json", err)
	}
	config, err := types.ToConfigFromJson(bytes)
	if err != nil {
		utils.Fatal("unable to read config.json", err)
	}
	return config
}

func newEndpoint(port int64) *types.Endpoint {
	return &types.Endpoint{Host: host, Port: port}
}

