package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/offlinebrain/cqrs-bank-example/command-app/config"
	"github.com/offlinebrain/cqrs-bank-example/command-app/domain/account"
	"github.com/offlinebrain/cqrs-bank-example/command-app/infrastructure"
	"github.com/offlinebrain/cqrs-bank-example/command-app/infrastructure/kafka"
	l "github.com/offlinebrain/cqrs-bank-example/command-app/infrastructure/log"
	"github.com/offlinebrain/cqrs-bank-example/command-app/infrastructure/metrics"
	er "github.com/offlinebrain/cqrs-bank-example/command-app/infrastructure/mongo"
	"github.com/offlinebrain/cqrs-bank-example/command-app/rest"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func main() {
	cfg, err := config.LoadConfig(".")
	if err != nil {
		panic(err)
	}

	metrics.RegisterMetrics()

	l.SetServiceName(cfg.ServiceName)
	l.Logger.Logger.SetLevel(logrus.DebugLevel)
	l.Logger.Logger.SetFormatter(&logrus.JSONFormatter{})

	mongoUri := fmt.Sprintf("mongodb://%s:%d/?connect=direct", cfg.MongoDbHost, cfg.MongoDbPort)
	client, _ := mongo.Connect(context.TODO(), options.Client().ApplyURI(mongoUri))
	err = client.Ping(context.Background(), readpref.Primary())
	if err != nil {
		panic(err)
	}
	eventRepository := infrastructure.NewPromMongoEventRepository(er.NewMongoEventRepository(*client))

	kafkaUri := fmt.Sprintf("%s:%d", cfg.KafkaHost, cfg.KafkaPort)
	producer, err := kafka.NewSyncProducer([]string{kafkaUri})
	if err != nil {
		panic(err)
	}
	publisher := kafka.NewEventProducer(producer)
	eventStore := infrastructure.NewEventStore(eventRepository, publisher, account.Topics(cfg.KafkaAccountTxTopic))
	accountRepository := account.NewAggregateRepository(eventStore)

	dispatcher := infrastructure.NewCommandDispatcher()

	account.NewCommandHandler(accountRepository, eventStore).Register(*dispatcher)

	engine := gin.Default()

	api := engine.Group("/api")
	api.Use(metrics.CommonMiddleware)
	{
		v1 := api.Group("/v1")
		{
			v1.POST("/accounts", rest.NewAccountCreateHandler(dispatcher))
			v1.POST("/accounts/:id/deposit", rest.NewAccountDepositHandler(dispatcher))
			v1.POST("/accounts/:id/withdraw", rest.NewAccountWithdrawHandler(dispatcher))
			v1.DELETE("/accounts/:id", rest.NewCloseAccountHandler(dispatcher))
		}
		evensApi := api.Group("/events")
		{
			evensApi.POST("accounts/replay", rest.NewReplayAccountHandler(dispatcher))
		}
	}

	engine.GET("/metrics", metrics.PrometheusHandler())

	addr := fmt.Sprintf(":%d", cfg.ServerPort)
	_ = engine.Run(addr)
}
