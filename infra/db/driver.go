package db

type DatabaseDriver int

const (
	DatabaseDriverPostgreSQL = iota
	DatabaseDriverMySQL
)

var dbDriver = map[string]DatabaseDriver{
	"postgres": DatabaseDriverPostgreSQL,
	"mysql":    DatabaseDriverMySQL,
}

func GetDatabaseDriver(driver string) DatabaseDriver {
	val, ok := dbDriver[driver]
	if !ok {
		panic("Unsupported database driver: " + driver)
	}

	return val
}
