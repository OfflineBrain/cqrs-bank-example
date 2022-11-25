package handler

import (
	"context"
	"encoding/json"
	"errors"
)

type Handler interface {
	Handle(ctx context.Context, model EventModel) error
}

type DbWriteHandler struct {
	repository AccountRepository
}

func NewDbWriteHandler(repository AccountRepository) *DbWriteHandler {
	return &DbWriteHandler{repository: repository}
}

func (d DbWriteHandler) Handle(ctx context.Context, model EventModel) error {
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
	case OpenAccountV1:
		var event OpenV1
		err := json.Unmarshal(model.Data, &event)
		if err != nil {
			return err
		}

		a := &Account{
			Id:         model.AggregateId,
			HolderName: event.HolderName,
			Balance:    0,
			Active:     true,
		}
		return d.repository.Save(*a)
	case CloseAccountV1:
		var event CloseV1
		err := json.Unmarshal(model.Data, &event)
		if err != nil {
			return err
		}
		return d.repository.SetInactive(model.AggregateId)
	default:
		return errors.New("event type cannot be handled")
	}
}
