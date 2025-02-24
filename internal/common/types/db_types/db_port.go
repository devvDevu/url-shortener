package db_types

type DbPort string

func (d DbPort) String() string {
	return string(d)
}
