package config

type SQLDatabaseProperties struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Name     string `yaml:"name"`
}

type DatabaseProperties struct {
	Master SQLDatabaseProperties `yaml:"master"`
}
