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
