package handler

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// Models
type account struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Balance   float64   `json:"balance"`
	CreatedAt time.Time `json:"createdAt"`
}

var accounts = []account{
	{ID: "1", Name: "User 01", Balance: 10, CreatedAt: time.Now()},
	{ID: "2", Name: "User 02", Balance: 5000, CreatedAt: time.Now()},
	{ID: "3", Name: "User 03", Balance: -10, CreatedAt: time.Now()},
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

func UpdateAccount(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "UPDATE account"})
}

func DeleteAccount(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "DELETE account"})
}
