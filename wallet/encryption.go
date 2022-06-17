package wallet

import (
	"crypto/ecdsa"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/crypto/ecies"
)

//this creates something new.
func CreateNewAccount(password string) string {

	ks := keystore.NewKeyStore("./wallets", keystore.StandardScryptN, keystore.StandardScryptP)

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
func Accounts() {
	fmt.Print("Not working yet, need to implement a retrieval process of the saving.")
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
	b, err := ioutil.ReadFile(file)

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
	fmt.Println(hexutil.Encode(pData))
	returnWallet.PublicKey = hexutil.Encode(pData)

	address := crypto.PubkeyToAddress(key.PrivateKey.PublicKey)
	fmt.Println("\n\nAddress:\n", address.String())
	fmt.Println("\n Finishing all the private key data and stuff. ")

	returnWallet.message = "We successfully retrieved your account"
	returnWallet.status = true
	returnWallet.Address = address.String()

	return returnWallet

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

	EncryptWithPublicKey(publicKey, "This data will be encrypted like a legend")
}

//behind the back gathering of privatekey.
func GetPrivateKey(password string) *ecdsa.PrivateKey {
	//privateKey, err := crypto.HexToECDSA("fad9c8855b740a0b7ed4c221dbad0f33a83a49cad6b3fe8d5817ac83d38b6a19")

	//retrieve the file is done by
	file := RetrieveWalletFile()

	b, err := ioutil.ReadFile(file)

	if err != nil {
		log.Fatal(err)
	}

	key, err := keystore.DecryptKey(b, password)

	//returnKey := hexutil.Encode(pData)

	//this private key can also be retrieved from the file and string.

	return key.PrivateKey
}

func GetPublicKey(password string) *ecdsa.PublicKey {
	file := RetrieveWalletFile()
	b, err := ioutil.ReadFile(file)

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

func RetrieveWalletFile() string {
	files, err := ioutil.ReadDir("./wallets")

	if err != nil {
		log.Fatal("Unable to read the file", err)
	}
	return "./wallets/" + files[0].Name()
}

//TODO: create function that verifies that this is the sender of the data.
func VerifyUser() bool {
	return false
}

func EncryptWithPublicKey(publicKey *ecdsa.PublicKey, data string) {

	acceptableKey := ecies.ImportECDSAPublic(publicKey)
	rdr := strings.NewReader(data)

	//encryptBytes, err := ecdsa.(sha256.New(), rand.Reader, publicKey, data, nil)

	//TODO: something here is not working correctly
	result, err := ecies.Encrypt(rdr, acceptableKey, nil, nil, nil)

	if err != nil {
		fmt.Println("We got the following error in encrypting the public key:")
		fmt.Println(err)
	}
	fmt.Println(result)
}
