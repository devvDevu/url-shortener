package db_types

type DbMaxIdleConnections int

func (d DbMaxIdleConnections) Int() int {
	return int(d)
}
