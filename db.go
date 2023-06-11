package sua

import (
	"database/sql"

	"github.com/dundunlabs/sua/core"
	"github.com/dundunlabs/sua/stmt"
)

func NewDB(sqldb *sql.DB) *DB {
	return &DB{core.NewDB(sqldb)}
}

type DB struct {
	*core.DB
}

func (db *DB) CreateTable(name string, fn func(t *stmt.CreateTable)) *stmt.StmtCreateTable {
	t := &stmt.CreateTable{}
	t.Name(name)
	fn(t)
	return stmt.NewStmtCreateTable(t, db.DB)
}
