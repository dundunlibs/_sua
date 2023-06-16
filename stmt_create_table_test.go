package sua

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateTable(t *testing.T) {
	table := &Table{}
	table.Name("users")
	table.ID("id")
	table.Col("name").Varchar(255).NotNull().Default("''")
	table.Col("email").Varchar(100).NotNull().Unique()

	ct := (&CreateTable{table: table}).IfNotExists()
	assert.Equal(t, `CREATE TABLE IF NOT EXISTS "users" (id SERIAL, name VARCHAR(255) NOT NULL DEFAULT '', email VARCHAR(100) UNIQUE NOT NULL, PRIMARY KEY(id));`, ct.SQL())
}
