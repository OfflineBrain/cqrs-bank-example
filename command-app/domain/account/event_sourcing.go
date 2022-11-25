package account

import (
	"context"
	"github.com/offlinebrain/cqrs-bank-example/command-app/base"
	"github.com/offlinebrain/cqrs-bank-example/command-app/infrastructure"
)

type AggregateRepository struct {
	infrastructure.AggregateRepositoryBase[*Aggregate]
}

func NewAggregateRepository(store *infrastructure.EventStore) infrastructure.AggregateRepository[*Aggregate] {
	return &AggregateRepository{AggregateRepositoryBase: *infrastructure.NewAggregateRepositoryBase[*Aggregate](store, aggregate)}
}

func (e *AggregateRepository) Get(id string) (*Aggregate, error) {
	return e.AggregateRepositoryBase.Get(id)
}

func (e *AggregateRepository) Save(ctx context.Context, aggregate *Aggregate) error {
	return e.AggregateRepositoryBase.Save(ctx, aggregate)
}

func aggregate(id string) *Aggregate {
	return &Aggregate{
		AggregateRoot: base.AggregateRoot{Id: id, Version: -1},
	}
}
