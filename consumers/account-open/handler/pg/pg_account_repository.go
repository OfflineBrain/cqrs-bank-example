package pg

import (
	"account-open/handler"
	"database/sql"
)

type AccountRepository struct {
	conn *sql.DB
}

func NewAccountRepository(conn *sql.DB) *AccountRepository {
	return &AccountRepository{conn: conn}
}

func (a *AccountRepository) Save(account handler.Account) error {

	prepare, err := a.conn.Prepare(`INSERT INTO account(id, holder_name, balance, active) VALUES ($1, $2, $3, $4)`)
	if err != nil {
		return err
	}

	ex, err := prepare.Exec(account.Id, account.HolderName, account.Balance, account.Active)
	if err != nil {
		return err
	}

	print(ex)
	return nil
}
