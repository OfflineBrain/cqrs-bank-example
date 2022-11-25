package mongo

import (
	"context"
	"github.com/google/uuid"
	"github.com/offlinebrain/cqrs-bank-example/command-app/base"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoEventRepository struct {
	collection *mongo.Collection
}

func (m *MongoEventRepository) GetAggregateIds(aggregateType string) []string {
	filter := bson.D{{"aggregate_type", aggregateType}}
	distinct, err := m.collection.Distinct(context.Background(), "aggregate_id", filter)
	if err != nil {
		return []string{}
	}

	ids := make([]string, len(distinct))
	for i, id := range distinct {
		ids[i] = id.(string)
	}

	return ids
}

func (m *MongoEventRepository) FindByAggregateId(id string) []base.EventModel {
	filter := bson.D{{"aggregate_id", id}}

	find, err := m.collection.Find(context.Background(), filter)
	if err != nil {
		return nil
	}

	var events []base.EventModel
	err = find.All(context.Background(), &events)
	if err != nil {
		return nil
	}

	return events
}

func (m *MongoEventRepository) Save(model *base.EventModel) {
	model.Id = uuid.New().String()
	result, err := m.collection.InsertOne(context.Background(), model)
	if err != nil {
		return
	}
	model.Id = result.InsertedID.(string)
}

func NewMongoEventRepository(client mongo.Client) *MongoEventRepository {
	collection := client.Database("event_sourcing").Collection("event")
	return &MongoEventRepository{collection: collection}
}
