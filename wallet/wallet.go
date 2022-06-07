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
	privateKey string
	publicKey  string
	address    string
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
