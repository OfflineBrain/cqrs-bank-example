package rest

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/offlinebrain/cqrs-bank-example/command-app/base"
	"github.com/offlinebrain/cqrs-bank-example/command-app/domain/account"
	"github.com/offlinebrain/cqrs-bank-example/command-app/infrastructure"
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
		method := c.Request.Method
		if method != http.MethodPost {
			c.Status(http.StatusMethodNotAllowed)
			return
		}

		var request *OpenAccountRequest
		if err := c.ShouldBindJSON(&request); err != nil {
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
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusCreated, gin.H{"id": command.AccountId})
	}
}
