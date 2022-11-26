package db

import "account-transactions/handler"

type AccountRepository interface {
	Save(account handler.Account) error
	IncreaseBalance(id string, amount uint64) error
	DecreaseBalance(id string, amount uint64) error
	SetInactive(id string) error
	Delete(id string) error
}

type NoOpAccountRepository struct {
}

func (m *NoOpAccountRepository) SetInactive(_ string) error {
	return nil
}

func (m *NoOpAccountRepository) IncreaseBalance(_ string, _ uint64) error {
	return nil
}

func (m *NoOpAccountRepository) DecreaseBalance(_ string, _ uint64) error {
	return nil
}

func (m *NoOpAccountRepository) Save(_ handler.Account) error {
	return nil
}

func (m *NoOpAccountRepository) Delete(string) error {
	return nil
}
