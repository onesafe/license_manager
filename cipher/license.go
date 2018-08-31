package cipher

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"encoding/pem"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"math/big"

	"github.com/onesafe/license_manager/views"
)

var (
	ErrInputSize  = errors.New("input size too large")
	ErrEncryption = errors.New("encryption error")
	publicKeyStr  = "LS0tLS1CRUdJTiBQVUJMSUMgS0VZLS0tLS0KTUlJQklqQU5CZ2txaGtpRzl3MEJBUUVGQUFPQ0FROEFNSUlCQ2dLQ0FRRUF6djYxZWFlS0piS1h5ZDFmcUs2WQpiRVNlaDZGYVVHZnhXejZ5eUROWVZZaElra3NXQmxhbDh6L2JiaXdrbGw5dGEramU4Rk5UbFRZZzVoWThCTUNhCmI3aHVnWTlTVFZpNFprcWR4WkF0Rnd6UDhTRmVvTmJpNjVGY09iSFFPdUJpRUdmNis3NC9Bd2cwNzFDVFFkelgKdmJSY1RvL2Yxd2tpVzdENytSSHFUOEFtWUEzbjRadkt4TDc3SDVlMExLMEh1cEFxNFJOd1cwVEpkakpCQWVQNQpMWWVLMXRsSEg0U2F5SHJzZ3V5NEFUMkcydWxYK0FseXBCMWl1Z2FaYzFoUThNYUFSQTZmQmZVV3BCNXFpMkRECnJ0dnRtMXJ6azhoWGlSK0FTQUZ1Vy9xTzkvVE5BK3hCaSsrTEpEZWxxZnU0dlNYU29pVjdsTUpZL1QrWXg1THYKcVFJREFRQUIKLS0tLS1FTkQgUFVCTElDIEtFWS0tLS0tCg=="
	privateKeyStr = "LS0tLS1CRUdJTiBSU0EgUFJJVkFURSBLRVktLS0tLQpNSUlFb2dJQkFBS0NBUUVBenY2MWVhZUtKYktYeWQxZnFLNlliRVNlaDZGYVVHZnhXejZ5eUROWVZZaElra3NXCkJsYWw4ei9iYml3a2xsOXRhK2plOEZOVGxUWWc1aFk4Qk1DYWI3aHVnWTlTVFZpNFprcWR4WkF0Rnd6UDhTRmUKb05iaTY1RmNPYkhRT3VCaUVHZjYrNzQvQXdnMDcxQ1RRZHpYdmJSY1RvL2Yxd2tpVzdENytSSHFUOEFtWUEzbgo0WnZLeEw3N0g1ZTBMSzBIdXBBcTRSTndXMFRKZGpKQkFlUDVMWWVLMXRsSEg0U2F5SHJzZ3V5NEFUMkcydWxYCitBbHlwQjFpdWdhWmMxaFE4TWFBUkE2ZkJmVVdwQjVxaTJERHJ0dnRtMXJ6azhoWGlSK0FTQUZ1Vy9xTzkvVE4KQSt4QmkrK0xKRGVscWZ1NHZTWFNvaVY3bE1KWS9UK1l4NUx2cVFJREFRQUJBb0lCQUE2dno5eUQ0Szl3RG9rSwpKZ0ZuOGZTc29icnAzZWlhbDJ0cjlXOXpCUzk4YmZuRFRockJCeUZOUkpwNmsxWUFDMWwrdW1LKzVzMU5lK1FYCmE4YkNpN0tsbVdHajg2ajlSSWd0ZnloSFpJdWdJeGp1enpJR1RTOTlydGVCSUU4dElJZ2RlSmRvRmp4MjFwcXMKVXVaM2J1OU0zc2ZNT1l1ZVlFdTFNUXVwRThENHRtdStqaGpBdytXcno1MDYrL1BUNkxMeVJVdmw1N2tsRVF1ZwoybHRmQ01jb1Q3WHoyd3RKT3NHdEw5emJmOXFJTlU5YStaWVNCaXdhUTlWakdlTU1wRHhPNFZ6WHpwUlZrMmtOCjU3QVV6dElGUWtkbkZ3WWI0cHUvdGdiSHRwbitEUDltc3BrUUp4QlhWaGJ2NmtwRTJJWkpyZXBUQ1Q2d1NvOUMKRDM5Mk95RUNnWUVBL25FejVSTVNJeEpiNE5reUFpLy9ZL0lYVUM2YjlycS9IRVp0d0o3TUgzZjdkQm12blYweQo5MTBGUW1XY1E1MU5YYVVQaGVCQkdtSTlHMmlGV3gzOWZnYllNdngwVnBNWlVZQ3pBek5OMFQ0b0pSTkNiVktwCndDN0VQdG9FZlJvb0VIdFNzWUhQSUZFK2FrL0NMZ1RRV2dwT0RpNHdWNmpvVDJyWmNxY2NUbVVDZ1lFQTBFTWoKNXVDb2N3cVpiTXhkZTNCb3J2ZUxjdmR3SG1aenQvQlo2eEdscVNRbE0xb3pNQWNDekl6VFNBR2FNQU1LUStGaAorTXNQSVdxenNFVWZFVXFJLzVDdXJZZG1rSE9DdFFjbzI1NUxtazhCU0ZSOWVEUm9OckxkaG5DbEhGTkxIaEFWCmdBeXlwaG9UK3loNGU1bFRTNVpzZ3ZCVG91UHpFcnhRMjBldE5mVUNnWUE1SlRXUmlsSDFmSWNVSGRQRWVBRTAKOGtkWUk4KzFmMFd5MVFLTFUxN201bXljSzdTc2RDVWhOMHdhR0hZYkhYWWx4UStTY3NaTVphbnh3T3pLaFJiTQp1ajdPWExMSVN1dFJ5Y1Rxd0JnSEdaMnNqZ0hLU2RtRUp6eStIUHNMR2RmTHM5Ymp0UkxPNEZCVFVpeDdaMnRmCi9aYUFTdEZpcnJYV09GbzBET0lublFLQmdCRU1mZ0xuZjBLenFtMnFxVGh5c2s4b2VxVDF0cHIyZmlDZnIyeTcKN0JqVm1hb0RoMDgyTTdkMUM5TElOc3daWTV5ODlaMDlXa2E5Q21xeXJlRm5mYUdXUVlaNUlCOVJKWEVXWGZUawpsNEhSVitTSTdpQ0tBY0lBa0h2eCtzSS8yMVZoc2JEaTJUa1p4MnIzSEMzYUZtU0lzdWRoTHllVmk0K01GUDV1CmRyS0ZBb0dBZTJIYjhYb2V3V3gzNEFCenFYejY3ZzVNd3FaemlOK2l0ajdxSDNUaXFYLysxVlpXYWdoV0RuencKbUtJTVB5VGttd3RQUzJHblZJUjg2VmMyWGM2c0xQMTVzalpuM1Q0Mm5XUTFaN3dwZkNGR1kwUkZsUWF3U0ExVgpWemx3NGtYYjJ4dTFHdW50NG1EeUhwNUtPZHZ4eUtqek1IYnNoNDRZMWw2SUZPMFoydmc9Ci0tLS0tRU5EIFJTQSBQUklWQVRFIEtFWS0tLS0tCg=="
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
	privateKeyBytes, err := base64.StdEncoding.DecodeString(privateKeyStr)
	if err != nil {
		err = fmt.Errorf("Decode private key error: %s", err)
		return
	}
	log.Println(string(privateKeyBytes))

	// take care private key block
	block, _ := pem.Decode(privateKeyBytes)

	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		err = fmt.Errorf("Parse private key interface error: %s", err)
		return
	}

	// rsa encrypt
	data, err := rsaEncrypt(privateKey, inputData)
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
	publicKeyBytes, err := base64.StdEncoding.DecodeString(publicKeyStr)
	if err != nil {
		err = fmt.Errorf("Decode public key error: %s", err)
		return
	}
	log.Println(string(publicKeyBytes))

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
