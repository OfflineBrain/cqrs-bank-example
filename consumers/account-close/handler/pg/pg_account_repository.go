package pg

import (
	"context"
	"database/sql"
	"io"
)

type AccountRepository struct {
	conn *sql.DB
}

func NewAccountRepository(conn *sql.DB) *AccountRepository {
	return &AccountRepository{conn: conn}
}

func (a *AccountRepository) SetInactive(id string) error {

	tx, err := a.conn.BeginTx(context.Background(), &sql.TxOptions{
		Isolation: sql.LevelWriteCommitted,
		ReadOnly:  false,
	})
	if err != nil {
		return err
	}
	defer tx.Rollback()

	stmt, err := tx.Prepare(`UPDATE account SET active = false WHERE id == $1`)
	defer closeIgnoring(stmt)

	if err != nil {
		return err
	}

	ex, err := stmt.Exec(id)
	if err != nil {
		return err
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return err
	}

	print(ex)
	return nil
}

func closeIgnoring(closer io.Closer) {
	_ = closer.Close()
}
