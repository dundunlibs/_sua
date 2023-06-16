package sua

import (
	"context"
	"database/sql"
	"fmt"
)

func (db *DB) CreateTable(name string, fn func(t *Table)) *CreateTable {
	t := &Table{name: name}
	fn(t)
	return &CreateTable{
		db:    db,
		table: t,
	}
}

type CreateTable struct {
	db    *DB
	table *Table
	ine   bool
}

func (ct *CreateTable) IfNotExists() *CreateTable {
	ct.ine = true
	return ct
}

func (ct *CreateTable) SQL() string {
	cmd := "CREATE TABLE"
	if ct.ine {
		cmd += " IF NOT EXISTS"
	}
	return fmt.Sprintf(
		"%s %q (%s, PRIMARY KEY(%s));",
		cmd,
		ct.table.name,
		ct.table.cols.string(),
		ct.table.pk.string(),
	)
}

func (ct *CreateTable) Exec(ctx context.Context) (sql.Result, error) {
	query := ct.SQL()
	return ct.db.ExecContext(ctx, query)
}
