package main

import (
	"log"
	"os"

	"github.com/dundunlabs/sua/cli/initialization"
	"github.com/urfave/cli/v3"
)

func main() {
	app := &cli.App{
		Name:  "sua",
		Usage: "An ORM package for Go",
		Commands: []*cli.Command{
			{
				Name:  "init",
				Usage: "Initialize project",
				Action: func(ctx *cli.Context) error {
					return initialization.Generate()
				},
			},
		},
	}
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
