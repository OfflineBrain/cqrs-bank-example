package handler

type AccountRepository interface {
	Save(account Account) error
	IncreaseBalance(id string, amount uint64) error
	DecreaseBalance(id string, amount uint64) error
	SetInactive(id string) error
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

func (m *NoOpAccountRepository) Save(_ Account) error {
	return nil
}
