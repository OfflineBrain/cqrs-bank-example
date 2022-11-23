package base

import (
	"encoding/json"
	"time"
)

type EventModel struct {
	Event         `bson:"event" json:"event"`
	Id            string    `bson:"_id" json:"id"`
	Date          time.Time `bson:"date" json:"date"`
	AggregateId   string    `bson:"aggregate_id" json:"aggregateId"`
	AggregateType string    `bson:"aggregate_type" json:"aggregateType"`
	Version       int64     `bson:"version" json:"version"`
}

type Event struct {
	Id   string `bson:"event_id" json:"id"`
	Type string `bson:"event_type" json:"type"`
	Data []byte `bson:"data" json:"data"`
}

func NewEvent(id string, eventType string, data interface{}) *Event {
	event, err := json.Marshal(data)
	if err != nil {
		return nil
	}
	return &Event{id, eventType, event}
}
