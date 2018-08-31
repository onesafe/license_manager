package modules

type RSAKeys struct {
	PrivateKey string `json:"privateKey"`
	PublicKey  string `json:"publicKey"`
}

type RSAKeysResp struct {
	Status     string  `json:"status"`
	BaseStatus string  `json:"baseStatus"`
	Message    string  `json:"msg"`
	Data       RSAKeys `json:"data"`
}
