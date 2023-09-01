package handler

import (
	"net/http"
	"strconv"

	"github.com/RafaZeero/go-pays/schemas"
	"github.com/gin-gonic/gin"
)

func CreateAccount(ctx *gin.Context) {
	// New account
	var na schemas.Account

	tx, _ := db.Begin()

	if err := ctx.BindJSON(&na); err != nil {
		return
	}

	query := "INSERT INTO accounts (name, balance, createdAt, updatedAt ) VALUES (?, ?, NOW(), NOW())"

	stmt, _ := tx.Prepare(query)

	// Rollback if error
	row, err := stmt.Exec(na.Name, na.Balance)
	if err != nil {
		logger.Errorf("Error creating new account: %s", err.Error())
		tx.Rollback()
		return
	}

	tx.Commit()

	id, _ := row.LastInsertId()

	response := gin.H{
		"success": true,
		"message": "Account created",
		"data": gin.H{
			"id":      id,
			"name":    na.Name,
			"balance": na.Balance,
		},
	}

	// Response
	ctx.JSON(http.StatusCreated, response)
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
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		u.CreatedAt, err = ParseDBDate(createdAtDB)
		if err != nil {
			logger.Errorf("Error handling createdAt: %s", err.Error())
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			continue
		}

		u.UpdatedAt, err = ParseDBDate(updatedAtDB)
		if err != nil {
			logger.Errorf("Error handling updatedAt: %s", err.Error())
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			continue
		}
		us = append(us, u)
	}
	// Response: Get accounts variable
	ctx.JSON(http.StatusOK, us)
}

func GetAccountByID(ctx *gin.Context) {
	// Get Param
	id := ctx.Param("id")
	// Make query
	query := "SELECT id, name, balance, createdAt, updatedAt FROM accounts WHERE id = ?"
	row, err := db.Query(query, id)
	// Validate query
	if err != nil {
		logger.Errorf("Error getting account: %s", err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Create Account
	var a schemas.Account
	var createdAtDB, updatedAtDB []uint8
	// Assign values from DB
	for row.Next() {
		err := row.Scan(&a.ID, &a.Name, &a.Balance, &createdAtDB, &updatedAtDB)
		if err != nil {
			logger.Errorf("Error: %s", err.Error())
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		a.CreatedAt, err = ParseDBDate(createdAtDB)
		if err != nil {
			logger.Errorf("Error handling createdAt: %s", err.Error())
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			continue
		}
		a.UpdatedAt, err = ParseDBDate(updatedAtDB)
		if err != nil {
			logger.Errorf("Error handling updatedAt: %s", err.Error())
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			continue
		}
	}
	// Validate User
	if err := a.Validate(); err != nil {
		logger.Errorf("Error while fetching account: %s", err.Error())
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	// Create response
	response := gin.H{
		"success": true,
		"message": "Account found",
		"data": gin.H{
			"account": a,
		},
	}
	// Sucess 🥳
	ctx.JSON(http.StatusNotFound, response)
}

func UpdateAccount(ctx *gin.Context) {
	// Get Param
	id := ctx.Param("id")
	// Create Account
	var a schemas.Account
	// Validate JSON format
	if err := ctx.ShouldBindJSON(&a); err != nil {
		logger.Errorf("Invalid format for transaction: %s", err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Begin Transaction
	tx, _ := db.Begin()
	// Make query
	query := "UPDATE accounts SET name = ?, updatedAt = NOW() WHERE id = ?"
	// Create statement
	stmt, _ := tx.Prepare(query)
	// Query to update acc
	if _, err := stmt.Exec(a.Name, id); err != nil {
		logger.Errorf("Error updating account: %s", err.Error())
		tx.Rollback()
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Commit changes
	tx.Commit()
	// Create response
	response := gin.H{
		"success": true,
		"message": "Account updated",
		"data": gin.H{
			"account_name": a.Name,
		},
	}
	// Sucess 🥳
	ctx.JSON(http.StatusOK, response)
}

func DeleteAccount(ctx *gin.Context) {
	id := ctx.Param("id")

	row, err := db.Query("SELECT id FROM accounts WHERE id = ?", id)
	if err != nil {
		logger.Errorf("Error fetching data from DB: %s", err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var u schemas.User

	for row.Next() {
		err := row.Scan(&u.ID)
		if err != nil {
			logger.Errorf("Error: %s", err.Error())
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
	}

	if err := u.Validate(); err != nil {
		logger.Errorf("Error: %s", err.Error())
		ctx.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
		return
	}

	tx, _ := db.Begin()
	stmt, err := tx.Prepare("DELETE FROM accounts WHERE id = ?")
	if err != nil {
		logger.Errorf("Error deleting account: %s", err.Error())
		tx.Rollback()
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if _, err := stmt.Exec(id); err != nil {
		logger.Errorf("Error deleting account: %s", err.Error())
		tx.Rollback()
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	tx.Commit()
	response := gin.H{
		"success":    true,
		"message":    "Account deleted",
		"account_id": id,
	}
	ctx.JSON(http.StatusOK, response)
}

func MakeTransaction(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	// Validate Param
	if err != nil {
		logger.Error("Failed to get param 'id'")
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Create transaction
	var t TransactionRequest
	t.ID = id
	// Validate JSON format
	if err := ctx.ShouldBindJSON(&t); err != nil {
		logger.Errorf("Invalid format for transaction: %s", err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Validate amount for 2 decimals only
	t.Amount = Truncate(t.Amount, 0.01)
	// Validate JSON data
	if err := t.Validate(); err != nil {
		logger.Errorf("Error request for transaction: %s", err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Begin Transaction
	tx, _ := db.Begin()
	// Create queries
	queries := struct {
		getAccount    string
		updateBalance string
	}{
		getAccount:    "SELECT id, name, balance FROM accounts WHERE id = ?",
		updateBalance: "UPDATE accounts SET balance = ? WHERE id = ?",
	}
	// First statement
	stmt1, _ := tx.Prepare(queries.getAccount)
	// Query to get account
	row, err := stmt1.Query(t.ID)
	// Validate query
	if err != nil {
		logger.Errorf("Error getting account: %s", err.Error())
		tx.Rollback()
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Create user
	var u schemas.User
	// Assign values from DB
	for row.Next() {
		err := row.Scan(&u.ID, &u.Name, &u.Balance)
		if err != nil {
			logger.Errorf("Error while fetching account: %s", err.Error())
			tx.Rollback()
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}
	// Validate User
	if err := u.Validate(); err != nil {
		logger.Errorf("Error while fetching account: %s", err.Error())
		tx.Rollback()
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	// Transaction process
	var newBalance float64
	if t.TransactionType == "deposit" {
		newBalance = u.Balance + t.Amount
	} else if t.TransactionType == "withdrawl" {
		// TODO: validate newBalance < 0
		newBalance = u.Balance - t.Amount
	}
	// Second statement
	stmt2, _ := tx.Prepare(queries.updateBalance)
	// Validate update new balance
	if _, err := stmt2.Exec(newBalance, t.ID); err != nil {
		logger.Errorf("Error getting account balance: %s", err.Error())
		tx.Rollback()
		return
	}
	// TODO: new statement in a new table to create history of transactions
	// Commit changes
	tx.Commit()
	// Create the response using gin.H
	response := gin.H{
		"success": true,
		"message": "Transaction successful",
		"data": gin.H{
			"account_id":  t.ID,
			"amount":      t.Amount,
			"new_balance": newBalance,
		},
	}
	// Sucess 🥳
	ctx.JSON(http.StatusOK, response)
}
