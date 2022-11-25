package main

import (
	"account-transactions/handler"
	"account-transactions/log"
	"account-transactions/pg"
	"context"
	"encoding/json"
	"github.com/Shopify/sarama"
	"github.com/sirupsen/logrus"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	log.Logger.Logger.SetLevel(logrus.DebugLevel)
	log.Logger.Logger.SetFormatter(&logrus.JSONFormatter{})

	topic := "account-transactions"
	worker, err := connectConsumer([]string{"localhost:9092"})
	if err != nil {
		panic(err)
	}

	connection, err := pg.NewPgConnection("localhost", 5432, "root", "root", "bank")
	if err != nil {
		panic(err)
	}

	repository := pg.NewAccountRepository(connection)
	writeHandler := handler.NewDbWriteHandler(repository)
	consumer, err := worker.ConsumePartition(topic, 0, sarama.OffsetNewest)
	if err != nil {
		panic(err)
	}

	log.Logger.Infof("Consumer started ")
	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)
	msgCount := 0

	// Get signal for finish
	doneCh := make(chan struct{})
	go func() {
		for {
			select {
			case err := <-consumer.Errors():
				log.Logger.Error(err)
			case msg := <-consumer.Messages():
				msgCount++
				var model handler.TracedEventModel
				err := json.Unmarshal(msg.Value, &model)
				if err != nil {
					log.Logger.Errorf("err unmarshalling %s", err)
				}
				l := log.Logger.WithField(handler.TraceIdKey, model.TraceId)
				ctx := context.WithValue(context.Background(), handler.TraceIdKey, model.TraceId)
				err = writeHandler.Handle(ctx, model.EventModel)
				if err != nil {
					l.Errorf("err saving %s \n", err.Error())
				}
				l.Infof("Received message Count %d: | Topic(%s) | Message(%s) \n", msgCount, msg.Topic, string(msg.Value))
			case <-sigchan:
				log.Logger.Infof("Interrupt is detected")
				doneCh <- struct{}{}
			}
		}
	}()

	<-doneCh
	log.Logger.Info("Processed", msgCount, "messages")

	if err := worker.Close(); err != nil {
		panic(err)
	}
}

func connectConsumer(brokersUrl []string) (sarama.Consumer, error) {
	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true
	config.Net.Proxy.Enable = false
	config.Net.TLS.Enable = false

	// Create new consumer
	conn, err := sarama.NewConsumer(brokersUrl, config)
	if err != nil {
		return nil, err
	}

	return conn, nil
}
