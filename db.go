package sua

import (
	"database/sql"

	"github.com/dundunlabs/sua/internal"
	"github.com/dundunlabs/sua/stmt"
)

func NewDB(sqldb *sql.DB) *DB {
	return &DB{
		DB: internal.NewDB(sqldb),
	}
}

type DB struct {
	*internal.DB
}

func (db *DB) CreateTable(name string, f func(stmt *stmt.CreateTableStmt)) *stmt.CreateTableStmt {
	stmt := stmt.NewCreateTableStmt(db.DB)
	stmt.Name(name)
	f(stmt)
	return stmt
}
