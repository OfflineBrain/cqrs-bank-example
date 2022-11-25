package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"query-app/config"
	"query-app/log"
	"query-app/pg"
	"query-app/rest"
	"query-app/usecase"
)

func main() {
	cfg, err := config.LoadConfig(".")
	if err != nil {
		return
	}

	log.SetServiceName(cfg.ServiceName)
	log.Logger.Logger.SetLevel(logrus.DebugLevel)
	log.Logger.Logger.SetFormatter(&logrus.JSONFormatter{})

	connString := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.PgHost,
		cfg.PgPort,
		cfg.PgUser,
		cfg.PgPassword,
		cfg.PgDatabase,
	)
	connection, err := pg.NewPgConnection(connString)
	if err != nil {
		panic(err)
	}

	accountRepository := pg.NewAccountRepository(connection)

	engine := gin.Default()

	api := engine.Group("/api")
	{
		v1 := api.Group("/v1")
		{
			v1.GET("/accounts/:id", rest.NewGetAccountHandler(usecase.NewGetAccountUseCase(accountRepository)))
		}
	}

	addr := fmt.Sprintf(":%d", cfg.ServerPort)
	_ = engine.Run(addr)
}
