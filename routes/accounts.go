package routes

import (
	"cashapp/core"
	"cashapp/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterAccountsRoutes(e *gin.Engine, s services.Services) {
	e.POST("/accounts", func(c *gin.Context) {

		var req core.CreateAccountRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		}

		response := s.Accounts.CreateAccount(req)

		if response.Error {
			c.JSON(response.Code, gin.H{
				"message": response.Meta.Message,
			})
			return
		}

		c.JSON(response.Code, response.Meta)
	})

}
