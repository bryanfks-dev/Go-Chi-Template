package db

import (
	dbmaster "skeleton/infra/db/master"
	"skeleton/infra/ent"
	"skeleton/pkg/config"

	_ "skeleton/infra/ent/runtime"
)

type Database struct {
	MasterClient *ent.Client
}

func NewDatabase(
	cfg config.DatabaseProperties,
	bcryptCfg config.BcryptProperties,
) *Database {
	masterDriver := GetDatabaseDriver(cfg.Master.Driver)
	masterClient := GetDatabaseClient(masterDriver, cfg.Master)
	dbmaster.RegisterHooks(masterClient, bcryptCfg)

	return &Database{
		MasterClient: masterClient,
	}
}
