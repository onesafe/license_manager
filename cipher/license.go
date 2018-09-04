package cipher

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"math/big"

	"github.com/onesafe/license_manager/views"
)

var (
	ErrInputSize      = errors.New("input size too large")
	ErrEncryption     = errors.New("encryption error")
	JavaPublicKeyStr  = "MIICIjANBgkqhkiG9w0BAQEFAAOCAg8AMIICCgKCAgEAnbfH2UL9tW7oXY608VzlHqbX2MH+iSz3ROGX7nheegvqXY9YjwXMP/Mtxg3X0DCcO+MaEXeOhVpn7X+NbIEYHf7Cl8tekSFeF0obPKRl1SG3a+LBBGLCVv3JRhPUh2lyl7J8yG1R0thsbflos0wdMMdUiReN4xzLXm6v/sdh8NPEtYkqNoN++FpNxqavi9zaLQsHBNsVK7/uSU/wyo92rCYpBSgX57/TC3aHrT6GDxxy1z65UVJFjawRjEmeCAnsbCaam3wdWNDVHC2Q5GH8cmCAvIC1O1gbC5U6+fLU5ZovdCPdvSEG0dPIVEAq0I8QjNqSzpGjpJdb3LbVlMPCoAEbfoRkejDmV89DZH6R+AGCZcREUz3Gofaqgg7uBIyMsr/5AwZlu6aH3P8l5Ib25+2Hbwpul/2LT0NBTxD44VZdIqw//TMt2L+AeFWT3ED69Q6LZtL1/PUDH/J+sj2dyGfcfL6eYyx8lMrnfd6D6+ol0HnAhAOr5lqMYe5PpYGF1U1BAGbr38Vh1bLLzexa6qGNEb4rG4/JAEiA51u6bucklwuF3uOQWZPJLQ8UI5Tlt2li++jjncsawyRwTcMOPEuwLuGAli2a/C/NmAAqjdsFQldn1Pl0P5e7zb0V9TIGhNqOOQr6jp2X/6egRsEcjM9m3UT57Vpycr/PpYStAgkCAwEAAQ=="
	JavaPrivateKeyStr = "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"
)

// get key str from file
func ReadKeyStr(keyPath string) (KeyStr string) {
	data, err := ioutil.ReadFile(keyPath)
	if err != nil {
		panic("key file not exist")
	}
	KeyStr = string(data)
	return
}

// Encrypt License
func EncryptLicense(inputData []byte) (encodedData string, err error) {
	//privateKeyStr := ReadKeyStr("resources/private.key")
	privateKeyBytes, err := base64.StdEncoding.DecodeString(JavaPrivateKeyStr)
	if err != nil {
		err = fmt.Errorf("Decode private key error: %s", err)
		return
	}
	// log.Println(string(privateKeyBytes))

	// take care private key block
	// block, _ := pem.Decode(privateKeyBytes)

	privateKeyObj, err := x509.ParsePKCS8PrivateKey(privateKeyBytes)
	if err != nil {
		err = fmt.Errorf("Parse private key interface error: %s", err)
		return
	}

	// rsa encrypt
	data, err := rsaEncrypt(privateKeyObj.(*rsa.PrivateKey), inputData)
	if err != nil {
		err = fmt.Errorf("RSA Encrypt error")
	}

	// base64 encrypt
	encodeString := base64.StdEncoding.EncodeToString(data)
	return encodeString, nil
}

// Decrypt License Content to struct License
func DecryptLicense(content string) (license *views.License, err error) {

	// base64 decode
	licenseBytes, err := base64.StdEncoding.DecodeString(content)
	if err != nil {
		err = fmt.Errorf("Decode license content error: %s", err)
		return
	}

	//publicKeyStr := ReadKeyStr("resources/public.key")
	publicKeyBytes, err := base64.StdEncoding.DecodeString(JavaPublicKeyStr)
	if err != nil {
		err = fmt.Errorf("Decode public key error: %s", err)
		return
	}
	// log.Println(string(publicKeyBytes))

	// block, _ := pem.Decode(publicKeyBytes)

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

func rsaEncrypt(priv *rsa.PrivateKey, data []byte) (enc []byte, err error) {

	k := (priv.N.BitLen() + 7) / 8
	tLen := len(data)
	// rfc2313, section 8:
	// The length of the data D shall not be more than k-11 octets
	if tLen > k-11 {
		err = ErrInputSize
		return
	}
	em := make([]byte, k)
	em[1] = 1
	for i := 2; i < k-tLen-1; i++ {
		em[i] = 0xff
	}
	copy(em[k-tLen:k], data)
	c := new(big.Int).SetBytes(em)
	if c.Cmp(priv.N) > 0 {
		err = ErrEncryption
		return
	}
	var m *big.Int
	var ir *big.Int
	if priv.Precomputed.Dp == nil {
		m = new(big.Int).Exp(c, priv.D, priv.N)
	} else {
		// We have the precalculated values needed for the CRT.
		m = new(big.Int).Exp(c, priv.Precomputed.Dp, priv.Primes[0])
		m2 := new(big.Int).Exp(c, priv.Precomputed.Dq, priv.Primes[1])
		m.Sub(m, m2)
		if m.Sign() < 0 {
			m.Add(m, priv.Primes[0])
		}
		m.Mul(m, priv.Precomputed.Qinv)
		m.Mod(m, priv.Primes[0])
		m.Mul(m, priv.Primes[1])
		m.Add(m, m2)

		for i, values := range priv.Precomputed.CRTValues {
			prime := priv.Primes[2+i]
			m2.Exp(c, values.Exp, prime)
			m2.Sub(m2, m)
			m2.Mul(m2, values.Coeff)
			m2.Mod(m2, prime)
			if m2.Sign() < 0 {
				m2.Add(m2, prime)
			}
			m2.Mul(m2, values.R)
			m.Add(m, m2)
		}
	}

	if ir != nil {
		// Unblind.
		m.Mul(m, ir)
		m.Mod(m, priv.N)
	}
	enc = m.Bytes()
	return
}
