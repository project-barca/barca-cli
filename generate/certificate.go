package generate

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/gob"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"math/big"
	"os"
	"time"
)

func CertificateX509() {
	// see Certificate structure at
	// references: http://golang.org/pkg/crypto/x509/#Certificate
	template := &x509.Certificate{
		IsCA:                  true,
		BasicConstraintsValid: true,
		SubjectKeyId:          []byte{1, 2, 3},
		SerialNumber:          big.NewInt(1234),
		Subject: pkix.Name{
			Country:      []string{"Earth"},
			Organization: []string{"Mother Nature"},
		},
		NotBefore:   time.Now(),
		NotAfter:    time.Now().AddDate(5, 5, 5),
		ExtKeyUsage: []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth, x509.ExtKeyUsageServerAuth},
		KeyUsage:    x509.KeyUsageDigitalSignature | x509.KeyUsageCertSign,
	}
	// generate private key
	privatekey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		fmt.Println(err)
	}
	// define public key
	publickey := &privatekey.PublicKey
	// create a self-signed certificate. template = parent
	var parent = template

	cert, err := x509.CreateCertificate(rand.Reader, template, parent, publickey, privatekey)
	if err != nil {
		fmt.Println(err)
	}
	// save private key
	pkey := x509.MarshalPKCS1PrivateKey(privatekey)
	ioutil.WriteFile("private.key", pkey, 0777)
	fmt.Println("private key salvo:    ./RSAprivate.key")

	// save public key
	pubkey, _ := x509.MarshalPKIXPublicKey(publickey)
	ioutil.WriteFile("public.key", pubkey, 0777)
	fmt.Println("public key salvo:     ./RSApublic.key")

	// save cert
	ioutil.WriteFile("certificate.pem", cert, 0777)
	fmt.Println("certificate salvo:    ./certificate.pem")

	// these are the files save with encoding/gob style
	// save privategob key
	privkeyfile, _ := os.Create("privategob.key")
	privkeyencoder := gob.NewEncoder(privkeyfile)
	privkeyencoder.Encode(privatekey)
	privkeyfile.Close()
	// save publicgob key
	pubkeyfile, _ := os.Create("publicgob.key")
	pubkeyencoder := gob.NewEncoder(pubkeyfile)
	pubkeyencoder.Encode(publickey)
	pubkeyfile.Close()

	// Create plain text PEM file.
	pemfile, _ := os.Create("cert.pem")
	var pemkey = &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(privatekey)}
	pem.Encode(pemfile, pemkey)
	pemfile.Close()

}
