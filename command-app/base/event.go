package base

import (
	"encoding/json"
	"time"
)

type EventModel struct {
	Event         `bson:"event"`
	Id            string    `bson:"_id"`
	Date          time.Time `bson:"date"`
	AggregateId   string    `bson:"aggregate_id"`
	AggregateType string    `bson:"aggregate_type"`
	Version       int64     `bson:"version"`
}

type Event struct {
	Id   string `bson:"event_id"`
	Type string `bson:"event_type"`
	Data []byte `bson:"data"`
}

func NewEvent(id string, eventType string, data interface{}) *Event {
	event, err := json.Marshal(data)
	if err != nil {
		return nil
	}
	return &Event{id, eventType, event}
}
