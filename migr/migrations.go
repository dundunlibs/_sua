package migr

import (
	"io/fs"
	"strings"
)

func NewMigrations() Migrations {
	return Migrations{}
}

type Migrations map[string]Migration

func (ms Migrations) Load(fsys fs.FS) error {
	return fs.WalkDir(fsys, ".", func(path string, d fs.DirEntry, err error) error {
		if !strings.HasSuffix(path, ".sql") {
			return nil
		}
		name, ext, _ := strings.Cut(path, ".")
		bytes, err := fs.ReadFile(fsys, path)
		if err != nil {
			return err
		}
		data := string(bytes)
		if ms[name] == nil {
			ms[name] = &MigrationSQL{}
		}
		m := ms[name].(*MigrationSQL)
		switch ext {
		case "up.sql":
			m.up = data
		case "down.sql":
			m.down = data
		}
		return nil
	})
}
