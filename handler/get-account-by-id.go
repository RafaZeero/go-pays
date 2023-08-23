package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAccountByID(ctx *gin.Context) {
	id := ctx.Param("id")

	// Looks for account with ID equals the params
	for _, a := range Accounts {
		if a.ID == id {
			ctx.JSON(http.StatusOK, a)
			return
		}
	}
	// Response
	ctx.JSON(http.StatusNotFound, gin.H{"message": "Account not found"})
}
