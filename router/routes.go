package router

import (
	"github.com/RafaZeero/go-pays/handler"
	"github.com/gin-gonic/gin"
)

func InitializeRoutes(r *gin.Engine) {
	v1 := r.Group("/api/v1")
	{
		v1.GET("/accounts", handler.GetAccounts)
		v1.GET("/accounts/:id", handler.GetAccountByID)
		v1.POST("/accounts", handler.CreateAccount)
		v1.PUT("/accounts", handler.UpdateAccount)
		v1.DELETE("/accounts", handler.DeleteAccount)
	}
}
