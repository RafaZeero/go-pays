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
		var createdAtDB, updatedAtDB []uint8
		err := rows.Scan(&u.ID, &u.Name, &u.Balance, &createdAtDB, &updatedAtDB)
		if err != nil {
			logger.Errorf("Error while fetching accounts: %s", err.Error())
			return
		}
		u.CreatedAt, err = ParseDBDate(createdAtDB)
		if err != nil {
			// Handle the error
			logger.Errorf("Error handling createdAt: %s", err.Error())
			continue
		}

		u.UpdatedAt, err = ParseDBDate(updatedAtDB)
		if err != nil {
			// Handle the error
			logger.Errorf("Error handling updatedAt: %s", err.Error())
			continue
		}
		us = append(us, u)
	}
	// Response: Get accounts variable
	ctx.JSON(http.StatusOK, us)
}
