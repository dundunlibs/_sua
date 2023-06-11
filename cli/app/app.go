package suacli

import (
	"strings"

	"github.com/dundunlabs/sua"
	"github.com/dundunlabs/sua/migr"
	"github.com/urfave/cli/v3"
)

func NewApp(db *sua.DB, migrations migr.Migrations) *App {
	migrator := migr.NewMigrator(db, migrations)
	app := &App{
		migrator: migrator,
	}
	app.App = newCLI(app)

	return app
}

type App struct {
	*cli.App
	migrator *migr.Migrator
}

func (app *App) Migrate(ctx *cli.Context) error {
	return app.migrator.Migrate(ctx.Context)
}

func (app *App) Rollback(ctx *cli.Context) error {
	return app.migrator.Rollback(ctx.Context)
}

func (app *App) Generate(ctx *cli.Context) error {
	name := strings.Join(ctx.Args().Slice(), "_")
	sql := ctx.Bool("sql")
	if sql {
		return genSQLs(name)
	}
	return nil
}
