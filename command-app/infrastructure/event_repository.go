package infrastructure

import "github.com/offlinebrain/cqrs-bank-example/command-app/base"

type EventRepository interface {
	FindByAggregateId(id string) []base.EventModel
	Save(model *base.EventModel)
	GetAggregateIds(aggregateType string) []string
}
