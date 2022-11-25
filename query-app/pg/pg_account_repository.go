package pg

import (
	"context"
	"database/sql"
	"query-app/db/entity"
)

type AccountRepository struct {
	conn *sql.DB
}

func NewAccountRepository(conn *sql.DB) *AccountRepository {
	return &AccountRepository{conn: conn}
}

func (a *AccountRepository) Get(id string) (account entity.Account, err error) {

	err = RunInTx(context.Background(), a.conn, func(ctx context.Context, tx *sql.Tx) error {
		rows := tx.QueryRowContext(ctx, `SELECT a.id, a.holder_name, a.balance, a.active FROM account a WHERE id = $1`, id)
		err := rows.Scan(&account.Id, &account.HolderName, &account.Balance, &account.Active)
		if err != nil || err != sql.ErrNoRows {
			return err
		}

		return nil
	})
	return
}
