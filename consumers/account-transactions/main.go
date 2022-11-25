package main

import (
	"account-transactions/handler"
	"account-transactions/pg"
	"context"
	"encoding/json"
	"fmt"
	"github.com/Shopify/sarama"
	"os"
	"os/signal"
	"syscall"
)

func main() {
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

	fmt.Println("Consumer started ")
	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)
	msgCount := 0

	// Get signal for finish
	doneCh := make(chan struct{})
	go func() {
		for {
			select {
			case err := <-consumer.Errors():
				fmt.Println(err)
			case msg := <-consumer.Messages():
				msgCount++
				var model handler.TracedEventModel
				err := json.Unmarshal(msg.Value, &model)
				if err != nil {
					fmt.Printf("err unmarshalling")
				}

				ctx := context.WithValue(context.Background(), handler.TraceIdKey, model.TraceId)
				err = writeHandler.Handle(ctx, model.EventModel)
				if err != nil {
					fmt.Printf("err saving %s \n", err.Error())
				}
				fmt.Printf("Received message Count %d: | Topic(%s) | Message(%s) \n", msgCount, msg.Topic, string(msg.Value))
			case <-sigchan:
				fmt.Println("Interrupt is detected")
				doneCh <- struct{}{}
			}
		}
	}()

	<-doneCh
	fmt.Println("Processed", msgCount, "messages")

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
