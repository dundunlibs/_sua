package sua

import (
	"context"
	"database/sql"
	"fmt"
)

type Tx struct {
	*DB
	*sql.Tx
}

func (tx *Tx) ExecContext(ctx context.Context, query string, args ...any) (sql.Result, error) {
	fmt.Println(query)
	return tx.Tx.ExecContext(ctx, query, args...)
}
