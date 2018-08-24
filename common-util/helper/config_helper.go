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
	"github.com/dispatchlabs/samples/common-util/configTypes"
)

var genesisTransaction = `{"hash":"a48ff2bd1fb99d9170e2bae2f4ed94ed79dbc8c1002986f8054a369655e29276","type":0,"from":"e6098cc0d5c20c6c31c4d69f0201a02975264e94","to":"3ed25f42484d517cdfc72cafb7ebc9e8baa52c2c","value":10000000,"data":"","time":0,"signature":"03c1fdb91cd10aa441e0025dd21def5ebe045762c1eeea0f6a3f7e63b27deb9c40e08b656a744f6c69c55f7cb41751eebd49c1eedfbd10b861834f0352c510b200","hertz":0,"fromName":"","toName":""}`

func CreateSeedAccount() *types.Account {
	return CreateAccount("seed")
}

func CreateDelegateAccount(delegateName string) *types.Account {
	return CreateAccount(delegateName)
}

func GetSeedConfig(ipAddress string, httpPort, grpcPort int64, seedAccount *types.Account) *types.Config {
	nodeName := "seed"

	configInstance := &types.Config{
		HttpEndpoint:       &types.Endpoint{Host: ipAddress, Port: httpPort},
		GrpcEndpoint:       &types.Endpoint{Host: ipAddress, Port: grpcPort},
		LocalHttpApiPort:   0,
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


func GetDelegateConfig(ipAddress string, httpPort, grpcPort int64, seedNodes []*types.Node) *types.Config {

	configInstance := &types.Config{
		HttpEndpoint:       &types.Endpoint{Host: ipAddress, Port: httpPort},
		GrpcEndpoint:       &types.Endpoint{Host: ipAddress, Port: grpcPort},
		LocalHttpApiPort:   0,
		GrpcTimeout:        5,
		GenesisTransaction: genesisTransaction,
		Seeds:				seedNodes,
		IsBookkeeper:       true,

	}
	return configInstance
}


func GetNewRemoteConfigs(seedNode *configTypes.NodeInfo, delegateNodes []*configTypes.NodeInfo) map[string]*configTypes.NodeInfo {
	configMap := map[string]*configTypes.NodeInfo{}

	seedAccount := CreateSeedAccount();
	seedConfig := GetSeedConfig(seedNode.Host, seedNode.HttpPort, seedNode.GrpcPort, seedAccount)
	seedNode.Account = seedAccount
	seedNode.Config = seedConfig
	configMap["seed"] = seedNode

	delegateAddressList := make([]string, len(delegateNodes))
	for i := 0; i < len(delegateNodes); i++ {
		delegateNodes[i].Config = GetDelegateConfig(delegateNodes[i].Host, delegateNodes[i].HttpPort, delegateNodes[i].GrpcPort, seedConfig.Seeds)
		delegateNodes[i].Account = CreateDelegateAccount(delegateNodes[i].Name)

		delegateAddressList[i] = delegateNodes[i].Account.Address

		configMap[delegateNodes[i].Name] = delegateNodes[i]
	}
	seedConfig.DelegateAddresses = delegateAddressList
	return configMap
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

