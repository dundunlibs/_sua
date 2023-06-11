package suacli

import (
	"github.com/urfave/cli/v3"
)

func newCLI(app *App) *cli.App {
	return &cli.App{
		Name:  "sua",
		Usage: "DB management toolkit",
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
						Action: app.Generate,
					},
					{
						Name:   "migrate",
						Usage:  "migrate DB to latest version",
						Action: app.Migrate,
					},
					{
						Name:   "rollback",
						Usage:  "rollback DB to latest version",
						Action: app.Rollback,
					},
				},
			},
		},
	}
}
