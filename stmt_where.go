package sua

import (
	"fmt"
)

type Where []WhereGroup

func (w Where) Where(condition string) Where {
	return w.where(WhereGroup{
		condition: condition,
	})
}

func (w Where) WhereNot(condition string) Where {
	return w.where(WhereGroup{
		condition: condition,
		not:       true,
	})
}

func (w Where) And(condition string) Where {
	return w.where(WhereGroup{
		condition: condition,
		operator:  FilterOperatorAND,
	})
}

func (w Where) Or(condition string) Where {
	return w.where(WhereGroup{
		condition: condition,
		operator:  FilterOperatorOR,
	})
}

func (w Where) AndNot(condition string) Where {
	return w.where(WhereGroup{
		condition: condition,
		operator:  FilterOperatorAND,
		not:       true,
	})
}

func (w Where) OrNot(condition string) Where {
	return w.where(WhereGroup{
		condition: condition,
		operator:  FilterOperatorOR,
		not:       true,
	})
}

func (w Where) WhereGroup(fn func(g Where) Where) Where {
	return w.filterGroup(fn, FilterOperatorAND, false)
}

func (w Where) WhereNotGroup(fn func(g Where) Where) Where {
	return w.filterGroup(fn, FilterOperatorAND, true)
}

func (w Where) AndGroup(fn func(g Where) Where) Where {
	return w.filterGroup(fn, FilterOperatorAND, false)
}

func (w Where) OrGroup(fn func(g Where) Where) Where {
	return w.filterGroup(fn, FilterOperatorOR, false)
}

func (w Where) AndNotGroup(fn func(g Where) Where) Where {
	return w.filterGroup(fn, FilterOperatorAND, true)
}

func (w Where) OrNotGroup(fn func(g Where) Where) Where {
	return w.filterGroup(fn, FilterOperatorOR, true)
}

func (w Where) filterGroup(
	fn func(w Where) Where,
	operator FilterOperator,
	not bool,
) Where {
	g := fn(Where{})
	return w.where(WhereGroup{
		groups:   g,
		operator: operator,
		not:      not,
	})
}

func (w Where) where(f WhereGroup) Where {
	nw := w
	nw = append(nw, f)
	return nw
}

func (w Where) string() string {
	r := ""
	for i, g := range w {
		if i > 0 {
			r += fmt.Sprintf(" %s ", g.operator)
		}
		if g.not {
			r += "NOT "
		}
		if len(g.groups) > 0 {
			r += g.groups.string()
		} else {
			r += g.condition
		}
	}
	return fmt.Sprintf("(%s)", r)
}

type FilterOperator string

const (
	FilterOperatorAND FilterOperator = "AND"
	FilterOperatorOR  FilterOperator = "OR"
)

type WhereGroup struct {
	groups    Where
	condition string
	not       bool
	operator  FilterOperator
}
