package router

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func InitializeRoutes(r *gin.Engine) {
	v1 := r.Group("/api/v1")
	{
		v1.GET("/accounts", GetAccounts)
		v1.GET("/accounts/:id", GetAccountByID)
		v1.POST("/accounts", CreateAccount)
	}
}

func GetAccounts(ctx *gin.Context) {
	// Response: Get accounts variable
	ctx.JSON(http.StatusOK, accounts)
}

func CreateAccount(ctx *gin.Context) {
	na := account{
		ID:        strconv.Itoa(len(accounts) + 1),
		CreatedAt: time.Now(),
	}

	if err := ctx.ShouldBindJSON(&na); err != nil {
		return
	}

	// Add new account to list
	accounts = append(accounts, na)

	// Response
	ctx.JSON(http.StatusCreated, na)
}

func GetAccountByID(ctx *gin.Context) {
	id := ctx.Param("id")

	// Looks for account with ID equals the params
	for _, a := range accounts {
		if a.ID == id {
			ctx.JSON(http.StatusOK, a)
			return
		}
	}
	// Response
	ctx.JSON(http.StatusNotFound, gin.H{"message": "Account not found"})
}
