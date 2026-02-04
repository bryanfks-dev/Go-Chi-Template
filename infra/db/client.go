package db

import (
	"context"
	"database/sql"
	"fmt"
	"skeleton/infra/ent"
	"skeleton/pkg/config"
	"skeleton/pkg/logger"
	"time"

	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	_ "github.com/jackc/pgx/v5/stdlib"
)

func GetDatabaseClient(
	driver DatabaseDriver,
	cfg any,
	env config.Environment,
	logger *logger.Logger,
) *ent.Client {
	if cfg == nil {
		panic("db config cannot be nil")
	}

	switch driver {
	case DatabaseDriverPostgreSQL:
		cfg, ok := cfg.(config.SQLDatabaseProperties)
		if !ok {
			panic(
				"postgresql config should have been follows SQLDatabaseProperties struct",
			)
		}
		db := newPostgreSQLDatabaseClient(&cfg)
		client := newEntClient(
			dialect.Postgres,
			db,
			env,
			logger,
		)
		return client
	}

	panic("Unsupported database driver")
}

func newPostgreSQLDatabaseClient(cfg *config.SQLDatabaseProperties) *sql.DB {
	dsn := fmt.Sprintf(
		"postgresql://%s:%s@%s:%d/%s?sslmode=%s",
		cfg.Username,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.Name,
		cfg.SSLMode,
	)
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		panic("Failed openning connection to SQL server" + err.Error())
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := db.PingContext(ctx); err != nil {
		panic("Failed pinging SQL server" + err.Error())
	}

	return db
}

func newEntClient(
	dialect string,
	db *sql.DB,
	env config.Environment,
	logger *logger.Logger,
) *ent.Client {
	drv := entsql.OpenDB(dialect, db)
	client := ent.NewClient(ent.Driver(drv), ent.Log(func(arg ...any) {
		logger.Info(fmt.Sprint(arg...))
	}))
	if env == config.EnvironmentDevelopment {
		client = client.Debug()
	}
	return client
}
