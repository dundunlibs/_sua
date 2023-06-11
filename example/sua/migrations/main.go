package migrations

import (
	"embed"

	"github.com/dundunlabs/sua/migr"
)

var Migrations = migr.NewMigrations()

//go:embed *
var fs embed.FS

func init() {
	if err := Migrations.Load(fs); err != nil {
		panic(err)
	}
}
