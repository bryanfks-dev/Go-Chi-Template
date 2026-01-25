package config

import "os"

type Environment int

const (
	EnvironmentDevelopment = iota
	EnvironmentProduction
)

var environment = map[string]Environment{
	"development": EnvironmentDevelopment,
	"production":  EnvironmentProduction,
}

func LoadEnvironment() Environment {
	appEnvironment := os.Getenv("APP_ENVIRONMENT")
	val, ok := environment[appEnvironment]
	if !ok {
		return EnvironmentDevelopment
	}

	return val
}
