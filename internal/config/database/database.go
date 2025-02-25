package database

type DatabaseConfig struct {
	Postgres PostgresConfig `yaml:"postgres"`
}

func (d *DatabaseConfig) GetPostgres() *PostgresConfig {
	return &d.Postgres
}
