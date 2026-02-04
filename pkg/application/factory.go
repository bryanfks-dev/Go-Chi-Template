package application

import (
	"skeleton/infra/db"
	"skeleton/pkg/config"
	"skeleton/pkg/logger"
	"skeleton/pkg/security"
	"skeleton/pkg/timezone"
)

type Application struct {
	Env    config.Environment
	Db     *db.Database
	Cfg    *config.Config
	Logger *logger.Logger
	Sec    *security.Security
}

func NewApplication() *Application {
	env := config.LoadEnvironment()
	cfg := config.NewConfig(env)

	timezone.SetupTimezone(&cfg.Timezone)
	logger := logger.NewLogger(env, &cfg.Logging)
	security := security.NewSecurity(
		&cfg.Application,
		&cfg.Bcrypt,
		&cfg.HMAC,
		&cfg.JWT,
	)

	db := db.NewDatabase(&cfg.Database, security, env, logger)

	return &Application{
		Env:    env,
		Db:     db,
		Cfg:    cfg,
		Logger: logger,
		Sec:    security,
	}
}
