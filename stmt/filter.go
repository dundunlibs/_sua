package stmt

import (
	"fmt"
)

type Filters []filter

func (fs Filters) Where(condition string) Filters {
	return fs.filter(filter{
		condition: condition,
	})
}

func (fs Filters) WhereNot(condition string) Filters {
	return fs.filter(filter{
		condition: condition,
		not:       true,
	})
}

func (fs Filters) And(condition string) Filters {
	return fs.filter(filter{
		condition: condition,
		operator:  FilterOperatorAND,
	})
}

func (fs Filters) Or(condition string) Filters {
	return fs.filter(filter{
		condition: condition,
		operator:  FilterOperatorOR,
	})
}

func (fs Filters) AndNot(condition string) Filters {
	return fs.filter(filter{
		condition: condition,
		operator:  FilterOperatorAND,
		not:       true,
	})
}

func (fs Filters) OrNot(condition string) Filters {
	return fs.filter(filter{
		condition: condition,
		operator:  FilterOperatorOR,
		not:       true,
	})
}

func (fs Filters) WhereGroup(fn func(g Filters) Filters) Filters {
	return fs.filterGroup(fn, FilterOperatorAND, false)
}

func (fs Filters) WhereNotGroup(fn func(g Filters) Filters) Filters {
	return fs.filterGroup(fn, FilterOperatorAND, true)
}

func (fs Filters) AndGroup(fn func(g Filters) Filters) Filters {
	return fs.filterGroup(fn, FilterOperatorAND, false)
}

func (fs Filters) OrGroup(fn func(g Filters) Filters) Filters {
	return fs.filterGroup(fn, FilterOperatorOR, false)
}

func (fs Filters) AndNotGroup(fn func(g Filters) Filters) Filters {
	return fs.filterGroup(fn, FilterOperatorAND, true)
}

func (fs Filters) OrNotGroup(fn func(g Filters) Filters) Filters {
	return fs.filterGroup(fn, FilterOperatorOR, true)
}

func (fs Filters) filterGroup(
	fn func(fs Filters) Filters,
	operator FilterOperator,
	not bool,
) Filters {
	g := fn(Filters{})
	return fs.filter(filter{
		group:    g,
		operator: operator,
		not:      not,
	})
}

func (fs Filters) filter(f filter) Filters {
	newFs := fs
	newFs = append(newFs, f)
	return newFs
}

func (fs Filters) string() string {
	r := ""
	for i, f := range fs {
		if i > 0 {
			r += fmt.Sprintf(" %s ", f.operator)
		}
		if f.not {
			r += "NOT "
		}
		if len(f.group) > 0 {
			r += f.group.string()
		} else {
			r += f.condition
		}
	}
	return fmt.Sprintf("(%s)", r)
}

type FilterOperator string

const (
	FilterOperatorAND FilterOperator = "AND"
	FilterOperatorOR  FilterOperator = "OR"
)

type filter struct {
	group     Filters
	condition string
	not       bool
	operator  FilterOperator
}
