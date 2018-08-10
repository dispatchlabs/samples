package main

import (
	"fmt"

	disgoSdk "github.com/dispatchlabs/disgo/sdk"
)

func main() {
	account, err := disgoSdk.CreateAccount("")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(account.String())
	}
}
