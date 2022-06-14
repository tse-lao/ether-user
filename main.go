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

	fmt.Println("== STARTING NEW ACCOUNT CREATION! ==")

	//wallet.CreateNewAccount("testing")

	fmt.Println("== GETTING ALL THEW DETAILS ABOUT THE WALLET ==")

	filePath := "./wallets/UTC--2022-06-12T08-10-20.631620000Z--36799a0a2e657721be34a92ff2e74c7261ee735f"

	wallet.GetAccount(filePath, "testing")
	/* d1 := []byte("Hello, world\n")
	writeFile("/tmp/keys", d1)
	readFile("/tmp/keys")
	*/
	//here we need to make sure that run afunction that checks if hte uesrs exist.
	//result := wallet.Login("1234567890-12345678901234567890-1234567890-")

	//fmt.Println(result)
	// loggedIn := wallet.CheckLoggedIn()
	// fmt.Println("User is logged in : ", loggedIn)
	//read file for testing
}
