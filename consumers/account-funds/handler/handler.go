package handler

import (
	"encoding/json"
	"errors"
)

type Handler interface {
	Handle(model EventModel) error
}

type DbWriteHandler struct {
	repository AccountRepository
}

func NewDbWriteHandler(repository AccountRepository) *DbWriteHandler {
	return &DbWriteHandler{repository: repository}
}

func (d DbWriteHandler) Handle(model EventModel) error {
	switch model.Type {
	case WithdrawFromAccountV1:
		var event WithdrawV1
		err := json.Unmarshal(model.Data, &event)
		if err != nil {
			return err
		}
		return d.repository.DecreaseBalance(model.AggregateId, event.Amount)
	case DepositToAccountV1:
		var event DepositV1
		err := json.Unmarshal(model.Data, &event)
		if err != nil {
			return err
		}
		return d.repository.IncreaseBalance(model.AggregateId, event.Amount)
	default:
		return errors.New("event type cannot be handled")
	}
}
