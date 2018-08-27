package helper

import (
	"github.com/dispatchlabs/disgo/commons/types"
	"github.com/dispatchlabs/disgo/commons/crypto"
	"encoding/hex"
	"math/big"
	"time"
	"os"
	"github.com/dispatchlabs/disgo/commons/utils"
	"github.com/dispatchlabs/samples/common-util/util"
	"io/ioutil"
)

func CreateSeedAccount() *types.Account {
	return CreateAccount("seed")
}

func CreateDelegateAccount(delegateName string) *types.Account {
	return CreateAccount(delegateName)
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

func GetAccount(dirName, nodeName string) *types.Account {
	fileName := dirName + string(os.PathSeparator) + "account.json"
	var account *types.Account
	if !utils.Exists(fileName) {
		account = CreateAccount(nodeName)
		util.WriteFile(dirName, fileName, account.ToPrettyJson())
	}
	bytes, err := ioutil.ReadFile(fileName)
	if err != nil {
		utils.Fatal("unable to read account.json", err)
	}
	account, err = types.ToAccountFromJson(bytes)
	if err != nil {
		utils.Fatal("unable to read account.json", err)
	}
	return account
}
