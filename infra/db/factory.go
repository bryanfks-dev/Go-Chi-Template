package db

import (
	dbmaster "skeleton/infra/db/master"
	"skeleton/infra/ent"
	"skeleton/pkg/config"
	"skeleton/pkg/security"

	_ "skeleton/infra/ent/runtime"
)

type Database struct {
	MasterClient *ent.Client
}

func NewDatabase(
	cfg *config.DatabaseProperties,
	security *security.Security,
) *Database {
	masterDriver := GetDatabaseDriver(cfg.Master.Driver)
	masterClient := GetDatabaseClient(masterDriver, cfg.Master)
	dbmaster.RegisterHooks(masterClient, security)

	return &Database{
		MasterClient: masterClient,
	}
}
