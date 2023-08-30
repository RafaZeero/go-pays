package handler

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func CreateAccount(ctx *gin.Context) {
	// New account
	na := Account{
		ID:        strconv.Itoa(len(Accounts) + 1),
		CreatedAt: time.Now(),
	}

	if err := ctx.ShouldBindJSON(&na); err != nil {
		return
	}

	// Add new account to list
	Accounts = append(Accounts, na)

	// Response
	ctx.JSON(http.StatusCreated, na)
}
