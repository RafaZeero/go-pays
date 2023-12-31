package router

import (
	"net/http"

	"github.com/RafaZeero/go-pays/handler"
	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.Engine) {
	// Initialize Handler
	handler.InitHandler()

	v1 := r.Group("/api/v1")
	{
		RoutePing(v1)
		InitAccountRoutes(v1)
	}
}

func InitAccountRoutes(r *gin.RouterGroup) {
	r.GET("/accounts", handler.GetAccounts)
	r.GET("/accounts/:id", handler.GetAccountByID)
	r.POST("/accounts", handler.CreateAccount)
	r.POST("/accounts/:id/transaction", handler.MakeTransaction)
	r.PATCH("/accounts/:id", handler.UpdateAccount)
	r.DELETE("/accounts/:id", handler.DeleteAccount)
}

func RoutePing(r *gin.RouterGroup) {
	r.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"response": "pong"})
	})
}
