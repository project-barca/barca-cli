package scan

import (
	"bufio"
	"crypto/dsa"
	"crypto/rsa"
	"encoding/gob"
	"encoding/pem"
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

func privKeyRSA(privateKey string) {
	// Private Key RSA
	privatekeyfile, err := os.Open(privateKey)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	decoder := gob.NewDecoder(privatekeyfile)

	var privatekey rsa.PrivateKey
	err = decoder.Decode(&privatekey)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	privatekeyfile.Close()
	fmt.Printf("Private Key: \n%x\n", privatekey)
}

func PEMFile(pemFile string) {
	pemfile, err := os.Open(pemFile)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	// convert pemfile to []byte for decoding
	pemfileinfo, _ := pemfile.Stat()
	var size int64 = pemfileinfo.Size()
	pembytes := make([]byte, size)

	// read pemfile content into pembytes
	buffer := bufio.NewReader(pemfile)
	_, err = buffer.Read(pembytes)

	// proper decoding now
	data, _ := pem.Decode([]byte(pembytes))

	pemfile.Close()
	fmt.Printf("PEM Type: \n%s\n", data.Type)
	fmt.Printf("PEM Headers: \n%s\n", data.Headers)
	fmt.Printf("PEM Bytes: \n%x\n", string(data.Bytes))
}
