package db_types

type DbPassword string

func (d DbPassword) String() string {
	return string(d)
}
