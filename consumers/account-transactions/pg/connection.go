package pg

import (
	"database/sql"
	"fmt"
)
import _ "github.com/lib/pq"

func NewPgConnection(
	host string,
	port int32,
	user string,
	password string,
	dbname string,
) (*sql.DB, error) {
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

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
