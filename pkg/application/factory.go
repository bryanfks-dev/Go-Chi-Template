package application

import (
	"skeleton/infra/db"
	"skeleton/pkg/config"
	"skeleton/pkg/logger"
	"skeleton/pkg/security"
	"skeleton/pkg/timezone"
)

type Application struct {
	Environment config.Environment
	Database    *db.Database
	Config      *config.Config
	Logger      *logger.Logger
	Security    *security.Security
}

func NewApplication() *Application {
	env := config.LoadEnvironment()
	cfg := config.NewConfig(env)

	timezone.SetupTimezone(&cfg.Timezone)
	logger := logger.NewLogger(env, &cfg.Logging)
	security := security.NewSecurity(&cfg.Bcrypt)

	db := db.NewDatabase(&cfg.Database, security)

	return &Application{
		Environment: env,
		Database:    db,
		Config:      cfg,
		Logger:      logger,
		Security:    security,
	}
}
