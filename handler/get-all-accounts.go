package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAccounts(ctx *gin.Context) {
	rows, _ := db.Query("select id, name, balance from accounts")
	defer rows.Close()

	// All users
	var us []User

	for rows.Next() {
		var u User
		rows.Scan(&u.ID, &u.Name, &u.Balance)
		us = append(us, u)
	}
	// Response: Get accounts variable
	ctx.JSON(http.StatusOK, us)
}
