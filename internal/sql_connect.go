package iternal

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func ConnectToDb( connStr string) (*sql.DB, error) {
	db, err := sql.Open("postres", connStr)
	if err != nil {
		Logger.Error(err)
		return nil, err
	}
	return db, nil
}