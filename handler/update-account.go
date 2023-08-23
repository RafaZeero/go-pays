package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func UpdateAccount(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "UPDATE account"})
}
