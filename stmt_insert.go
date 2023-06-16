package sua

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
)

func (db *DB) Insert(table string, data map[string]any) Insert {
	return Insert{
		db:    db,
		table: table,
		data:  data,
	}
}

type Insert struct {
	db    *DB
	table string
	data  map[string]any
}

func (i Insert) SQL() (string, []any) {
	keys := []string{}
	params := []string{}
	values := []any{}

	idx := 1
	for key, value := range i.data {
		keys = append(keys, key)
		params = append(params, fmt.Sprintf("$%d", idx))
		values = append(values, value)
		idx++
	}

	return fmt.Sprintf(
		"INSERT INTO %q (%s) VALUES (%s)",
		i.table,
		strings.Join(keys, ", "),
		strings.Join(params, ", "),
	), values
}

func (i Insert) Exec(ctx context.Context) (sql.Result, error) {
	query, args := i.SQL()
	return i.db.ExecContext(ctx, query, args...)
}
