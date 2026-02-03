package config

import "os"

type Environment int

const (
	EnvironmentDevelopment = iota
	EnvironmentProduction
)

func LoadEnvironment() Environment {
	appEnvironment := os.Getenv(AppEnvironmentKey)
	val, ok := EnvironmentValue[appEnvironment]
	if !ok {
		return EnvironmentDevelopment
	}

	return val
}
