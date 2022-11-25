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

func NewCloseAccountHandler(dispatcher base.CommandDispatcher) func(c *gin.Context) {
	return func(c *gin.Context) {
		ctx := context.WithValue(context.Background(), infrastructure.TraceIdKey, "close-"+uuid.New().String())
		method := c.Request.Method
		if method != http.MethodDelete {
			c.Status(http.StatusMethodNotAllowed)
			return
		}

		id := c.Param("id")
		command := account.CloseAccountCommand{
			AccountId: id,
		}

		if err := dispatcher.Send(ctx, command); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.Status(http.StatusNoContent)
	}
}
