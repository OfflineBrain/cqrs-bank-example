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
	if model.Type == CloseAccountV1 {
		var event CloseV1
		err := json.Unmarshal(model.Data, &event)
		if err != nil {
			return err
		}

		return d.repository.SetInactive(model.AggregateId)
	}
	return errors.New("event type cannot be handled")
}
