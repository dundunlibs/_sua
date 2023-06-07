package internal

import "database/sql"

func NewDB(sqldb *sql.DB) *DB {
	if err := sqldb.Ping(); err != nil {
		panic(err)
	}
	return &DB{
		DB: sqldb,
	}
}

type DB struct {
	*sql.DB
}
