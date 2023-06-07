package main

import (
	"database/sql"
	"log"
	"os"
	"strings"

	"github.com/dundunlabs/sua"
	"github.com/dundunlabs/sua/x/migration"
	_ "github.com/joho/godotenv/autoload"
	_ "github.com/lib/pq"
	"github.com/urfave/cli/v3"
)

const dsnKey = "DATABASE_URL"

func main() {
	sqldb, err := sql.Open("postgres", os.Getenv(dsnKey))
	if err != nil {
		log.Fatal(err)
	}

	db := sua.NewDB(sqldb)
	defer db.Close()

	migrator := migration.NewMigrator(db)

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
			{
				Name:  "db",
				Usage: "manage db",
				Commands: []*cli.Command{
					{
						Name:  "migrate",
						Usage: "migrate db",
						Action: func(ctx *cli.Context) error {
							return migrator.Migrate(ctx.Context)
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
