package rest

import (
	"github.com/gin-gonic/gin"
	"github.com/offlinebrain/cqrs-bank-example/command-app/base"
	"github.com/offlinebrain/cqrs-bank-example/command-app/domain/account"
	"net/http"
)

type DepositRequest struct {
	Amount uint64 `json:"amount"`
}

func NewAccountDepositHandler(dispatcher base.CommandDispatcher) func(c *gin.Context) {
	return func(c *gin.Context) {
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
			CommandBase: base.CommandBase{
				Id: id,
			},
			Amount: request.Amount,
		}

		if err := dispatcher.Send(command); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"id": command.Id})
	}
}
