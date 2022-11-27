package infrastructure

import (
	"github.com/offlinebrain/cqrs-bank-example/command-app/base"
	"github.com/offlinebrain/cqrs-bank-example/command-app/infrastructure/metrics"
	"github.com/prometheus/client_golang/prometheus"
)

type EventRepository interface {
	FindByAggregateId(id string) []base.EventModel
	Save(model *base.EventModel)
	GetAggregateIds(aggregateType string) []string
}

type PromMongoEventRepository struct {
	EventRepository
}

func NewPromMongoEventRepository(eventRepository EventRepository) *PromMongoEventRepository {
	return &PromMongoEventRepository{EventRepository: eventRepository}
}

func (p *PromMongoEventRepository) FindByAggregateId(id string) []base.EventModel {
	timer := prometheus.NewTimer(prometheus.ObserverFunc(func(f float64) {
		metrics.TotalMongoAccessDuration.WithLabelValues("read").Observe(f)
	}))
	defer timer.ObserveDuration()
	return p.EventRepository.FindByAggregateId(id)
}

func (p *PromMongoEventRepository) Save(model *base.EventModel) {
	timer := prometheus.NewTimer(prometheus.ObserverFunc(func(f float64) {
		metrics.TotalMongoAccessDuration.WithLabelValues("write").Observe(f)
	}))
	defer timer.ObserveDuration()
	p.EventRepository.Save(model)
}

func (p *PromMongoEventRepository) GetAggregateIds(aggregateType string) []string {
	timer := prometheus.NewTimer(prometheus.ObserverFunc(func(f float64) {
		metrics.TotalMongoAccessDuration.WithLabelValues("read").Observe(f)
	}))
	defer timer.ObserveDuration()
	return p.EventRepository.GetAggregateIds(aggregateType)
}
