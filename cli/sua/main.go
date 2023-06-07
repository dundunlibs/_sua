package main

import (
	"log"
	"os"
	"strings"

	"github.com/dundunlabs/sua/x/migration"
	"github.com/urfave/cli/v3"
)

func main() {
	migrator := migration.NewMigrator()

	suacli := &cli.App{
		Name:  "sua",
		Usage: "DB management toolkit",
		Commands: []*cli.Command{
			{
				Name:    "generate",
				Aliases: []string{"g"},
				Usage:   "generate sua's resources",
				Commands: []*cli.Command{
					{
						Name:    "migration",
						Aliases: []string{"migr"},
						Usage:   "generate migration",
						Flags: []cli.Flag{
							&cli.BoolFlag{
								Name:  "sql",
								Usage: "generate sql migration",
							},
						},
						Action: func(ctx *cli.Context) error {
							name := strings.Join(ctx.Args().Slice(), "_")
							sql := ctx.Bool("sql")
							return migrator.GenerateMigration(name, sql)
						},
					},
				},
			},
		},
	}

	if err := suacli.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
