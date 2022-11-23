package handler

type AccountRepository interface {
	Save(account Account) error
}

type MockAccountRepository struct {
}

func (m MockAccountRepository) Save(account Account) error {
	return nil
}
