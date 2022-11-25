package kafka

import (
	"context"
	"encoding/json"
	"github.com/Shopify/sarama"
	"github.com/offlinebrain/cqrs-bank-example/command-app/base"
	"github.com/offlinebrain/cqrs-bank-example/command-app/infrastructure"
	l "github.com/offlinebrain/cqrs-bank-example/command-app/log"
	"github.com/sirupsen/logrus"
)

type EventProducer struct {
	producer sarama.SyncProducer
}

func NewEventProducer(producer sarama.SyncProducer) *EventProducer {
	return &EventProducer{producer: producer}
}

func (e *EventProducer) Publish(ctx context.Context, topic string, event base.EventModel) error {
	log := l.Logger.WithFields(logrus.Fields{
		infrastructure.TraceIdKey:       ctx.Value(infrastructure.TraceIdKey),
		infrastructure.AggregateIdKey:   event.AggregateId,
		infrastructure.AggregateTypeKey: event.AggregateType,
	})
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

	log.Debugf("Message [%s] is stored in topic(%s)/partition(%d)/offset(%d)\n", event.Id, topic, partition, offset)

	return nil
}
