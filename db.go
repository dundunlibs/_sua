package sua

import (
	"context"
	"database/sql"
	"fmt"
)

func NewDB(sqldb *sql.DB) *DB {
	return &DB{sqldb}
}

type DB struct {
	*sql.DB
}

func (db *DB) ExecContext(ctx context.Context, query string, args ...any) (sql.Result, error) {
	fmt.Println(query)
	return db.DB.ExecContext(ctx, query, args...)
}

func (db *DB) WithTx(ctx context.Context, opts *sql.TxOptions, fn func(tx *Tx) error) error {
	tx, err := db.BeginTx(ctx, opts)
	if err != nil {
		return err
	}

	if err := fn(&Tx{db, tx}); err != nil {
		if err := tx.Rollback(); err != nil {
			return err
		}
		return err
	}

	return tx.Commit()
}
