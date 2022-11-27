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

func NewCloseAccountHandler(dispatcher base.CommandDispatcher) func(c *gin.Context) {
	return func(c *gin.Context) {
		ctx := context.WithValue(context.Background(), infrastructure.TraceIdKey, "close-"+uuid.New().String())
		log := l.Logger.WithField(infrastructure.TraceIdKey, ctx.Value(infrastructure.TraceIdKey))
		method := c.Request.Method
		if method != http.MethodDelete {
			log.Warnf("endpoint handler called with unsupported method %s - %s", method, c.Request.RequestURI)
			c.Status(http.StatusMethodNotAllowed)
			return
		}

		id := c.Param("id")
		command := account.CloseAccountCommand{
			AccountId: id,
		}
		if err := dispatcher.Send(ctx, command); err != nil {
			log.Warnf("failed to process command %s", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		log.Infof("successfully processed, account [%s] closed", id)
		c.Status(http.StatusNoContent)
	}
}
