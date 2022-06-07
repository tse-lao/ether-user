package ether_user

import (
	"crypto/ecdsa"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
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
	if !AlreadyLoggedIn() {
		CreateEthWallet()
	}

	//read file for testing
}

// now we want some kind of function that can create a path, data to be written and store like a propper json file
func WriteFile(path string, data []byte) string {
	//this will return a string with the message.

	fmt.Println("The directory is ", path)
	fmt.Println("The data to bes stored is ", data)

	//here we write the file and the path.
	err := os.WriteFile(path, data, 0644)
	check(err)

	return "Succesfully added the file to the following directory:" + path
}

func ReadFile(path string) []byte {
	//here we will read the file

	data, err := ioutil.ReadFile(path)
	check(err)
	fmt.Print(data)

	return data
}

//now we need to  create a function that is providing somehting new, namely, the interaction with a wallet of rnot.
type Wallet struct {
	PrivateKey string
	PublicKey  string
	Address    string
}

func SetupCheck() string {
	return "Succesfully called the ether-user package"
}

func CreateEthWallet() {
	newWallet := Wallet{}

	//Create private Key
	privateKey, err := crypto.GenerateKey()

	check(err)
	privateKeyBytes := crypto.FromECDSA(privateKey)

	fmt.Println("\nPrivate Key Bytes \n", hexutil.Encode(privateKeyBytes))
	newWallet.PrivateKey = hexutil.Encode(privateKeyBytes)

	//Create Public Key
	publicKey := privateKey.Public()

	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)

	if !ok {
		fmt.Println("error casting public key to ECDSA")
	}

	publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)
	fmt.Println("\n\nPublic key in bytes:\n", hexutil.Encode(publicKeyBytes))
	newWallet.PublicKey = hexutil.Encode(publicKeyBytes)

	//Create Address
	address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()
	fmt.Println("\n\nAddress:\n", address)

	newWallet.Address = address

	fmt.Println("\n\n", newWallet)

	//now we write all the files to towards the system to check if they exists and correspond.

	// we will be writing them as bytes to store it correctl.y

	WriteFile("/tmp/public", publicKeyBytes)
	WriteFile("/tmp/private", privateKeyBytes)
	WriteFile("/tmp/address", []byte(address))

}

func FileExists(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}

	return true
}

//Cheack if we are already logged in.
func AlreadyLoggedIn() bool {
	if FileExists("/tmp/private") || FileExists("/tmp/public") || FileExists("/tmp/address") {
		fmt.Println("All files exists and therefore we can start investigating")
		//now we need to check if it all matches.
		if matchingKeys() {
			fmt.Print("We have matching keys")
			return true
		}

		return true
	}
	return false
}

func matchingKeys() bool {

	privateKeyBytes := ReadFile("/tmp/private")
	publicKeyBytes := ReadFile("/tmp/public")
	//addressBytes := readFile("/tmp/address")

	//we check if the privateKey is corresponding to the poublic bytes key
	storedPrivateKey, _ := crypto.ToECDSA(privateKeyBytes)

	publicKey := storedPrivateKey.Public()

	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)

	if !ok {

		fmt.Println("error casting public key to ECDSA")
	}

	storedPublicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)

	if hexutil.Encode(storedPublicKeyBytes) == hexutil.Encode(publicKeyBytes) {
		fmt.Println("That is nice the private key and public key are corresponding.")
		return true
	}

	return false
}
