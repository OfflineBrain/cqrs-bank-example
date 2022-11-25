package usecase

import (
	"query-app/db"
)

type AccountRequest struct {
	id string
}

func NewAccountRequest(id string) *AccountRequest {
	return &AccountRequest{id: id}
}

type GetAccountResponse struct {
	Id         string `json:"id"`
	HolderName string `json:"holder_name"`
	Balance    uint64 `json:"balance"`
	Active     bool   `json:"active"`
}

type GetAccountUseCase struct {
	ar db.AccountRepository
}

func NewGetAccountUseCase(ar db.AccountRepository) *GetAccountUseCase {
	return &GetAccountUseCase{ar: ar}
}

func (g *GetAccountUseCase) Request(request AccountRequest) *GetAccountResponse {
	acc, err := g.ar.Get(request.id)
	if err != nil {
		return nil
	}

	return &GetAccountResponse{
		Id:         acc.Id,
		HolderName: acc.HolderName,
		Balance:    acc.Balance,
		Active:     acc.Active,
	}
}
