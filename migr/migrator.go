package migr

import (
	"context"
	"database/sql"
	"sort"

	"github.com/dundunlabs/sua"
	"github.com/dundunlabs/xidau/maps"
)

func NewMigrator(db *sua.DB, migrations Migrations) *Migrator {
	return &Migrator{
		db:         db,
		migrations: migrations,
	}
}

type Migrator struct {
	db         *sua.DB
	migrations Migrations
}

func (m *Migrator) Migrate(ctx context.Context) error {
	ms := m.sortedMigrations()
	for _, k := range ms {
		migr := m.migrations[k]
		if err := m.db.ExecInTx(ctx, &sql.TxOptions{}, func(tx *sql.Tx) error {
			_, err := tx.ExecContext(ctx, migr.Up())
			return err
		}); err != nil {
			return err
		}
	}
	return nil
}

func (m *Migrator) Rollback(ctx context.Context) error {
	ms := m.sortedMigrations()
	migr := m.migrations[ms[len(ms)-1]]
	return m.db.ExecInTx(ctx, &sql.TxOptions{}, func(tx *sql.Tx) error {
		_, err := tx.ExecContext(ctx, migr.Down())
		return err
	})
}

func (m *Migrator) sortedMigrations() []string {
	keys := maps.Keys(m.migrations)
	sort.Strings(keys)
	return keys
}
