package handler

type AccountRepository interface {
	IncreaseBalance(id string, amount uint64) error
	DecreaseBalance(id string, amount uint64) error
}

type MockAccountRepository struct {
}

func (m *MockAccountRepository) IncreaseBalance(id string, amount uint64) error {
	return nil
}

func (m *MockAccountRepository) DecreaseBalance(id string, amount uint64) error {
	return nil
}
