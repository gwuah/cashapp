package routes

import (
	"cashapp/core"
	"cashapp/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterPaymentRoutes(e *gin.Engine, s services.Services) {
	e.POST("/payments", func(c *gin.Context) {
		var req core.CreatePaymentRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		}

		response := s.Payments.SendMoney(req)
		if response.Error {
			c.JSON(response.Code, gin.H{
				"message": response.Meta.Message,
			})
			return
		}

		c.JSON(response.Code, response.Meta)
	})

	e.POST("/payments/deposit", func(c *gin.Context) {
	})

	e.POST("/payments/withdraw", func(c *gin.Context) {
	})

}
