package handler

import "time"

const (
	CloseAccountV1 = "closeAccountV1"
)

type EventModel struct {
	Event         `json:"event"`
	Id            string    `json:"id"`
	Date          time.Time `json:"date"`
	AggregateId   string    `json:"aggregateId"`
	AggregateType string    `json:"aggregateType"`
	Version       int64     `json:"version"`
}

type Event struct {
	Id   string `json:"id"`
	Type string `json:"type"`
	Data []byte `json:"data"`
}

type CloseV1 struct {
}
