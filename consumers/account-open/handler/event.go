package handler

import "time"

const (
	OpenAccountV1 = "openAccountV1"
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

type OpenV1 struct {
	HolderName string `json:"holder_name"`
	Balance    uint64 `json:"balance"`
}
