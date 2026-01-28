package config

type JWTPayloadProperties struct {
	ExpirationMinutes int `yaml:"expiration_minutes"`
}

type JWTProperties struct {
	Secret  string               `yaml:"secret"`
	Access  JWTPayloadProperties `yaml:"access"`
	Refresh JWTPayloadProperties `yaml:"refresh"`
}
