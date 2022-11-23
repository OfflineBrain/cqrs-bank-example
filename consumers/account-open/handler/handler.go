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
	if model.Type == OpenAccountV1 {
		var event OpenV1
		err := json.Unmarshal(model.Data, &event)
		if err != nil {
			return err
		}

		a := &Account{
			Id:         model.AggregateId,
			HolderName: event.HolderName,
			Balance:    event.Balance,
			Active:     true,
		}

		return d.repository.Save(*a)
	}
	return errors.New("event type cannot be handled")
}
