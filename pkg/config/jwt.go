package config

type JWTPayloadProperties struct {
	ExpirationMinutes int `yaml:"expiration_minutes"`
}

type JWTProperties struct {
	Secret    string               `yaml:"secret"`
	Algorithm string               `yaml:"algorithm"`
	Access    JWTPayloadProperties `yaml:"access"`
	Refresh   JWTPayloadProperties `yaml:"refresh"`
}
