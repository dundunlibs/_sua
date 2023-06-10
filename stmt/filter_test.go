package stmt

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFilters(t *testing.T) {
	f := Filters{}.Where("id = 1").Or("id = 2").And("id IS NOT NULL").AndNot("name IS NULL").OrNot("name = ''")
	assert.Equal(t, "(id = 1 OR id = 2 AND id IS NOT NULL AND NOT name IS NULL OR NOT name = '')", f.string())
}

func TestFiltersGroup(t *testing.T) {
	f := Filters{}.WhereGroup(func(g Filters) Filters {
		return g.WhereNotGroup(func(g Filters) Filters {
			return g.Where("id = 1").Or("id = 2")
		})
	}).AndGroup(func(g Filters) Filters {
		return g.Where("name IS NOT NULL").AndNotGroup(func(g Filters) Filters {
			return g.WhereNot("name = ''").AndNot("name = ' ")
		})
	}).OrGroup(func(g Filters) Filters {
		return g.WhereNot("email LIKE '%gmail.com'").OrNotGroup(func(g Filters) Filters {
			return g.Where("email IS NULL")
		})
	})
	assert.Equal(t, "((NOT (id = 1 OR id = 2)) AND (name IS NOT NULL AND NOT (NOT name = '' AND NOT name = ' )) OR (NOT email LIKE '%gmail.com' OR NOT (email IS NULL)))", f.string())
}
