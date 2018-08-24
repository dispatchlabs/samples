package configHelpers

import (
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"log"
	"math/big"
	"os"
	"os/user"
	"time"

	"github.com/dispatchlabs/disgo/commons/crypto"
	"github.com/dispatchlabs/disgo/commons/types"
	"github.com/dispatchlabs/disgo/commons/utils"
)

// ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~
// Account
// ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~
func CreateAccount(nodeName string) *types.Account {
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

	return account
}

// ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~
// Config -
// ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~
func GetSeedConfig(ipAddress string, httpPort, grpcPort int64, seedAccounts []*types.Account) *types.Config {
	configInstance := types.GetDefaultConfig()

	configInstance.HttpEndpoint = &types.Endpoint{Host: ipAddress, Port: httpPort}
	configInstance.GrpcEndpoint = &types.Endpoint{Host: ipAddress, Port: grpcPort}

	seeds := []*types.Node{}
	for i := 0; i < len(seedAccounts); i++ {
		seeds = append(seeds, &types.Node{
			Address:      seedAccounts[i].Address,
			GrpcEndpoint: configInstance.GrpcEndpoint,
			HttpEndpoint: configInstance.HttpEndpoint,
		})
	}

	configInstance.Seeds = seeds
	return configInstance
}

func GetSeedConfigAndSave(ipAddress string, httpPort, grpcPort int64, seedAccount []*types.Account) *types.Config {
	nodeName := "seed"
	seedConfig := GetSeedConfig(ipAddress, httpPort, grpcPort, seedAccount)
	SaveConfig(nodeName, seedConfig)

	return seedConfig
}

func GetDelegateConfig(ipAddress string, httpPort, grpcPort int64, seedNodes []*types.Node) *types.Config {
	configInstance := types.GetDefaultConfig()

	configInstance.HttpEndpoint = &types.Endpoint{Host: ipAddress, Port: httpPort}
	configInstance.GrpcEndpoint = &types.Endpoint{Host: ipAddress, Port: grpcPort}
	configInstance.Seeds = seedNodes

	return configInstance
}

func GetNewRemoteConfigs(seedNode *NodeInfo, delegateNodes []*NodeInfo) map[string]*NodeInfo {
	configMap := map[string]*NodeInfo{}

	seedAccount := CreateAccount("seed")
	seedConfig := GetSeedConfig(seedNode.Host, seedNode.HttpPort, seedNode.GrpcPort, []*types.Account{seedAccount})
	seedNode.Account = seedAccount
	seedNode.Config = seedConfig
	configMap["seed"] = seedNode

	delegateAddressList := make([]string, len(delegateNodes))
	for i := 0; i < len(delegateNodes); i++ {
		delegateNodes[i].Config = GetDelegateConfig(delegateNodes[i].Host, delegateNodes[i].HttpPort, delegateNodes[i].GrpcPort, seedConfig.Seeds)
		delegateNodes[i].Account = CreateAccount(delegateNodes[i].Name)

		delegateAddressList[i] = delegateNodes[i].Account.Address

		configMap[delegateNodes[i].Name] = delegateNodes[i]
	}
	seedConfig.DelegateAddresses = delegateAddressList
	return configMap
}

func GetConfigDir(nodeName string) string {
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
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

func writeAccount(dirName string, account *types.Account) error {
	fileName := dirName + string(os.PathSeparator) + "account.json"

	// Write account.
	file, err := os.Create(fileName)
	defer file.Close()
	if err != nil {
		utils.Fatal(fmt.Sprintf("unable to write %s", fileName), err)
	}
	fmt.Fprintf(file, account.ToPrettyJson())

	return nil
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
	return &types.Endpoint{Host: "127.0.0.1", Port: port}
}
