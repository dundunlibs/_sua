package sua

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
