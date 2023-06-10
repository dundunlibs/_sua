package stmt

import (
	"fmt"
	"strings"

	"github.com/dundunlabs/xidau/slices"
)

type StmtSelect struct {
	table   string
	fields  fields
	limit   *int
	offset  *int
	orders  orders
	filters Filters
}

func (stmt StmtSelect) Select(fields ...string) StmtSelect {
	newStmt := stmt
	newStmt.fields = fields
	return newStmt
}

func (stmt StmtSelect) Limit(limit int) StmtSelect {
	newStmt := stmt
	newStmt.limit = &limit
	return newStmt
}

func (stmt StmtSelect) Offset(offset int) StmtSelect {
	newStmt := stmt
	newStmt.offset = &offset
	return newStmt
}

func (stmt StmtSelect) Order(field string, sort Sort) StmtSelect {
	newStmt := stmt
	newStmt.orders = append(newStmt.orders, order{field, sort})
	return newStmt
}

func (stmt StmtSelect) ReverseOrder() StmtSelect {
	newStmt := stmt
	newStmt.orders = slices.Map(newStmt.orders, func(o order) order {
		var sort Sort
		switch o.sort {
		case SortASC:
			sort = SortDESC
		case SortDESC:
			sort = SortASC
		}
		return order{o.field, sort}
	})
	return newStmt
}

func (stmt StmtSelect) Where(condition string) StmtSelect {
	return stmt.filter(filter{
		condition: condition,
	})
}

func (stmt StmtSelect) WhereNot(condition string) StmtSelect {
	return stmt.filter(filter{
		condition: condition,
		not:       true,
	})
}

func (stmt StmtSelect) And(condition string) StmtSelect {
	return stmt.filter(filter{
		condition: condition,
		operator:  FilterOperatorAND,
	})
}

func (stmt StmtSelect) Or(condition string) StmtSelect {
	return stmt.filter(filter{
		condition: condition,
		operator:  FilterOperatorOR,
	})
}

func (stmt StmtSelect) AndNot(condition string) StmtSelect {
	return stmt.filter(filter{
		condition: condition,
		not:       true,
		operator:  FilterOperatorAND,
	})
}

func (stmt StmtSelect) OrNot(condition string) StmtSelect {
	return stmt.filter(filter{
		condition: condition,
		not:       true,
		operator:  FilterOperatorOR,
	})
}

func (stmt StmtSelect) WhereGroup(fn func(g Filters) Filters) StmtSelect {
	return stmt.filterGroup(fn, FilterOperatorAND, false)
}

func (stmt StmtSelect) WhereNotGroup(fn func(g Filters) Filters) StmtSelect {
	return stmt.filterGroup(fn, FilterOperatorAND, true)
}

func (stmt StmtSelect) AndGroup(fn func(g Filters) Filters) StmtSelect {
	return stmt.filterGroup(fn, FilterOperatorAND, false)
}

func (stmt StmtSelect) OrGroup(fn func(g Filters) Filters) StmtSelect {
	return stmt.filterGroup(fn, FilterOperatorOR, false)
}

func (stmt StmtSelect) AndNotGroup(fn func(g Filters) Filters) StmtSelect {
	return stmt.filterGroup(fn, FilterOperatorAND, true)
}

func (stmt StmtSelect) OrNotGroup(fn func(g Filters) Filters) StmtSelect {
	return stmt.filterGroup(fn, FilterOperatorOR, true)
}

func (stmt StmtSelect) SQL() string {
	query := fmt.Sprintf(
		"SELECT %s FROM %s",
		stmt.fields.string(),
		stmt.table,
	)
	if len(stmt.filters) > 0 {
		query += fmt.Sprintf(" WHERE %s", stmt.filters.string())
	}
	if stmt.limit != nil {
		query += fmt.Sprintf(" LIMIT %d", *stmt.limit)
	}
	if stmt.offset != nil {
		query += fmt.Sprintf(" OFFSET %d", *stmt.offset)
	}
	if len(stmt.orders) > 0 {
		query += fmt.Sprintf(" ORDER BY %s", stmt.orders.string())
	}
	return fmt.Sprintf("%s;", query)
}

func (stmt StmtSelect) filterGroup(
	fn func(fs Filters) Filters,
	operator FilterOperator,
	not bool,
) StmtSelect {
	g := fn(Filters{})
	return stmt.filter(filter{
		group:    g,
		not:      not,
		operator: operator,
	})
}

func (stmt StmtSelect) filter(f filter) StmtSelect {
	newStmt := stmt
	newStmt.filters = append(newStmt.filters, f)
	return newStmt
}

type fields []string

func (fs fields) string() string {
	if len(fs) == 0 {
		return "*"
	}
	return strings.Join(fs, ", ")
}

type orders []order

func (os orders) string() string {
	r := slices.Map(os, func(o order) string {
		return fmt.Sprintf("%s %s", o.field, o.sort)
	})
	return strings.Join(r, ", ")
}

type Sort string

const (
	SortASC  Sort = "ASC"
	SortDESC Sort = "DESC"
)

type order struct {
	field string
	sort  Sort
}
