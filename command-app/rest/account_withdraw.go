package rest

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/offlinebrain/cqrs-bank-example/command-app/base"
	"github.com/offlinebrain/cqrs-bank-example/command-app/domain/account"
	"github.com/offlinebrain/cqrs-bank-example/command-app/infrastructure"
	l "github.com/offlinebrain/cqrs-bank-example/command-app/infrastructure/log"
	"net/http"
)

type WithdrawRequest struct {
	Amount uint64 `json:"amount"`
}

func NewAccountWithdrawHandler(dispatcher base.CommandDispatcher) func(c *gin.Context) {
	return func(c *gin.Context) {
		ctx := context.WithValue(context.Background(), infrastructure.TraceIdKey, "withdraw-"+uuid.New().String())
		log := l.Logger.WithField(infrastructure.TraceIdKey, ctx.Value(infrastructure.TraceIdKey))
		method := c.Request.Method
		if method != http.MethodPost {
			log.Warnf("endpoint handler called with unsupported method %s - %s", method, c.Request.RequestURI)
			c.Status(http.StatusMethodNotAllowed)
			return
		}

		id := c.Param("id")

		var request *WithdrawRequest
		if err := c.ShouldBindJSON(&request); err != nil {
			log.Warnf("request body failed to be parsed %s", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		command := account.WithdrawFundsCommand{
			AccountId: id,
			Amount:    request.Amount,
		}

		if err := dispatcher.Send(ctx, command); err != nil {
			log.Warnf("failed to process command %s", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		log.Infof("successfully processed")
		c.Status(http.StatusAccepted)
	}
}
