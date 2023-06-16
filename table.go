package sua

import (
	"strings"

	"github.com/dundunlabs/xidau/slices"
)

type Table struct {
	name string
	cols Columns
	pk   PrimaryKey
}

func (t *Table) Name(v string) *Table {
	t.name = v
	return t
}

func (t *Table) PK(cols ...string) *Table {
	t.pk = cols
	return t
}

func (t *Table) ID(name string) *Column {
	return t.PK(name).Col(name).Serial()
}

func (t *Table) Col(name string) *Column {
	col := &Column{name: name}
	t.cols = append(t.cols, col)
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
