package stmt

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestColumn(t *testing.T) {
	name := NewColumn().Name("name").Text().NotNull().Default("''")
	email := NewColumn().Name("bio").Varchar(255).Unique()
	assert.Equal(t, "name TEXT NOT NULL DEFAULT ''", name.string())
	assert.Equal(t, "bio VARCHAR(255) UNIQUE", email.string())
}
