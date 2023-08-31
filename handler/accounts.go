package handler

import (
	"net/http"

	"github.com/RafaZeero/go-pays/schemas"
	"github.com/gin-gonic/gin"
)

func CreateAccount(ctx *gin.Context) {
	// New account
	var na schemas.Account

	tx, _ := db.Begin()

	if err := ctx.ShouldBindJSON(&na); err != nil {
		return
	}

	query := "INSERT INTO accounts (name, balance, createdAt, updatedAt ) VALUES (?, ?, NOW(), NOW())"

	stmt, _ := tx.Prepare(query)

	// Add new account to list
	_, err := stmt.Exec(na.Name, na.Balance)

	// Rollback if error
	if err != nil {
		logger.Errorf("Error creating new account: %s", err.Error())
		tx.Rollback()
		return
	}

	tx.Commit()

	// Response
	ctx.JSON(http.StatusCreated, gin.H{"response": "Account created successfully."})
}

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

func GetAccountByID(ctx *gin.Context) {
	id := ctx.Param("id")

	// Looks for account with ID equals the params
	// for _, a := range Accounts {
	// 	if a.ID == id {
	// 		ctx.JSON(http.StatusOK, a)
	// 		return
	// 	}
	// }

	logger.Infof("The id is %s", id)
	ctx.JSON(http.StatusNotFound, gin.H{"message": "Account not found"})
}

func UpdateAccount(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "UPDATE account"})
}

func DeleteAccount(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "DELETE account"})
}
