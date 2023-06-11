package stmt

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateTable(t *testing.T) {
	ct := NewCreateTable("users")
	ct.ID("id")
	ct.Col("name").Varchar(255).NotNull().Default("''")
	ct.Col("email").Varchar(100).NotNull().Unique()

	stmt := NewStmtCreateTable(ct, nil).IfNotExists()
	assert.Equal(t, `CREATE TABLE IF NOT EXISTS "users" (id SERIAL, name VARCHAR(255) NOT NULL DEFAULT '', email VARCHAR(100) UNIQUE NOT NULL, PRIMARY KEY(id));`, stmt.SQL())
}
