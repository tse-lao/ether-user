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

	publicKey := wallet.GetPublicKey("somepassword")

	fmt.Println(publicKey)
	data := []byte("encrypt this data please")

	fmt.Println("We encrypted data below with public key: \n\n ")
	encrypted := wallet.EncryptWithPublicKey(publicKey, data)

	fmt.Println("We descrypted with private key: \n\n ")
	wallet.DecryptWithPrivateKey("somepassword", encrypted)
}

func checking() {

	fmt.Println("== STARTING NEW ACCOUNT CREATION! ==")
	//wallet.CreateNewAccount("testing")
	fmt.Println("== GETTING ALL THEW DETAILS ABOUT THE WALLET ==")
	filePath := "./wallets/UTC--2022-06-12T08-10-20.631620000Z--36799a0a2e657721be34a92ff2e74c7261ee735f"
	info := wallet.GetAccount(filePath, "testing")
	fmt.Println(info)
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
