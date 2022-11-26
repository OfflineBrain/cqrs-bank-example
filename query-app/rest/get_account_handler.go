package rest

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"query-app/infrastructure"
	l "query-app/infrastructure/log"
	"query-app/usecase"
)

func NewGetAccountHandler(uc *usecase.GetAccountUseCase) func(c *gin.Context) {
	return func(c *gin.Context) {
		ctx := context.WithValue(context.Background(), infrastructure.TraceIdKey, "get-"+uuid.New().String())
		log := l.Logger.Logger.WithField(infrastructure.TraceIdKey, ctx.Value(infrastructure.TraceIdKey))
		method := c.Request.Method
		if method != http.MethodGet {
			log.Warnf("endpoint handler called with unsupported method %s - %s", method, c.Request.RequestURI)
			c.Status(http.StatusMethodNotAllowed)
			return
		}

		id := c.Param("id")
		response := uc.Request(*usecase.NewAccountRequest(id))
		if response == nil {
			c.Status(http.StatusNotFound)
		}

		log.Infof("successfully processed, account [%s] closed", id)
		c.JSON(http.StatusOK, response)
	}
}
