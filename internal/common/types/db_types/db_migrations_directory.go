package db_types

type DbMigrationsDirectory string

func (d DbMigrationsDirectory) String() string {
	return string(d)
}
