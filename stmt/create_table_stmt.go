package stmt

import (
	"fmt"
	"strings"

	"github.com/dundunlabs/sua/internal"
	"github.com/dundunlabs/xidau/slices"
)

func NewCreateTableStmt(db *internal.DB) *CreateTableStmt {
	stmt := &CreateTableStmt{}
	stmt.BaseStmt = &BaseStmt{
		db:   db,
		self: stmt,
	}

	return stmt
}

type CreateTableStmt struct {
	*BaseStmt
	name    string
	pk      []string
	columns []*Column
	ine     bool
}

func (stmt *CreateTableStmt) SQL() string {
	cmd := "create table"
	if stmt.ine {
		cmd += " if not exists"
	}

	if len(stmt.pk) == 0 {
		stmt.addDefaultPK()
	}

	cols := strings.Join(slices.Map(stmt.columns, func(c *Column) string {
		return c.string()
	}), ", ")

	pk := fmt.Sprintf("primary key (%s)", strings.Join(stmt.pk, ", "))

	return fmt.Sprintf("%s %s (%s, %s)", cmd, stmt.name, cols, pk)
}

func (stmt *CreateTableStmt) ID(col string) {
	stmt.Col(col).Serial()
	stmt.PK(col)
}

func (stmt *CreateTableStmt) Name(v string) {
	stmt.name = v
}

func (stmt *CreateTableStmt) Col(name string) *Column {
	col := &Column{
		name: name,
	}
	stmt.columns = append(stmt.columns, col)
	return col
}

func (stmt *CreateTableStmt) PK(cols ...string) {
	stmt.pk = cols
}

func (stmt *CreateTableStmt) IfNotExists() *CreateTableStmt {
	stmt.ine = true
	return stmt
}

func (stmt *CreateTableStmt) addDefaultPK() {
	stmt.columns = append([]*Column{
		{
			name:     "id",
			datatype: "serial",
			nullable: true,
		},
	}, stmt.columns...)
	stmt.PK("id")
}

type Column struct {
	name     string
	datatype string
	nullable bool
	args     []string
}

func (c *Column) Serial() *Column {
	return c.Datatype("serial")
}

func (c *Column) Varchar(length int) *Column {
	return c.Datatype(fmt.Sprintf("varchar(%d)", length))
}

func (c *Column) Timestamp(length int) *Column {
	return c.Datatype(fmt.Sprintf("timestamp(%d)", length))
}

func (c *Column) Nullable() *Column {
	c.nullable = true
	return c
}

func (c *Column) Unique() *Column {
	return c.Arg("unique")
}

func (c *Column) Default(v string) *Column {
	return c.Arg(fmt.Sprintf("default %s", v))
}

func (c *Column) Datatype(datatype string) *Column {
	c.datatype = datatype
	return c
}

func (c *Column) Arg(arg string) *Column {
	c.args = append(c.args, arg)
	return c
}

func (c *Column) string() string {
	strs := []string{c.name, c.datatype}
	if !c.nullable {
		strs = append(strs, "not null")
	}
	strs = append(strs, c.args...)
	return strings.Join(strs, " ")
}
