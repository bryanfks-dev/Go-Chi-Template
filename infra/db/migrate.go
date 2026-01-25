package db

import "context"

func (db *Database) Migrate() {
	if err := db.MasterClient.Schema.Create(context.Background()); err != nil {
		panic("Failed running database migration: " + err.Error())
	}
}
