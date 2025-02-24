package database

import (
	"fmt"
	"time"
	"url-shortener/internal/common/types/db_types"
)

type PostgresConfig struct {
	Host                  db_types.DbHost                `yaml:"host" env:"POSTGRES_HOST" env-default:"127.0.0.1" env-required:"true" env-upd:""`
	Port                  db_types.DbPort                `yaml:"port" env:"POSTGRES_PORT" env-default:"5432" env-required:"true" env-upd:""`
	UserName              db_types.DbUserName            `yaml:"username" env:"POSTGRES_USERNAME" env-default:"default" env-required:"true" env-upd:""`
	Password              db_types.DbPassword            `yaml:"password" env:"POSTGRES_PASSWORD" env-default:"default" env-required:"true" env-upd:""`
	Database              db_types.DbName                `yaml:"database" env:"POSTGRES_DATABASE" env-default:"default" env-required:"true" env-upd:""`
	Schema                db_types.DbSchema              `yaml:"schema" env:"POSTGRES_SCHEMA" env-default:"public" env-required:"true" env-upd:""`
	MaxIdleConnections    db_types.DbMaxIdleConnections  `yaml:"max_idle_connections" env:"POSTGRES_MAX_IDLE_CONNECTIONS" env-default:"5" env-required:"true" env-upd:""`
	MaxOpenConnections    db_types.DbMaxOpenConnections  `yaml:"max_open_connections" env:"POSTGRES_MAX_OPEN_CONNECTIONS" env-default:"20" env-required:"true" env-upd:""`
	ConnectionMaxLifeTime time.Duration                  `yaml:"connection_max_lifetime" env:"POSTGRES_CONNECTION_MAX_LIFETIME" env-default:"1h" env-required:"true" env-upd:""`
	UpMigrations          db_types.DbUpMigrations        `yaml:"up_migrations" env:"POSTGRES_UP_MIGRATIONS" env-default:"true"`
	MigrationsDir         db_types.DbMigrationsDirectory `yaml:"migration_dir" env-required:"true"`
}

func (p *PostgresConfig) GetHost() db_types.DbHost {
	return p.Host
}

func (p *PostgresConfig) GetPort() db_types.DbPort {
	return p.Port
}

func (p *PostgresConfig) GetAddr() db_types.DbAddr {
	return db_types.DbAddr(fmt.Sprintf("%s:%s", p.GetHost(), p.GetPort()))
}

func (p *PostgresConfig) GetDatabase() db_types.DbName {
	return p.Database
}

func (p *PostgresConfig) GetUserName() db_types.DbUserName {
	return p.UserName
}

func (p *PostgresConfig) GetPassword() db_types.DbPassword {
	return p.Password
}

func (p *PostgresConfig) GetSchema() db_types.DbSchema {
	return p.Schema
}

func (p *PostgresConfig) GetUpMigrations() db_types.DbUpMigrations {
	return p.UpMigrations
}

func (p *PostgresConfig) GetMaxIdleConnections() db_types.DbMaxIdleConnections {
	return p.MaxIdleConnections
}

func (p *PostgresConfig) GetMaxOpenConnections() db_types.DbMaxOpenConnections {
	return p.MaxOpenConnections
}

func (p *PostgresConfig) GetMigrationsDir() db_types.DbMigrationsDirectory {
	return p.MigrationsDir
}

func (p *PostgresConfig) GetConnectionMaxLifeTime() time.Duration {
	return p.ConnectionMaxLifeTime
}
