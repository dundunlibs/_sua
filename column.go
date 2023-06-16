package sua

import (
	"fmt"
	"strings"
)

type Column struct {
	name         string
	datatype     string
	unique       bool
	notnull      bool
	defaultvalue string
	others       []string
}

func (c *Column) Name(v string) *Column {
	c.name = v
	return c
}

func (c *Column) Serial() *Column {
	return c.DataType("SERIAL")
}

func (c *Column) Varchar(limit int) *Column {
	return c.DataType(fmt.Sprintf("VARCHAR(%d)", limit))
}

func (c *Column) Text() *Column {
	return c.DataType("TEXT")
}

func (c *Column) DateTime() *Column {
	return c.Timestamp(3)
}

func (c *Column) Timestamp(p int) *Column {
	return c.DataType(fmt.Sprintf("TIMESTAMP(%d)", p))
}

func (c *Column) DataType(v string) *Column {
	c.datatype = v
	return c
}

func (c *Column) Unique() *Column {
	c.unique = true
	return c
}

func (c *Column) NotNull() *Column {
	c.notnull = true
	return c
}

func (c *Column) DefaultNow() *Column {
	return c.Default("CURRENT_TIMESTAMP")
}

func (c *Column) Default(v string) *Column {
	c.defaultvalue = v
	return c
}

func (c *Column) Others(args ...string) *Column {
	c.others = args
	return c
}

func (c *Column) string() string {
	strs := []string{c.name, c.datatype}
	if c.unique {
		strs = append(strs, "UNIQUE")
	}
	if c.notnull {
		strs = append(strs, "NOT NULL")
	}
	if c.defaultvalue != "" {
		strs = append(strs, fmt.Sprintf("DEFAULT %s", c.defaultvalue))
	}
	strs = append(strs, c.others...)
	return strings.Join(strs, " ")
}
