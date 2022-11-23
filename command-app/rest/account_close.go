package rest

import (
	"github.com/gin-gonic/gin"
	"github.com/offlinebrain/cqrs-bank-example/command-app/base"
	"github.com/offlinebrain/cqrs-bank-example/command-app/domain/account"
	"net/http"
)

func NewCloseAccountHandler(dispatcher base.CommandDispatcher) func(c *gin.Context) {
	return func(c *gin.Context) {
		method := c.Request.Method
		if method != http.MethodDelete {
			c.Status(http.StatusMethodNotAllowed)
			return
		}

		id := c.Param("id")
		command := account.CloseAccountCommand{
			CommandBase: base.CommandBase{
				Id: id,
			},
		}

		if err := dispatcher.Send(command); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.Status(http.StatusNoContent)
	}
}
