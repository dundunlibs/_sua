package stmt

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/dundunlabs/sua/internal"
)

type BaseStmt struct {
	self stmt
	db   *internal.DB
}

func (stmt *BaseStmt) Exec(ctx context.Context) (sql.Result, error) {
	query := stmt.self.SQL()
	fmt.Println(query)
	return stmt.db.ExecContext(ctx, query)
}

type stmt interface {
	SQL() string
}
