package migration

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/dundunlabs/sua"
	"github.com/dundunlabs/sua/stmt"
)

const (
	migrationsDir = "sua/migrations"
	timeFormat    = "20060102150405"
	migrTable     = "_sua_migrations"
	dirPerm       = 0o755
	filePerm      = 0o644
)

var (
	goExt   = "go"
	sqlExts = []string{"up.sql", "down.sql"}
)

func NewMigrator(db *sua.DB) *Migrator {
	return &Migrator{
		DB: db,
	}
}

type Migrator struct {
	*sua.DB
}

func (m *Migrator) Migrate(ctx context.Context) error {
	if err := m.Init(ctx); err != nil {
		return err
	}
	return nil
}

func (m *Migrator) Init(ctx context.Context) error {
	_, err := m.CreateTable(migrTable, func(stmt *stmt.CreateTableStmt) {
		stmt.Col("name").Varchar(255).Unique()
		stmt.Col("migrated_at").Timestamp(3).Default("current_timestamp")
		stmt.Col("rollbacked_at").Timestamp(3).Nullable()
	}).IfNotExists().Exec(ctx)
	return err
}

func (m *Migrator) GenerateMigration(name string, sql bool) error {
	err := os.MkdirAll(migrationsDir, dirPerm)
	if err != nil && !os.IsExist(err) {
		return err
	}
	if sql {
		return generateSQLMigration(name)
	}
	return generateGoMigration(name)
}

func generateSQLMigration(name string) error {
	for _, ext := range sqlExts {
		path := generateFilename(name, ext)
		if err := os.WriteFile(path, nil, filePerm); err != nil {
			return err
		}
	}
	return nil
}

func generateGoMigration(name string) error {
	path := generateFilename(name, goExt)
	return os.WriteFile(path, nil, filePerm)
}

func generateFilename(name string, ext string) string {
	ver := generateVersion()
	return fmt.Sprintf("%s/%s_%s.%s", migrationsDir, ver, name, ext)
}

func generateVersion() string {
	return time.Now().UTC().Format(timeFormat)
}
