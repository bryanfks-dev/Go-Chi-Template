package db

import (
	"skeleton/infra/ent"
	"skeleton/pkg/config"
)

func GetDatabaseClient(driver DatabaseDriver, cfg any) *ent.Client {
	if cfg == nil {
		panic("Database configuration is nil")
	}

	switch driver {
	case DatabaseDriverPostgreSQL:
		return NewPostgreSQLDatabaseClient(cfg.(config.SQLDatabaseProperties))
	}

	panic("Unsupported database driver")
}
