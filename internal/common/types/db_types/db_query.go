package db_types

type DbQuery string

func (d DbQuery) String() string {
	return string(d)
}
