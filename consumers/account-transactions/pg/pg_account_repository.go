package pg

import (
	"account-transactions/handler"
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

	err := RunInTx(context.Background(), a.conn, func(ctx context.Context, tx *sql.Tx) error {
		stmt, err := tx.PrepareContext(ctx, `INSERT INTO account(id, holder_name, balance, active) VALUES ($1, $2, $3, $4)`)
		if err != nil {
			return err
		}
		defer closeIgnoring(stmt)

		_, err = stmt.Exec(account.Id, account.HolderName, account.Balance, account.Active)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}

func (a *AccountRepository) IncreaseBalance(id string, amount uint64) error {

	err := RunInTx(context.Background(), a.conn, func(ctx context.Context, tx *sql.Tx) error {
		stmt, err := tx.Prepare(`UPDATE account SET balance = balance + $1 WHERE id = $2`)
		if err != nil {
			return err
		}
		defer closeIgnoring(stmt)

		_, err = stmt.Exec(amount, id)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}

func (a *AccountRepository) DecreaseBalance(id string, amount uint64) error {

	err := RunInTx(context.Background(), a.conn, func(ctx context.Context, tx *sql.Tx) error {
		stmt, err := tx.Prepare(`UPDATE account SET balance = balance - $1 WHERE id = $2`)
		if err != nil {
			return err
		}
		defer closeIgnoring(stmt)

		_, err = stmt.Exec(amount, id)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}

func (a *AccountRepository) SetInactive(id string) error {
	err := RunInTx(context.Background(), a.conn, func(ctx context.Context, tx *sql.Tx) error {
		stmt, err := tx.Prepare(`UPDATE account SET active = false WHERE id = $1`)
		if err != nil {
			return err
		}
		defer closeIgnoring(stmt)

		_, err = stmt.Exec(id)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return err
	}

	return nil
}

func closeIgnoring(closer io.Closer) {
	_ = closer.Close()
}
