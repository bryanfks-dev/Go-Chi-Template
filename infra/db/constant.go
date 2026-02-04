package db

var DbDriver = map[string]DatabaseDriver{
	"postgres": DatabaseDriverPostgreSQL,
	"mysql":    DatabaseDriverMySQL,
}
