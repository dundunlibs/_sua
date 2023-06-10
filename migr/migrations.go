package migr

func NewMigrations() Migrations {
	return Migrations{}
}

type Migrations []Migration
