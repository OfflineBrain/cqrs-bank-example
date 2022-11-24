package pg

import (
	"account-open/handler"
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

func (a *AccountRepository) Save(account handler.Account) error {

	tx, err := a.conn.BeginTx(context.Background(), &sql.TxOptions{
		Isolation: sql.LevelReadCommitted,
		ReadOnly:  false,
	})
	if err != nil {
		return err
	}
	defer tx.Rollback()

	stmt, err := tx.Prepare(`INSERT INTO account(id, holder_name, balance, active) VALUES ($1, $2, $3, $4)`)
	if err != nil {
		return err
	}
	defer closeIgnoring(stmt)

	ex, err := stmt.Exec(account.Id, account.HolderName, account.Balance, account.Active)
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
