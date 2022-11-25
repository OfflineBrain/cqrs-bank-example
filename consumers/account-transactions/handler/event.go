package handler

import "time"

const (
	OpenAccountV1         = "openAccountV1"
	DepositToAccountV1    = "depositToAccountV1"
	WithdrawFromAccountV1 = "withdrawFromAccountV1"
	CloseAccountV1        = "closeAccountV1"
)

const TraceIdKey = "trace_id"

type TracedEventModel struct {
	TraceId    string `json:"trace_id"`
	EventModel `json:"event_model"`
}

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
}

type DepositV1 struct {
	Amount uint64 `json:"amount"`
}

type WithdrawV1 struct {
	Amount uint64 `json:"amount"`
}

type CloseV1 struct {
}
