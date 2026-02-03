package config

const AppEnvironmentKey = "APP_ENVIRONMENT"

var EnvironmentValue = map[string]Environment{
	"development": EnvironmentDevelopment,
	"production":  EnvironmentProduction,
}

var CfgFilePath = map[Environment]string{
	EnvironmentDevelopment: "files/config/development.yaml",
	EnvironmentProduction:  "files/config/production.yaml",
}
