package initialization

import (
	"fmt"
	"os"
	"os/exec"

	"golang.org/x/mod/modfile"
)

const (
	suaDir   = "sua"
	migrDir  = "migrations"
	cmdDir   = "cmd/sua"
	dirPerm  = 0755
	filePerm = 0644
)

const migrMain = `package migrations

import "github.com/dundunlabs/sua/migr"

var Migrations = migr.NewMigrations()
`

const suaDB = `package sua

import (
	"database/sql"

	"github.com/dundunlabs/sua"
)

func NewDB() (*sua.DB, error) {
	sqldb, err := sql.Open("", "")
	if err != nil {
		return nil, err
	}
	return sua.NewDB(sqldb), nil
}
`

const cmdMain = `package main

import (
	"log"
	"os"

	suacli "github.com/dundunlabs/sua/cli/app"
	"%[1]s/%[2]s"
	"%[1]s/%[2]s/%[3]s"
)

func main() {
	db, err := sua.NewDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	app := suacli.NewApp(db, %[3]s.Migrations)
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
`

func Generate() error {
	for _, gen := range []genFunc{genDB, genMigrations, genCmd} {
		if err := gen(); err != nil {
			return err
		}
	}
	return tidy()
}

type genFunc func() error

func genDB() error {
	return genFile(suaDir, "db.go", suaDB)
}

func genMigrations() error {
	return genFile(fmt.Sprintf("%s/%s", suaDir, migrDir), "main.go", migrMain)
}

func genCmd() error {
	path, err := modPath()
	if err != nil {
		return err
	}
	str := fmt.Sprintf(cmdMain, path, suaDir, migrDir)
	return genFile(cmdDir, "main.go", str)
}

func genFile(dir string, file string, data string) error {
	if err := mkdir(dir); err != nil {
		return err
	}
	return os.WriteFile(dir+"/main.go", []byte(data), filePerm)

}

func mkdir(dir string) error {
	err := os.MkdirAll(dir, dirPerm)
	if err != nil && !os.IsExist(err) {
		return err
	}
	return nil
}

func tidy() error {
	return exec.Command("go", "mod", "tidy").Run()
}

func modPath() (string, error) {
	mod, err := os.ReadFile("go.mod")
	if err != nil {
		return "", err
	}
	return modfile.ModulePath(mod), nil
}
