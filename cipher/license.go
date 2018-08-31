package cipher

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/big"

	"github.com/onesafe/license_manager/views"
)

// get key str from file
func ReadKeyStr(keyPath string) (KeyStr string) {
	data, err := ioutil.ReadFile(keyPath)
	if err != nil {
		panic("public key file not exist")
	}
	KeyStr = string(data)
	return
}

// Decrypt License Content to struct License
func DecryptLicense(content string) (license *views.License, err error) {

	// base64 decode
	licenseBytes, err := base64.StdEncoding.DecodeString(content)
	if err != nil {
		err = fmt.Errorf("Decode license content error: %s", err)
		return
	}

	publicKeyStr := ReadKeyStr("resources/public.key")
	publicKeyBytes, err := base64.StdEncoding.DecodeString(publicKeyStr)
	if err != nil {
		err = fmt.Errorf("Decode public key error: %s", err)
		return
	}

	publicKeyObj, err := x509.ParsePKIXPublicKey(publicKeyBytes)
	if err != nil {
		err = fmt.Errorf("Parse public key interface error: %s", err)
		return
	}

	publicKey, ok := publicKeyObj.(*rsa.PublicKey)
	if !ok {
		err = fmt.Errorf("Asseration error: %s", err)
	}

	// use rsa public key decrypt content
	data := rsaDecrypt(publicKey, licenseBytes)
	fmt.Println("data is: " + string(data))

	// parse to License struct
	err = json.Unmarshal(data, &license)
	if err != nil {
		err = fmt.Errorf("Invalid License %s", err)
		return
	}
	return license, nil
}

func rsaDecrypt(publicKey *rsa.PublicKey, data []byte) []byte {
	c := new(big.Int)
	m := new(big.Int)
	m.SetBytes(data)
	e := big.NewInt(int64(publicKey.E))
	c.Exp(m, e, publicKey.N)
	out := c.Bytes()
	skip := 0
	for i := 2; i < len(out); i++ {
		if i+1 >= len(out) {
			break
		}
		if out[i] == 0xff && out[i+1] == 0 {
			skip = i + 2
			break
		}
	}
	return out[skip:]
}
