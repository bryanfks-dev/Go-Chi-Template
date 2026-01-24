package application

import (
	"skeleton/pkg/config"
	"skeleton/pkg/logger"
	"skeleton/pkg/timezone"
)

type Application struct {
	Environment *config.Environment
	Config      *config.Config
	Logger      *logger.Logger
}

func NewApplication() Application {
	environment := config.LoadEnvironment()
	config := config.NewConfig(environment)

	timezone.SetupTimezone(config.Timezone)
	logger := logger.NewLogger(environment, config.Logging)

	return Application{
		Environment: &environment,
		Config:      config,
		Logger:      &logger,
	}
}
