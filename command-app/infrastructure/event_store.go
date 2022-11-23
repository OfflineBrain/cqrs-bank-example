package infrastructure

import (
	"errors"
	"github.com/offlinebrain/cqrs-bank-example/command-app/base"
	"time"
)

type EventStore struct {
	repository EventRepository
	publisher  EventPublisher
	topics     map[string]string
}

func NewEventStore(repository EventRepository, publisher EventPublisher, topics map[string]string) *EventStore {
	return &EventStore{repository: repository, publisher: publisher, topics: topics}
}

func (s *EventStore) SaveEvents(
	aggregateId string,
	aggregateType string,
	events []base.Event,
	expectedVersion int64,
) error {

	eventStream := s.repository.FindByAggregateId(aggregateId)
	if expectedVersion != -1 && eventStream[len(eventStream)-1].Version != expectedVersion {
		return errors.New("concurrency error")
	}

	version := expectedVersion

	for _, event := range events {
		version++
		model := &base.EventModel{
			Event:         event,
			AggregateType: aggregateType,
			AggregateId:   aggregateId,
			Date:          time.Now(),
			Version:       version,
		}
		s.repository.Save(model)
		topic := s.topics[event.Type]
		_ = s.publisher.Publish(topic, *model)
	}

	return nil
}

func (s *EventStore) GetEvents(aggregateId string) ([]base.Event, error) {
	eventStream := s.repository.FindByAggregateId(aggregateId)
	if eventStream == nil || len(eventStream) == 0 {
		return []base.Event{}, errors.New("incorrect account id")
	}

	events := make([]base.Event, len(eventStream))
	for i, model := range eventStream {
		events[i] = model.Event
	}

	return events, nil
}
