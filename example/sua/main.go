package sua

import (
	"database/sql"

	"github.com/dundunlabs/sua"
	_ "github.com/lib/pq"
)

func NewDB() (*sua.DB, error) {
	sqldb, err := sql.Open("postgres", "postgresql://dundun:@localhost:5432/sua-example?sslmode=disable")
	if err != nil {
		return nil, err
	}
	return sua.NewDB(sqldb), nil
}
