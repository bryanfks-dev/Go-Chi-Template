package db

import (
	"skeleton/infra/ent"
	"skeleton/pkg/config"
)

type Database struct {
	MasterClient *ent.Client
}

func NewDatabase(cfg config.DatabaseProperties) *Database {
	masterDriver := GetDatabaseDriver(cfg.Master.Driver)
	masterClient := GetDatabaseClient(masterDriver, cfg.Master)

	return &Database{
		MasterClient: masterClient,
	}
}
