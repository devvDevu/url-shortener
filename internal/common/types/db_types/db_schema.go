package db_types

type DbSchema string

func (d DbSchema) String() string {
	return string(d)
}
