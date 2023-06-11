package suacli

import (
	"github.com/dundunlabs/sua"
	migrcli "github.com/dundunlabs/sua/cli/app/migr"
	"github.com/dundunlabs/sua/migr"
	"github.com/urfave/cli/v3"
)

func NewApp(db *sua.DB, migrations migr.Migrations) *App {
	return &App{
		App:        cliApp,
		db:         db,
		migrations: migrations,
	}
}

var cliApp = &cli.App{
	Name:  "sua",
	Usage: "db management toolkit",
	Commands: []*cli.Command{
		{
			Name:  "migr",
			Usage: "generate, migrate, rollback migrations",
			Commands: []*cli.Command{
				{
					Name:  "gen",
					Usage: "generate migration",
					Flags: []cli.Flag{
						&cli.BoolFlag{
							Name:  "sql",
							Usage: "generate sql migration",
						},
					},
					Action: func(ctx *cli.Context) error {
						sql := ctx.Bool("sql")
						if sql {
							return migrcli.GenerateSQLs(ctx)
						}
						return nil
					},
				},
			},
		},
	},
}

type App struct {
	*cli.App
	db         *sua.DB
	migrations migr.Migrations
}
