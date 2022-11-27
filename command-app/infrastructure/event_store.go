package infrastructure

import (
	"context"
	"errors"
	"github.com/offlinebrain/cqrs-bank-example/command-app/base"
	l "github.com/offlinebrain/cqrs-bank-example/command-app/infrastructure/log"
	"github.com/sirupsen/logrus"
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
	ctx context.Context,
	aggregateId string,
	aggregateType string,
	events []base.Event,
	expectedVersion int64,
) error {
	log := l.Logger.WithFields(logrus.Fields{
		TraceIdKey:       ctx.Value(TraceIdKey),
		AggregateIdKey:   aggregateId,
		AggregateTypeKey: aggregateType,
	})

	log.Infof("saving %d events", len(events))
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
		_ = s.publisher.Publish(ctx, topic, *model)
	}

	log.Debugf("saved successfuly")
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

func (s *EventStore) GetAggregateIds(aggregateType string) []string {
	return s.repository.GetAggregateIds(aggregateType)
}

func (s *EventStore) RepublishEvents(
	ctx context.Context,
	aggregateType string,
	clearEvent func(id string) *base.Event,
) error {
	log := l.Logger.WithField(TraceIdKey, ctx.Value(TraceIdKey))
	ids := s.GetAggregateIds(aggregateType)
	log.Infof("republishing events for %d %s aggregates", len(ids), aggregateType)
	for _, id := range ids {
		event := clearEvent(id)
		err := s.publisher.Publish(ctx, s.topics[event.Type], base.EventModel{
			Event:         *event,
			AggregateType: aggregateType,
			AggregateId:   id,
			Date:          time.Now(),
			Version:       -1,
		})
		if err != nil {
			return err
		}
		eventStream := s.repository.FindByAggregateId(id)
		if eventStream == nil {
			return errors.New("incorrect account id")
		}
		log.Infof("republishing %d events for %s aggregate [%s]", len(eventStream), aggregateType, id)

		err = s.republish(ctx, eventStream)
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *EventStore) republish(ctx context.Context, events []base.EventModel) error {

	for _, event := range events {
		err := s.publisher.Publish(ctx, s.topics[event.Type], event)
		if err != nil {
			return err
		}
	}
	return nil
}
