package handler

import (
	"account-transactions/db"
	l "account-transactions/infrastructure/log"
	"context"
	"encoding/json"
	"errors"
)

type Handler interface {
	Handle(ctx context.Context, model EventModel) error
}

type DbWriteHandler struct {
	repository db.AccountRepository
}

func NewDbWriteHandler(repository db.AccountRepository) *DbWriteHandler {
	return &DbWriteHandler{repository: repository}
}

func (d DbWriteHandler) Handle(ctx context.Context, model EventModel) error {
	log := l.Logger.WithField(TraceIdKey, ctx.Value(TraceIdKey))

	switch model.Type {
	case WithdrawFromAccountV1:
		log.Debugf("received %s event", WithdrawFromAccountV1)
		var event WithdrawV1
		err := json.Unmarshal(model.Data, &event)
		if err != nil {
			return err
		}
		return d.repository.DecreaseBalance(model.AggregateId, event.Amount)

	case DepositToAccountV1:
		log.Debugf("received %s event", DepositToAccountV1)
		var event DepositV1
		err := json.Unmarshal(model.Data, &event)
		if err != nil {
			return err
		}
		return d.repository.IncreaseBalance(model.AggregateId, event.Amount)

	case OpenAccountV1:
		log.Debugf("received %s event", OpenAccountV1)
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
		log.Debugf("received %s event", CloseAccountV1)
		var event CloseV1
		err := json.Unmarshal(model.Data, &event)
		if err != nil {
			return err
		}
		return d.repository.SetInactive(model.AggregateId)

	case ClearEvent:
		log.Debugf("received %s event", ClearEvent)
		return d.repository.Delete(model.AggregateId)

	default:
		return errors.New("event type cannot be handled")
	}
}
