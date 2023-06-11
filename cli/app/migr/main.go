package migr

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/dundunlabs/sua/cli/constants"
	"github.com/urfave/cli/v3"
)

const (
	timeFormat = "20060102150405"
	file
)

func GenerateSQLs(ctx *cli.Context) error {
	migrName := genMigrName(ctx)
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

func genMigrName(ctx *cli.Context) string {
	ver := genVersion()
	name := strings.Join(ctx.Args().Slice(), "_")
	return fmt.Sprintf("%s_%s", ver, name)
}

func genVersion() string {
	return time.Now().Format(timeFormat)
}
