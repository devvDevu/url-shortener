package db_types

type DbHost string

func (d DbHost) String() string {
	return string(d)
}
