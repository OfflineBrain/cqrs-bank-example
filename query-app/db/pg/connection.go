package pg

import (
	"database/sql"
)
import _ "github.com/lib/pq"

func NewPgConnection(
	psqlconn string,
) (*sql.DB, error) {
	db, err := sql.Open("postgres", psqlconn)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
