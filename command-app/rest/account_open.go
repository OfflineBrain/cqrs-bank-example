package rest

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/offlinebrain/cqrs-bank-example/command-app/base"
	"github.com/offlinebrain/cqrs-bank-example/command-app/domain/account"
	"net/http"
)

type OpenAccountRequest struct {
	HolderName     string `json:"holderName" binding:"required"`
	AccountType    string `json:"accountType" binding:"required"`
	OpeningBalance uint64 `json:"openingBalance" binding:"required"`
}

func NewAccountCreateHandler(dispatcher base.CommandDispatcher) func(c *gin.Context) {
	return func(c *gin.Context) {
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
			CommandBase: base.CommandBase{
				Id: uuid.New().String(),
			},
			HolderName:     request.HolderName,
			AccountType:    request.AccountType,
			OpeningBalance: request.OpeningBalance,
		}

		if err := dispatcher.Send(command); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"id": command.Id})
	}
}
