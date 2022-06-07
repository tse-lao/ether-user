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
	/* d1 := []byte("Hello, world\n")
	writeFile("/tmp/keys", d1)
	readFile("/tmp/keys")
	*/
	//here we need to make sure that run afunction that checks if hte uesrs exist.
	if !wallet.CheckLoggedIn() {
		wallet.CreateEthWallet()
	}

	loggedIn := wallet.CheckLoggedIn()
	fmt.Println("User is logged in : ", loggedIn)
	//read file for testing
}
