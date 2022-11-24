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

func (a *AccountRepository) IncreaseBalance(id string, amount uint64) error {

	tx, err := a.conn.BeginTx(context.Background(), &sql.TxOptions{
		Isolation: sql.LevelWriteCommitted,
		ReadOnly:  false,
	})
	if err != nil {
		return err
	}
	defer tx.Rollback()

	updateStmt, err := tx.Prepare(`UPDATE account SET balance = balance + $1 WHERE id == $2`)
	defer closeIgnoring(updateStmt)

	if err != nil {
		return err
	}

	ex, err := updateStmt.Exec(amount, id)
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

func (a *AccountRepository) DecreaseBalance(id string, amount uint64) error {

	tx, err := a.conn.BeginTx(context.Background(), &sql.TxOptions{
		Isolation: sql.LevelWriteCommitted,
		ReadOnly:  false,
	})
	if err != nil {
		return err
	}
	defer tx.Rollback()

	updateStmt, err := tx.Prepare(`UPDATE account SET balance = balance - $1 WHERE id == $2`)
	defer closeIgnoring(updateStmt)

	if err != nil {
		return err
	}

	ex, err := updateStmt.Exec(amount, id)
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
