package wallet

import (
	"bytes"
	"fmt"
	"testing"
)

func TestEncryptDescrypt(t *testing.T) {

	//GET PUBLIC KEY
	publicKey := GetPublicKey("somepassword")

	fmt.Println(publicKey)
	data := []byte("encrypt this data please")

	fmt.Println("We encrypted data below with public key: \n\n ")
	encrypted := EncryptWithPublicKey(publicKey, data)

	fmt.Println("We descrypted with private key: \n\n ")
	decrypted := DecryptWithPrivateKey("somepassword", encrypted)

	if !bytes.Equal(encrypted, decrypted) {
		t.Fatal("the key pairs do not match")
	}

}
