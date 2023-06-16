package sua

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestColumn(t *testing.T) {
	name := (&Column{}).Name("name").Text().NotNull().Default("''")
	email := (&Column{}).Name("bio").Varchar(255).Unique()
	assert.Equal(t, "name TEXT NOT NULL DEFAULT ''", name.string())
	assert.Equal(t, "bio VARCHAR(255) UNIQUE", email.string())
}
