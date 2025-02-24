package db_types

type DbUpMigrations bool

func (d DbUpMigrations) Bool() bool {
	return bool(d)
}
