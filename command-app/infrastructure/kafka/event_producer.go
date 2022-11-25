package kafka

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/Shopify/sarama"
	"github.com/offlinebrain/cqrs-bank-example/command-app/base"
	"github.com/offlinebrain/cqrs-bank-example/command-app/infrastructure"
)

type EventProducer struct {
	producer sarama.SyncProducer
}

func NewEventProducer(producer sarama.SyncProducer) *EventProducer {
	return &EventProducer{producer: producer}
}

func (e *EventProducer) Publish(ctx context.Context, topic string, event base.EventModel) error {
	bytes, err := json.Marshal(base.TracedEventModel{
		TraceId:    ctx.Value(infrastructure.TraceIdKey).(string),
		EventModel: event,
	})
	if err != nil {
		return err
	}

	msg := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.StringEncoder(bytes),
	}

	partition, offset, err := e.producer.SendMessage(msg)
	if err != nil {
		return err
	}

	fmt.Printf("Message is stored in topic(%s)/partition(%d)/offset(%d)\n", topic, partition, offset)

	return nil
}
