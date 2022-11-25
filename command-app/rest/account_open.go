package rest

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/offlinebrain/cqrs-bank-example/command-app/base"
	"github.com/offlinebrain/cqrs-bank-example/command-app/domain/account"
	"github.com/offlinebrain/cqrs-bank-example/command-app/infrastructure"
	l "github.com/offlinebrain/cqrs-bank-example/command-app/log"
	"net/http"
)

type OpenAccountRequest struct {
	HolderName     string `json:"holderName" binding:"required"`
	AccountType    string `json:"accountType" binding:"required"`
	OpeningBalance uint64 `json:"openingBalance" binding:"required"`
}

func NewAccountCreateHandler(dispatcher base.CommandDispatcher) func(c *gin.Context) {
	return func(c *gin.Context) {
		ctx := context.WithValue(context.Background(), infrastructure.TraceIdKey, "open-"+uuid.New().String())
		log := l.Logger.WithField(infrastructure.TraceIdKey, ctx.Value(infrastructure.TraceIdKey))
		method := c.Request.Method
		if method != http.MethodPost {
			log.Warnf("endpoint handler called with unsupported method %s - %s", method, c.Request.RequestURI)
			c.Status(http.StatusMethodNotAllowed)
			return
		}

		var request *OpenAccountRequest
		if err := c.ShouldBindJSON(&request); err != nil {
			log.Warnf("request body failed to be parsed %s", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		command := account.OpenAccountCommand{
			AccountId:      uuid.New().String(),
			HolderName:     request.HolderName,
			AccountType:    request.AccountType,
			OpeningBalance: request.OpeningBalance,
		}

		if err := dispatcher.Send(ctx, command); err != nil {
			log.Warnf("failed to process command %s", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		log.Infof("successfully processed, new created account id [%s]", command.AccountId)
		c.JSON(http.StatusCreated, gin.H{"id": command.AccountId})
	}
}
