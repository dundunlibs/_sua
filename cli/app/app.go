package suacli

import (
	"github.com/dundunlabs/sua"
	"github.com/dundunlabs/sua/migr"
	"github.com/urfave/cli/v3"
)

func NewApp(db *sua.DB, migrations migr.Migrations) *App {
	return &App{
		App:        &cli.App{},
		db:         db,
		migrations: migrations,
	}
}

type App struct {
	*cli.App
	db         *sua.DB
	migrations migr.Migrations
}
