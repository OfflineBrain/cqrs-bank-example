package infrastructure

import (
	"context"
	"github.com/offlinebrain/cqrs-bank-example/command-app/base"
)

type AggregateRepository[A base.Aggregate] interface {
	Exist(id string) bool
	Get(id string) (A, error)
	Save(ctx context.Context, aggregate A) error
}

type AggregateRepositoryBase[A base.Aggregate] struct {
	store     *EventStore
	aggregate func(id string) A
}

func NewAggregateRepositoryBase[A base.Aggregate](store *EventStore, aggregate func(id string) A) *AggregateRepositoryBase[A] {
	return &AggregateRepositoryBase[A]{store: store, aggregate: aggregate}
}

func (e *AggregateRepositoryBase[A]) Exist(id string) bool {
	events, err := e.store.GetEvents(id)
	if err != nil || len(events) == 0 {
		return false
	}
	return true
}

func (e *AggregateRepositoryBase[A]) Save(ctx context.Context, a A) error {
	err := e.store.SaveEvents(ctx, a.GetId(), a.GetType(), a.GetChanges(), a.GetVersion())
	if err != nil {
		return err
	}
	a.MarkCommitted()
	return nil
}

func (e *AggregateRepositoryBase[A]) Get(id string) (A, error) {
	events, err := e.store.GetEvents(id)
	if err != nil {
		return e.aggregate(id), err
	}

	var aggregate = e.aggregate(id)

	aggregate.Replay(events)
	return aggregate, nil
}
