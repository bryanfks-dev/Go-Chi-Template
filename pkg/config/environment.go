package config

import "os"

type Environment string

const (
	EnvironmentDevelopment Environment = "development"
	EnvironmentProduction  Environment = "production"
)

func LoadEnvironment() Environment {
	fallback := EnvironmentDevelopment
	environment := os.Getenv("APP_ENVIRONMENT")

	if environment == "production" {
		return EnvironmentProduction
	}
	return fallback
}
