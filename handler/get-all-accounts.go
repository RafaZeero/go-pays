package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAccounts(ctx *gin.Context) {
	// Response: Get accounts variable
	ctx.JSON(http.StatusOK, Accounts)
}
