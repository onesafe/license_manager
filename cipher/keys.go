package cipher

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"io/ioutil"
	"log"
	"os"
)

func GenRsaKeys(bits int) (privateKeyStr, publicKeyStr string, err error) {
	// generate private key
	privateKey, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		return "", "", err
	}

	derStream := x509.MarshalPKCS1PrivateKey(privateKey)
	block := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: derStream,
	}

	// private key temp file
	privateFile, err := ioutil.TempFile("", "private.pem")
	if err != nil {
		panic(err)
	}
	log.Println(privateFile.Name())
	defer os.Remove(privateFile.Name())

	err = pem.Encode(privateFile, block)
	if err != nil {
		return "", "", err
	}

	// generate public key
	publicKey := &privateKey.PublicKey
	derPkix, err := x509.MarshalPKIXPublicKey(publicKey)
	if err != nil {
		return "", "", err
	}
	block = &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: derPkix,
	}

	// public key temp file
	publicFile, err := ioutil.TempFile("", "public.pem")
	if err != nil {
		panic(err)
	}
	log.Println(publicFile.Name())
	defer os.Remove(publicFile.Name())
	err = pem.Encode(publicFile, block)
	if err != nil {
		return "", "", err
	}

	// read string from key files
	privateKeyBytes, err := ioutil.ReadFile(privateFile.Name())
	if err != nil {
		return "", "", err
	}
	publicKeyBytes, err := ioutil.ReadFile(publicFile.Name())
	if err != nil {
		return "", "", err
	}
	return string(privateKeyBytes), string(publicKeyBytes), nil
}
