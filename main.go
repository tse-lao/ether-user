package main

import (
	"fmt"

	"github.com/tse-lao/ether-user/wallet"
)

func check(e error) {
	if e != nil {
		fmt.Println("Error has been found here: ", e)
		return
	}
}

func main() {
	//initializing
	//wallet.CreateNewAccount("somepassword")
	//wallet.GetPrivateKey("somepassword")

	//[] Make sure that we get the public key correctly.

	//wallet.CreateNewAccount("password")

	result := wallet.GetAccount("password")

	fmt.Println(result)
}
