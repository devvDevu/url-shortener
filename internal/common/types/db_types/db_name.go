package db_types

type DbName string

func (d DbName) String() string {
	return string(d)
}
