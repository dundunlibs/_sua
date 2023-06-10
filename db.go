package sua

import "database/sql"

func NewDB(sqldb *sql.DB) *DB {
	return &DB{sqldb}
}

type DB struct {
	*sql.DB
}
