package suacli

import (
	"fmt"
	"os"
	"time"

	"github.com/dundunlabs/sua/cli/constants"
)

const (
	timeFormat = "20060102150405"
)

func genSQLs(name string) error {
	migrName := genMigrName(name)
	for _, suffix := range []string{"up", "down"} {
		if err := genMigr(fmt.Sprintf("%s.%s.sql", migrName, suffix), ""); err != nil {
			return err
		}
	}
	return nil
}

func genMigr(filename string, data string) error {
	return os.WriteFile(fmt.Sprintf("%s/%s/%s", constants.SuaDir, constants.MigrDir, filename), []byte(data), constants.FilePerm)
}

func genMigrName(name string) string {
	ver := genVersion()
	return fmt.Sprintf("%s_%s", ver, name)
}

func genVersion() string {
	return time.Now().Format(timeFormat)
}
