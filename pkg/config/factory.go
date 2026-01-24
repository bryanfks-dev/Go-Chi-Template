package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Server   ServerProperties   `yaml:"server"`
	Logging  LoggingProperties  `yaml:"logging"`
	Database DatabaseProperties `yaml:"database"`
	Timezone TimezoneProperties `yaml:"timezone"`
}

var configFilePath = map[Environment]string{
	EnvironmentDevelopment: "development.yaml",
	EnvironmentProduction:  "production.yaml",
}

func getConfigFilePath(environment Environment) string {
	fallback := configFilePath[EnvironmentDevelopment]
	path, exists := configFilePath[environment]

	if exists {
		return path
	}
	return fallback
}

func NewConfig(environment Environment) *Config {
	config := &Config{}

	configFilePath := getConfigFilePath(environment)
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
