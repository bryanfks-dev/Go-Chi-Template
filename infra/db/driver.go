package db

type DatabaseDriver int

const (
	DatabaseDriverPostgreSQL = iota
	DatabaseDriverMySQL
)

func GetDatabaseDriver(driver string) DatabaseDriver {
	val, ok := DbDriver[driver]
	if !ok {
		panic("Unsupported database driver: " + driver)
	}
	return val
}
