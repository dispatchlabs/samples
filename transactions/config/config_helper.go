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

type TestConfig struct {
	Account 			*types.Account
	SeedConfig 			*types.Config
	SeedAddress 		string
	IsSeed				bool
	HttpEndpoint		*types.Endpoint
	GrpcEndpoint		*types.Endpoint
	GenesisTx			string
}

var genesisTransaction = `{"hash":"a48ff2bd1fb99d9170e2bae2f4ed94ed79dbc8c1002986f8054a369655e29276","type":0,"from":"e6098cc0d5c20c6c31c4d69f0201a02975264e94","to":"3ed25f42484d517cdfc72cafb7ebc9e8baa52c2c","value":10000000,"data":"","time":0,"signature":"03c1fdb91cd10aa441e0025dd21def5ebe045762c1eeea0f6a3f7e63b27deb9c40e08b656a744f6c69c55f7cb41751eebd49c1eedfbd10b861834f0352c510b200","hertz":0,"fromName":"","toName":""}`

var Delegate_1 = &TestConfig{IsSeed: false, HttpEndpoint: &types.Endpoint{Host: "127.0.0.1", Port: 1175}, GrpcEndpoint: &types.Endpoint{Host: "127.0.0.1", Port: 1173}, GenesisTx: genesisTransaction}
var Delegate_2 = &TestConfig{IsSeed: false, HttpEndpoint: &types.Endpoint{Host: "127.0.0.1", Port: 1275}, GrpcEndpoint: &types.Endpoint{Host: "127.0.0.1", Port: 1273}, GenesisTx: genesisTransaction}
var Delegate_3 = &TestConfig{IsSeed: false, HttpEndpoint: &types.Endpoint{Host: "127.0.0.1", Port: 1375}, GrpcEndpoint: &types.Endpoint{Host: "127.0.0.1", Port: 1373}, GenesisTx: genesisTransaction}
var Delegate_4 = &TestConfig{IsSeed: false, HttpEndpoint: &types.Endpoint{Host: "127.0.0.1", Port: 1475}, GrpcEndpoint: &types.Endpoint{Host: "127.0.0.1", Port: 1473}, GenesisTx: genesisTransaction}
var Seed = &TestConfig{IsSeed: true, HttpEndpoint: &types.Endpoint{Host: "127.0.0.1", Port: 1975}, GrpcEndpoint: &types.Endpoint{Host: "127.0.0.1", Port: 1973}, GenesisTx: genesisTransaction}

func SetUp(nbrDelegates int, startingPort int64) []*TestConfig {
	nodeName := "seed"
	dir := GetConfigDir(nodeName)
	seedAccount := GetAccount(dir, nodeName)
	fmt.Printf("\nSeed Account: %s\n", seedAccount.Address)
	seedConfig := GetConfig(dir, Seed)

	for i := 1; i <= nbrDelegates; i++ {
		delegateName := fmt.Sprintf("delegate-%d", i)
		startingPort++
		grpcPort := startingPort
		startingPort++
		httpPort := startingPort

		setupDelegate(delegateName, seedAccount.Address, httpPort, grpcPort, seedConfig)

	}
	fmt.Printf("%s\n %s\n", seedAccount.ToPrettyJson(), seedConfig.String())

	configs := []*TestConfig{Delegate_1, Delegate_2, Delegate_3, Delegate_4}
	return configs
}

func setupDelegate(delegateName, seedAddress string, httpPort, grpcPort int64, seedConfig *types.Config) {
	delegateConfig := &TestConfig{IsSeed: false, HttpEndpoint: &types.Endpoint{Host: "127.0.0.1", Port: httpPort}, GrpcEndpoint: &types.Endpoint{Host: "127.0.0.1", Port: grpcPort}, GenesisTx: genesisTransaction}

	dir := GetConfigDir(delegateName)
	GetAccount(dir, delegateName)
	delegateConfig.SeedAddress = seedAddress
	delegateConfig.SeedConfig = seedConfig
	GetConfig(dir, delegateConfig)
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

// GetConfig -
func GetConfig(dirName string, testConfig *TestConfig) *types.Config {
	fileName := dirName + string(os.PathSeparator) + "config.json"
	if !utils.Exists(fileName) {

		configInstance := &types.Config{
			HttpEndpoint:       testConfig.HttpEndpoint,
			GrpcEndpoint:       testConfig.GrpcEndpoint,
			GrpcTimeout:        5,
			GenesisTransaction: testConfig.GenesisTx,
		}
		if !testConfig.IsSeed {
			configInstance.SeedAddresses = []string{testConfig.SeedAddress}
			configInstance.IsBookkeeper = true
			configInstance.SeedEndpoints = []*types.Endpoint{
				{
					Host: testConfig.SeedConfig.GrpcEndpoint.Host,
					Port: testConfig.SeedConfig.GrpcEndpoint.Port,
				},
			}

		} else {
			configInstance.SeedEndpoints = []*types.Endpoint{
				{
					Host: testConfig.GrpcEndpoint.Host,
					Port: testConfig.GrpcEndpoint.Port,
				},
			}
		}
		// Write config.
		file, err := os.Create(fileName)
		defer file.Close()
		if err != nil {
			utils.Fatal(fmt.Sprintf("unable to write %s", fileName), err)
		}
		fmt.Printf(configInstance.ToPrettyJson())

		fmt.Fprintf(file, configInstance.ToPrettyJson())
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
