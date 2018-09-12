package helper

import (
	"net/http"
	"fmt"
	"os"
	"io/ioutil"
	"bytes"
	"github.com/dispatchlabs/disgo/commons/utils"
	"github.com/dispatchlabs/disgo/commons/types"
)

func GetQueue(endpoint string) string {
	response, err := http.Get(fmt.Sprintf("%s", endpoint))
	//var gossips []types.Gossip
	var result string
	if err != nil {
		fmt.Printf("%s", err)
		os.Exit(1)
	} else {
		defer response.Body.Close()
		contents, _ := ioutil.ReadAll(response.Body)
		//receipt, err = unmarshalReceipt(contents)
		if err != nil {
			fmt.Printf("%s", err)
			os.Exit(1)
		}
		result = string(contents)
		//fmt.Printf("%s\n", receipt.ToPrettyJson())
	}
	return result

}

func PostTx(tx *types.Transaction, endpoint string) {
	fmt.Printf("PostTx(): %s\n%s\n\n", endpoint, tx.ToPrettyJson())
	data := new(bytes.Buffer)
	data.WriteString(tx.String())

	response, err := http.Post(
		endpoint,
		"application/json; charset=utf-8",
		data,
	)
	if err != nil {
		utils.Error(err)
		return
	}
	contents, _ := ioutil.ReadAll(response.Body)
	// If NOT then this happens https://stackoverflow.com/questions/37454236/net-http-server-too-many-open-files-error
	fmt.Printf("Response: %v\n", string(contents))
	response.Body.Close()
}

