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

type DepositRequest struct {
	Amount uint64 `json:"amount"`
}

func NewAccountDepositHandler(dispatcher base.CommandDispatcher) func(c *gin.Context) {
	return func(c *gin.Context) {
		ctx := context.WithValue(context.Background(), infrastructure.TraceIdKey, "deposit-"+uuid.New().String())
		method := c.Request.Method
		if method != http.MethodPost {
			c.Status(http.StatusMethodNotAllowed)
			return
		}

		id := c.Param("id")

		var request *DepositRequest
		if err := c.ShouldBindJSON(&request); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		command := account.DepositFundsCommand{
			AccountId: id,
			Amount:    request.Amount,
		}

		if err := dispatcher.Send(ctx, command); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.Status(http.StatusAccepted)
	}
}
