package helper

import (
	"github.com/dispatchlabs/disgo/commons/types"
	"github.com/dispatchlabs/disgo/commons/utils"
	"os"
	"io/ioutil"
	"github.com/pkg/errors"
)

var genesisTransaction = `{"hash":"a48ff2bd1fb99d9170e2bae2f4ed94ed79dbc8c1002986f8054a369655e29276","type":0,"from":"e6098cc0d5c20c6c31c4d69f0201a02975264e94","to":"3ed25f42484d517cdfc72cafb7ebc9e8baa52c2c","value":10000000,"data":"","time":0,"signature":"03c1fdb91cd10aa441e0025dd21def5ebe045762c1eeea0f6a3f7e63b27deb9c40e08b656a744f6c69c55f7cb41751eebd49c1eedfbd10b861834f0352c510b200","hertz":0,"fromName":"","toName":""}`

func CreateSeedConfig(ipAddress string, httpPort, grpcPort int64, seedAccount *types.Account) *types.Config {
	seedConfig := createConfig(ipAddress, httpPort, grpcPort)

	seedConfig.Seeds = []*types.Node{
		{
			Address:      seedAccount.Address,
			GrpcEndpoint: seedConfig.GrpcEndpoint,
			HttpEndpoint: seedConfig.HttpEndpoint,
		},
	}
	return seedConfig
}

func CreateDelegateConfig(ipAddress string, httpPort, grpcPort int64, seedNodes []*types.Node) *types.Config {
	delegateConfig := createConfig(ipAddress, httpPort, grpcPort)
	delegateConfig.Seeds = seedNodes
	delegateConfig.IsBookkeeper = true
	return delegateConfig
}

func createConfig(ipAddress string, httpPort, grpcPort int64) *types.Config  {
	configInstance := &types.Config{
		HttpEndpoint:       &types.Endpoint{Host: ipAddress, Port: httpPort},
		GrpcEndpoint:       &types.Endpoint{Host: ipAddress, Port: grpcPort},
		LocalHttpApiPort:   0,
		DelegateAddresses:  []string{},
		GrpcTimeout:        5,
		GenesisTransaction: genesisTransaction,
	}
	return configInstance
}

func GetConfig(dirName, nodeName string) (*types.Config, error) {
	fileName := dirName + string(os.PathSeparator) + "config.json"
	if !utils.Exists(fileName) {
		return nil, errors.New("Config file does not exist")
	}
	bytes, err := ioutil.ReadFile(fileName)
	if err != nil {
		utils.Fatal("unable to read config.json", err)
	}

	config, err := types.ToConfigFromJson(bytes)
	if err != nil {
		utils.Fatal("unable to read config.json", err)
	}
	return config, nil
}