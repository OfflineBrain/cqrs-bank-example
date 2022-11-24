package handler

type AccountRepository interface {
	SetInactive(id string) error
}

type MockAccountRepository struct {
}

func (m *MockAccountRepository) SetInactive(_ string) error {
	return nil
}
