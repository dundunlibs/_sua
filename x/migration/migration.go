package migration

import (
	"fmt"
	"os"
	"time"
)

const (
	migrationsDir = "sua/migrations"
	timeFormat    = "20060102150405"
	dirPerm       = 0o755
	filePerm      = 0o644
)

var (
	goExt   = "go"
	sqlExts = []string{"up.sql", "down.sql"}
)

func NewMigrator() *Migrator {
	return &Migrator{}
}

type Migrator struct {
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
