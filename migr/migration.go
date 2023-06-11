package migr

type Migration interface {
	Up() string
	Down() string
}

type MigrationSQL struct {
	up   string
	down string
}

func (m *MigrationSQL) Up() string {
	return m.up
}

func (m *MigrationSQL) Down() string {
	return m.down
}
