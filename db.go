package sua

import (
	"context"
	"database/sql"
)

func NewDB(sqldb *sql.DB) *DB {
	return &DB{sqldb}
}

type DB struct {
	*sql.DB
}

func (db *DB) ExecInTx(ctx context.Context, opts *sql.TxOptions, fn func(tx *sql.Tx) error) error {
	tx, err := db.BeginTx(ctx, opts)
	if err != nil {
		return err
	}

	if err := fn(tx); err != nil {
		if err := tx.Rollback(); err != nil {
			return err
		}
		return err
	}

	return tx.Commit()
}
