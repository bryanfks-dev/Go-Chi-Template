package config

type BcryptProperties struct {
	Cost int `yaml:"cost"`
}

type HMACProperties struct {
	Algorithm   string `yaml:"algorithm"`
	TokenSecret string `yaml:"token_secret"`
}

type JWTPayloadProperties struct {
	ExpirationMinutes int `yaml:"expiration_minutes"`
}

type JWTProperties struct {
	Secret    string               `yaml:"secret"`
	Algorithm string               `yaml:"algorithm"`
	Access    JWTPayloadProperties `yaml:"access"`
	Refresh   JWTPayloadProperties `yaml:"refresh"`
}
