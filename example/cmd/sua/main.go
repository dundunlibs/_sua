package main

import (
	"log"
	"os"

	suacli "github.com/dundunlabs/sua/cli/app"
	"github.com/dundunlabs/sua/example/sua"
	"github.com/dundunlabs/sua/example/sua/migrations"
)

func main() {
	db, err := sua.NewDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	app := suacli.NewApp(db, migrations.Migrations)
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
