package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Application ApplicationProperties `yaml:"application"`
	Server      ServerProperties      `yaml:"server"`
	Logging     LoggingProperties     `yaml:"logging"`
	Database    DatabaseProperties    `yaml:"database"`
	Bcrypt      BcryptProperties      `yaml:"bcrypt"`
	HMAC        HMACProperties        `yaml:"hmac"`
	JWT         JWTProperties         `yaml:"jwt"`
	Timezone    TimezoneProperties    `yaml:"timezone"`
}

func NewConfig(env Environment) *Config {
	config := &Config{}

	configFilePath := getConfigFilePath(env)
	configFile, err := os.Open(configFilePath)
	if err != nil {
		panic("Failed to open config file: " + err.Error())
	}
	defer configFile.Close()

	configFileDecoder := yaml.NewDecoder(configFile)
	if err := configFileDecoder.Decode(&config); err != nil {
		panic("Failed to decode config file: " + err.Error())
	}

	return config
}

func getConfigFilePath(env Environment) string {
	val, ok := CfgFilePath[env]
	if !ok {
		return CfgFilePath[EnvironmentDevelopment]
	}
	return val
}
