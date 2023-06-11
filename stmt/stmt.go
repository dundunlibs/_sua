package stmt

import (
	"context"
	"database/sql"

	"github.com/dundunlabs/sua/core"
)

type Stmt interface {
	SQL() string
}

type BaseStmt struct {
	self Stmt
	db   *core.DB
}

func (stmt *BaseStmt) Exec(ctx context.Context) (sql.Result, error) {
	query := stmt.self.SQL()
	return stmt.db.ExecContext(ctx, query)
}
