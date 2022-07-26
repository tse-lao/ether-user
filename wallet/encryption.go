package wallet

import (
	"crypto/ecdsa"
	"crypto/rand"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/crypto/ecies"
)

var filePath = "./keystore/"

func CreateKeyStore() *keystore.KeyStore {
	//check if there is a keystore available.
	ks := keystore.NewKeyStore(filePath, keystore.StandardScryptN, keystore.StandardScryptP)

	return ks
}

//this creates something new.
func CreateNewAccount(password string) string {

	ks := CreateKeyStore()
	//check ENVIRONMENT.
	//create a new account here,
	account, err := ks.NewAccount(password)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(account.Address)
	result := ks.Wallets()

	fmt.Println("Wallets")
	fmt.Println(result)

	return account.Address.String()
}

//GEt aLL tHe acCOunts
func GetAccounts() interface{} {
	fmt.Print("Not working yet, need to implement a retrieval process of the saving.")
	ks := CreateKeyStore()
	accounts := ks.Accounts()

	return accounts
}

func importKs() {
	file := "./wallets/UTC--2022-06-11T12-20-12.532168000Z--0a934422f86f899cd02983edd79825e5ac6f9db2"
	ks := keystore.NewKeyStore("./tmp", keystore.StandardScryptN, keystore.StandardScryptP)
	jsonBytes, err := ioutil.ReadFile(file)

	if err != nil {
		log.Fatal(err)
	}

	password := "secret"
	account, err := ks.Import(jsonBytes, password, password)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(account.Address.Hex()) // 0x20F8D42FB0F667F2E53930fed426f225752453b3

	if err := os.Remove(file); err != nil {
		log.Fatal(err)
	}
}

func GetAccount(file string, password string) Wallet {
	returnWallet := Wallet{}

	b, err := ioutil.ReadFile(file)

	if err != nil {
		log.Fatal(err)
		returnWallet.Status = false
		returnWallet.Message = "An error occured when searching for the file.. "
	}

	key, err := keystore.DecryptKey(b, password)

	if err != nil {
		log.Fatal(err)
		returnWallet.Status = false
		returnWallet.Message = "Your password is incorrect!, or another error occured"
	}

	pData := crypto.FromECDSA(key.PrivateKey)

	fmt.Println("Print PRIVATE KEY:: ")
	fmt.Println(hexutil.Encode(pData))
	returnWallet.PrivateKey = hexutil.Encode(pData)

	pData = crypto.FromECDSAPub(&key.PrivateKey.PublicKey)
	fmt.Println(hexutil.Encode(pData))
	returnWallet.PublicKey = hexutil.Encode(pData)

	address := crypto.PubkeyToAddress(key.PrivateKey.PublicKey)
	fmt.Println("\n\nAddress:\n", address.String())

	returnWallet.Message = "We successfully retrieved your account"
	returnWallet.Status = true
	returnWallet.Address = address.String()

	return returnWallet

}

func RetrievePrivateKey(address []byte, password string) *keystore.Key {
	//byte[] ass argument.
	//this is not possible
	key, err := keystore.DecryptKey(address, password)

	if err != nil {
		fmt.Println("We have and")
		fmt.Println(err)
	}

	fmt.Println("Private key is: " + hexutil.Encode(crypto.FromECDSA(key.PrivateKey)))

	return key

}
func EncryptData(data []byte) {

	hash := crypto.Keccak256Hash(data)

	fmt.Println(hash.Hex())

	privateKey := GetPrivateKey("testing")

	signature, err := crypto.Sign(hash.Bytes(), privateKey)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(hexutil.Encode(signature))

	fmt.Println(privateKey)

	publicKey := GetPublicKey("testing")

	fmt.Println("==== PUBLIC KEY ====")
	fmt.Println(publicKey)

	EncryptWithPublicKey(publicKey, data)
}

//behind the back gathering of privatekey.
func GetPrivateKey(password string) *ecdsa.PrivateKey {
	//privateKey, err := crypto.HexToECDSA("fad9c8855b740a0b7ed4c221dbad0f33a83a49cad6b3fe8d5817ac83d38b6a19")

	//retrieve the file is done by
	//this is not possible.,
	file := RetrieveWalletFile()

	b, err := ioutil.ReadFile(file.Message)

	if err != nil {
		log.Fatal(err)
	}

	key, err := keystore.DecryptKey(b, password)

	fmt.Println(key)

	//returnKey := hexutil.Encode(pData)

	//this private key can also be retrieved from the file and string.

	return key.PrivateKey
}

func GetPublicKey(password string) *ecdsa.PublicKey {
	file := RetrieveWalletFile()
	b, err := ioutil.ReadFile(file.Message)

	if err != nil {
		log.Fatal(err)
	}
	returnWallet := Wallet{}

	key, err := keystore.DecryptKey(b, password)

	pData := crypto.FromECDSA(key.PrivateKey)

	fmt.Println("Print PRIVATE KEY:: ")
	fmt.Println(hexutil.Encode(pData))
	returnWallet.PrivateKey = hexutil.Encode(pData)
	fmt.Println("\n Print PUBLIC_KEY::")

	pData = crypto.FromECDSAPub(&key.PrivateKey.PublicKey)

	return &key.PrivateKey.PublicKey
}

func RetrieveWalletFile() Notification {
	note := Notification{}
	files, err := ioutil.ReadDir(filePath)

	if err != nil {
		log.Fatal("Unable to read the file", err)
		note.Status = false
		note.Message = "Unable to read the file "
	}
	note.Message = "path: " + filePath + files[0].Name() + ""
	note.Status = true
	return note
}

//TODO: create function that verifies that this is the sender of the data.
func VerifyUser() bool {
	return false
}

func EncryptWithPublicKey(publicKey *ecdsa.PublicKey, data []byte) []byte {

	acceptableKey := ecies.ImportECDSAPublic(publicKey)

	//encryptBytes, err := ecdsa.(sha256.New(), rand.Reader, publicKey, data, nil)

	//TODO: something here is not working correctly
	result, err := ecies.Encrypt(rand.Reader, acceptableKey, data, nil, nil)

	if err != nil {
		fmt.Println("We got the following error in encrypting the public key:")
		fmt.Println(err)
	}
	fmt.Println(result)

	return result
}

func DecryptWithPrivateKey(password string, data []byte) []byte {
	privateKey := GetPrivateKey(password)
	private := ecies.ImportECDSA(privateKey)

	result, err := private.Decrypt(data, nil, nil)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(result))

	return result
}
