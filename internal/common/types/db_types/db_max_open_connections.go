package db_types

type DbMaxOpenConnections int

func (d DbMaxOpenConnections) Int() int {
	return int(d)
}
