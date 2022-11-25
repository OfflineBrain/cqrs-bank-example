package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/offlinebrain/cqrs-bank-example/command-app/domain/account"
	"github.com/offlinebrain/cqrs-bank-example/command-app/infrastructure"
	"github.com/offlinebrain/cqrs-bank-example/command-app/infrastructure/kafka"
	er "github.com/offlinebrain/cqrs-bank-example/command-app/infrastructure/mongo"
	l "github.com/offlinebrain/cqrs-bank-example/command-app/log"
	"github.com/offlinebrain/cqrs-bank-example/command-app/rest"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func main() {
	l.Logger.Logger.SetLevel(logrus.DebugLevel)
	l.Logger.Logger.SetFormatter(&logrus.JSONFormatter{})

	client, _ := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017/?connect=direct"))
	err := client.Ping(context.Background(), readpref.Primary())
	if err != nil {
		panic(err)
	}
	eventRepository := er.NewMongoEventRepository(*client)

	producer, err := kafka.NewSyncProducer([]string{"localhost:9092"})
	if err != nil {
		panic(err)
	}
	publisher := kafka.NewEventProducer(producer)
	eventStore := infrastructure.NewEventStore(eventRepository, publisher, account.Topics)
	accountRepository := account.NewAggregateRepository(eventStore)

	dispatcher := infrastructure.NewCommandDispatcher()

	account.NewCommandHandler(accountRepository).Register(*dispatcher)

	engine := gin.Default()

	api := engine.Group("/api")
	{
		v1 := api.Group("/v1")
		{
			v1.POST("/accounts", rest.NewAccountCreateHandler(dispatcher))
			v1.POST("/accounts/:id/deposit", rest.NewAccountDepositHandler(dispatcher))
			v1.POST("/accounts/:id/withdraw", rest.NewAccountWithdrawHandler(dispatcher))
			v1.DELETE("/accounts/:id", rest.NewCloseAccountHandler(dispatcher))
		}
	}

	_ = engine.Run(":8080")
}
