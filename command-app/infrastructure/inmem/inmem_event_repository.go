package account_inmem

import (
	"github.com/google/uuid"
	"github.com/offlinebrain/cqrs-bank-example/command-app/base"
)

type InmemRepository struct {
	Cache map[string][]base.EventModel
}

func (i *InmemRepository) GetAggregateIds(string) []string {
	ids := make([]string, len(i.Cache), len(i.Cache))
	for s := range i.Cache {
		ids = append(ids, s)
	}
	return ids
}

func (i *InmemRepository) FindByAggregateId(id string) []base.EventModel {
	models := i.Cache[id]
	if models != nil {
		return models
	}
	return []base.EventModel{}
}

func (i *InmemRepository) Save(model *base.EventModel) {
	models := i.Cache[model.AggregateId]
	if models == nil {
		models = make([]base.EventModel, 1)
	}
	model.Id = uuid.New().String()
	models = append(models, *model)
	i.Cache[model.AggregateId] = models
}
