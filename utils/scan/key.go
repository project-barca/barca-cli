package scan

import (
	"crypto/dsa"
	"crypto/rsa"
	"encoding/gob"
	"fmt"
	"os"
)

func pubKeyDSA(publicKey string) {
	//Public key DSA
	pubKeyFile, err := os.Open(publicKey)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	decoder := gob.NewDecoder(pubKeyFile)

	var publickey dsa.PublicKey
	err = decoder.Decode(&publickey)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	pubKeyFile.Close()
	fmt.Printf("Public key: \n%x\n", pubKeyFile)

	fmt.Printf("Public key parameter P: %v\n", publickey.Parameters.P)
	fmt.Printf("Public key parameter Q: %v\n", publickey.Parameters.Q)
	fmt.Printf("Public key parameter G: %v\n", publickey.Parameters.G)
	fmt.Printf("Public key Y: %v\n", publickey.Y)
}

func pubKeyRSA(publicKey string) {
	// Public Key RSA
	publickeyfile, err := os.Open(publicKey)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	decoder := gob.NewDecoder(publickeyfile)

	var publickey rsa.PublicKey
	err = decoder.Decode(&publickey)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	publickeyfile.Close()
	fmt.Printf("Public key: \n%x\n", publickey)
}
