package db_types

type DbAddr string

func (d DbAddr) String() string {
	return string(d)
}
