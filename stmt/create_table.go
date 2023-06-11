package stmt

import (
	"fmt"
	"strings"

	"github.com/dundunlabs/sua/core"
	"github.com/dundunlabs/xidau/slices"
)

func NewStmtCreateTable(t *CreateTable, db *core.DB) *StmtCreateTable {
	stmt := &StmtCreateTable{t: t}
	stmt.BaseStmt = &BaseStmt{
		db:   db,
		self: stmt,
	}
	return stmt
}

type StmtCreateTable struct {
	*BaseStmt
	t   *CreateTable
	ine bool
}

func (stmt *StmtCreateTable) IfNotExists() *StmtCreateTable {
	stmt.ine = true
	return stmt
}

func (stmt *StmtCreateTable) SQL() string {
	cmd := "CREATE TABLE"
	if stmt.ine {
		cmd += " IF NOT EXISTS"
	}
	return fmt.Sprintf(
		"%s %q (%s, PRIMARY KEY(%s));",
		cmd,
		stmt.t.name,
		stmt.t.columns.string(),
		stmt.t.pk.string(),
	)
}

func NewCreateTable(name string) *CreateTable {
	return &CreateTable{name: name}
}

type CreateTable struct {
	name    string
	columns Columns
	pk      PrimaryKey
}

func (ct *CreateTable) Name(v string) *CreateTable {
	ct.name = v
	return ct
}

func (ct *CreateTable) PK(cols ...string) *CreateTable {
	ct.pk = cols
	return ct
}

func (ct *CreateTable) ID(name string) *Column {
	return ct.PK(name).Col(name).Serial()
}

func (ct *CreateTable) Col(name string) *Column {
	col := &Column{name: name}
	ct.columns = append(ct.columns, col)
	return col
}

type Columns []*Column

func (cs Columns) string() string {
	cols := slices.Map(cs, func(c *Column) string {
		return c.string()
	})
	return strings.Join(cols, ", ")
}

type PrimaryKey []string

func (pk PrimaryKey) string() string {
	return strings.Join(pk, ", ")
}
