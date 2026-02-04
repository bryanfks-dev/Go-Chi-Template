package db

import (
	dbmaster "skeleton/infra/db/master"
	"skeleton/infra/ent"
	"skeleton/pkg/config"
	"skeleton/pkg/logger"
	"skeleton/pkg/security"

	_ "skeleton/infra/ent/runtime"
)

type Database struct {
	env    config.Environment
	logger *logger.Logger

	MasterClient *ent.Client
}

func NewDatabase(
	cfg *config.DatabaseProperties,
	sec *security.Security,
	env config.Environment,
	logger *logger.Logger,
) *Database {
	masterDriver := GetDatabaseDriver(cfg.Master.Driver)
	masterClient := GetDatabaseClient(masterDriver, cfg.Master, env, logger)
	dbmaster.RegisterHooks(masterClient, sec)

	return &Database{
		MasterClient: masterClient,
	}
}
