package handler

import (
	"net/http"

	"github.com/RafaZeero/go-pays/schemas"
	"github.com/gin-gonic/gin"
)

func GetAccounts(ctx *gin.Context) {
	rows, _ := db.Query("select id, name, balance, createdAt, updatedAt from accounts")
	defer rows.Close()

	// All users
	var us []schemas.Account

	for rows.Next() {
		var u schemas.Account
		rows.Scan(&u.ID, &u.Name, &u.Balance, &u.CreatedAt, &u.UpdatedAt)
		us = append(us, u)
	}
	// Response: Get accounts variable
	ctx.JSON(http.StatusOK, us)
}
