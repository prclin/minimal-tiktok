package config

type Jwt struct {
	RSA *RSA
}

type RSA struct {
	//私钥
	PrivateKey string
	//公钥
	PublicKey string
}

var DefaultJwt = &Jwt{
	RSA: &RSA{},
}
