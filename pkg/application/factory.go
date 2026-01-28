package application

import (
	"skeleton/infra/db"
	"skeleton/pkg/config"
	"skeleton/pkg/logger"
	"skeleton/pkg/timezone"
)

type Application struct {
	Environment *config.Environment
	Database    *db.Database
	Config      *config.Config
	Logger      *logger.Logger
}

func NewApplication() *Application {
	env := config.LoadEnvironment()
	cfg := config.NewConfig(env)
	db := db.NewDatabase(cfg.Database, cfg.Bcrypt)

	timezone.SetupTimezone(cfg.Timezone)
	logger := logger.NewLogger(env, cfg.Logging)

	return &Application{
		Environment: &env,
		Database:    db,
		Config:      cfg,
		Logger:      logger,
	}
}
