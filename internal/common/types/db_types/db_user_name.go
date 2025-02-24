package db_types

type DbUserName string

func (d DbUserName) String() string {
	return string(d)
}
