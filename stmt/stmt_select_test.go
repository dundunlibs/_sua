package stmt

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var sq = StmtSelect{
	table: "users",
}

func TestSelectAll(t *testing.T) {
	assert.Equal(t, "SELECT * FROM users;", sq.SQL())
}

func TestSelectFields(t *testing.T) {
	q := sq.Select("id", "name")
	assert.Equal(t, "SELECT id, name FROM users;", q.SQL())
}

func TestSelectLimitOffset(t *testing.T) {
	q := sq.Limit(10).Offset(100)
	assert.Equal(t, "SELECT * FROM users LIMIT 10 OFFSET 100;", q.SQL())
}

func TestSelectOrder(t *testing.T) {
	q1 := sq.Order("id", SortASC)
	q2 := q1.Order("name", SortDESC)
	assert.Equal(t, "SELECT * FROM users ORDER BY id ASC;", q1.SQL())
	assert.Equal(t, "SELECT * FROM users ORDER BY id ASC, name DESC;", q2.SQL())
}

func TestSelectOrderReversed(t *testing.T) {
	q := sq.Order("id", SortASC).Order("name", SortDESC).ReverseOrder()
	assert.Equal(t, "SELECT * FROM users ORDER BY id DESC, name ASC;", q.SQL())
}

func TestSelectFilters(t *testing.T) {
	q1 := sq.Where("id = 1")
	q2 := q1.And("name = 'foo'")
	q3 := q2.Or("name = 'bar'")
	q4 := sq.WhereNot("id IS NULL").AndNot("name = ''").OrNot("name IS NULL")
	assert.Equal(t, "SELECT * FROM users WHERE (id = 1);", q1.SQL())
	assert.Equal(t, "SELECT * FROM users WHERE (id = 1 AND name = 'foo');", q2.SQL())
	assert.Equal(t, "SELECT * FROM users WHERE (id = 1 AND name = 'foo' OR name = 'bar');", q3.SQL())
	assert.Equal(t, "SELECT * FROM users WHERE (NOT id IS NULL AND NOT name = '' OR NOT name IS NULL);", q4.SQL())
}

func TestSelectFiltersGroup(t *testing.T) {
	q1 := sq.WhereGroup(func(g Filters) Filters {
		return g.Where("id = 1").And("name = 'foo'")
	}).OrGroup(func(g Filters) Filters {
		return g.Where("id = 2").And("name = 'bar'")
	}).AndGroup(func(g Filters) Filters {
		return g.WhereNot("name = ''").AndNot("name IS NULL")
	})
	q2 := sq.WhereNotGroup(func(g Filters) Filters {
		return g.Where("1 = 1")
	}).AndNotGroup(func(g Filters) Filters {
		return g.Where("TRUE")
	}).OrNotGroup(func(g Filters) Filters {
		return g.WhereNot("FALSE")
	})
	assert.Equal(t, "SELECT * FROM users WHERE ((id = 1 AND name = 'foo') OR (id = 2 AND name = 'bar') AND (NOT name = '' AND NOT name IS NULL));", q1.SQL())
	assert.Equal(t, "SELECT * FROM users WHERE (NOT (1 = 1) AND NOT (TRUE) OR NOT (NOT FALSE));", q2.SQL())
}
