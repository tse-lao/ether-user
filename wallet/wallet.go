package wallet

import (
	"crypto/ecdsa"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

type Wallet struct {
	status     bool
	message    string
	PrivateKey string
	PublicKey  string
	Address    string
}

func check(e error) {
	if e != nil {
		fmt.Println("Error has been found here: ", e)
		return
	}
}

func CheckLoggedIn() bool {
	if FileExists("/tmp/private") || FileExists("/tmp/public") || FileExists("/tmp/address") {
		fmt.Println("All files exists and therefore we can start investigating")
		//now we need to check if it all matches.
		if MatchingKeys() {
			fmt.Print("We have matching keys")
			return true
		}

		return true
	}
	return false
}

func MatchingKeys() bool {

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

func ReadFile(path string) []byte {
	//here we will read the file

	data, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println("ERoor with reading the file")
	}

	fmt.Print(data)

	return data
}

func FileExists(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}

	return true
}

func Login(privateKey string) Wallet {
	response := Wallet{}
	if len(privateKey) > 32 {
		response.status = true
		response.message = "Valid Private key, lets try to use it. "
	} else {
		response.message = "Your private key entered does not exists, or is invalid"
		response.status = false
	}

	return response
}

func ValidPrivateKey(privateKey string) bool {

	return true
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
func WriteFile(path string, data []byte) string {
	//this will return a string with the message.

	fmt.Println("The directory is ", path)
	fmt.Println("The data to bes stored is ", data)

	//here we write the file and the path.
	err := os.WriteFile(path, data, 0644)
	check(err)

	return "Succesfully added the file to the following directory:" + path
}
