package database

import (
	"context"
	"database/sql"
	"errors"
	"time"
	"url-shortener/internal/common/types/db_types"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/pressly/goose/v3"
	"github.com/sirupsen/logrus"
)

type PostgresAdapter struct {
	cfg    postgresConfigI
	connDb *sql.DB
	conn   *sqlx.DB
}

type postgresConfigI interface {
	GetHost() db_types.DbHost
	GetPort() db_types.DbPort
	GetAddr() db_types.DbAddr
	GetDatabase() db_types.DbName
	GetUserName() db_types.DbUserName
	GetPassword() db_types.DbPassword
	GetSchema() db_types.DbSchema
	GetUpMigrations() db_types.DbUpMigrations
	GetMaxIdleConnections() db_types.DbMaxIdleConnections
	GetMaxOpenConnections() db_types.DbMaxOpenConnections
	GetMigrationsDir() db_types.DbMigrationsDirectory
	GetConnectionMaxLifeTime() time.Duration
}

func (p *PostgresAdapter) Close() error {
	return p.conn.Close()
}

const adapterName = "PostgresAdapter "

func NewPostgresAdapter(ctx context.Context, cfg postgresConfigI) (*PostgresAdapter, error) {
	p := new(PostgresAdapter)
	p.cfg = cfg

	url := "postgres://" + cfg.GetUserName().String() + ":" +
		cfg.GetPassword().String() + "@" +
		cfg.GetAddr().String() + "/" +
		cfg.GetDatabase().String()

	conn, err := sqlx.Open("pgx", url)
	if err != nil {
		logrus.Fatalf("sqlx.Open(): %s", err.Error())
		return nil, err
	}

	err = conn.PingContext(ctx)
	if err != nil {
		logrus.Fatalf("conn.PingContext(ctx): %s", err.Error())
	}

	p.conn = conn

	p.conn.SetMaxIdleConns(p.cfg.GetMaxIdleConnections().Int())
	p.conn.SetMaxOpenConns(p.cfg.GetMaxOpenConnections().Int())
	p.conn.SetConnMaxLifetime(p.cfg.GetConnectionMaxLifeTime())

	if p.cfg.GetUpMigrations().Bool() {
		err := p.upMigrations(ctx, p.cfg.GetMigrationsDir())
		if err != nil {
			return nil, err
		}
	}

	return p, nil
}

func (p *PostgresAdapter) upMigrations(ctx context.Context, migrationsDir db_types.DbMigrationsDirectory) error {
	const dialect = "postgres"
	err := goose.SetDialect("pgx")
	if err != nil {
		return err
	}

	goose.SetLogger(logrus.WithFields(logrus.Fields{
		"dialect":        dialect,
		"event":          "goose_up",
		"migrations_dir": migrationsDir,
	}))

	err = goose.UpContext(ctx, p.getDbSqlConn(), migrationsDir.String())
	if err != nil {
		return err
	}

	return nil
}

func (p *PostgresAdapter) getDbSqlConn() *sql.DB {
	return p.conn.DB
}

// TODO Get, Select, NamedExecFetchRow

func (p *PostgresAdapter) Get(ctx context.Context, dest interface{}, query db_types.DbQuery, params ...interface{}) (ok bool, err error) {
	const action = "PostgresAdapter Get "
	const method = "Get"

	err = p.conn.GetContext(ctx, dest, query.String(), params...)
	if errors.Is(err, sql.ErrNoRows) {
		return ok, err
	}

	if err == nil {
		ok = true
	}

	if err != nil && ctx.Err() == nil {
		logrus.WithFields(logrus.Fields{
			"adapterName": adapterName,
			"method":      method,
			"query":       query.String(),
		}).WithError(err).Error(action)
	}

	return ok, err
}

func (p *PostgresAdapter) Select(ctx context.Context, dest interface{}, query db_types.DbQuery) (ok bool, err error) {
	const action = "PostgresAdapter Select "
	const method = "Select"

	err = p.conn.SelectContext(ctx, dest, query.String())
	if errors.Is(err, sql.ErrNoRows) {
		return ok, err
	}

	if err == nil {
		ok = true
	}

	if err != nil && ctx.Err() == nil {
		logrus.WithFields(logrus.Fields{
			"adapterName": adapterName,
			"method":      method,
			"query":       query.String(),
		}).WithError(err).Error(action)
	}

	return ok, err
}

func (p *PostgresAdapter) NamedExecFetchRow(ctx context.Context, dest interface{}, query db_types.DbQuery, arg interface{}) error {
	const action = "PostgresAdapter NamedExecFetchRow "
	const method = "NamedExecFetchRow"

	queryStr, params, err := p.conn.BindNamed(query.String()+" RETURNING *;", arg)
	if err != nil && ctx.Err() == nil {
		logrus.WithFields(logrus.Fields{
			"adapterName": adapterName,
			"method":      method,
			"query":       queryStr,
		}).WithError(err).Error(action)
	}

	err = p.conn.GetContext(ctx, dest, queryStr, params...)
	if err != nil && ctx.Err() == nil {
		logrus.WithFields(logrus.Fields{
			"adapterName": adapterName,
			"method":      method,
			"query":       queryStr,
		}).WithError(err).Error(action)
	}

	return err
}

func (p *PostgresAdapter) NamedExec(ctx context.Context, query db_types.DbQuery, arg interface{}) error {
	const action = "PostgresAdapter NamedExec "
	const method = "NamedExec"

	_, err := p.conn.NamedExecContext(ctx, query.String(), arg)

	if errors.Is(err, sql.ErrNoRows) {
		return nil
	}

	if err != nil && ctx.Err() == nil {
		logrus.WithFields(logrus.Fields{
			"adapterName": adapterName,
			"method":      method,
			"query":       query.String(),
		}).WithError(err).Error(action)
	}

	return err
}
