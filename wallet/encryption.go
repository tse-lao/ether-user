package wallet

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

//this creates something new.
func CreateNewAccount(password string) {

	ks := keystore.NewKeyStore("./wallets", keystore.StandardScryptN, keystore.StandardScryptP)

	fmt.Println("Starting the keystore initializing process")

	fmt.Println("finished the keystore intializing process. ")
	account, err := ks.NewAccount(password)
	if err != nil {
		log.Fatal(err)

	}

	fmt.Println(account.Address.Hex())
	result := ks.Wallets()

	fmt.Println("Wallets")
	fmt.Println(result)

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
	fmt.Println("\n\nAddress:\n", address)
	fmt.Println("\n Finishing all the private key data and stuff. ")

	returnWallet.message = "We successfully retrieved your account"
	returnWallet.status = true
	returnWallet.Address = address

	return returnWallet

}
