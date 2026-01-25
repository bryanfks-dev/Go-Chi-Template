package db

import (
	"fmt"
	"skeleton/infra/ent"
	"skeleton/pkg/config"

	_ "github.com/lib/pq"
)

func NewPostgreSQLDatabaseClient(cfg config.SQLDatabaseProperties) *ent.Client {
	dsn := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.Host,
		cfg.Port,
		cfg.Username,
		cfg.Password,
		cfg.Name,
	)
	client, err := ent.Open(cfg.Driver, dsn)
	if err != nil {
		panic("Failed openning connection to MySQL server" + err.Error())
	}

	return client
}
