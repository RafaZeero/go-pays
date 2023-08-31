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

	if err := ctx.ShouldBindJSON(&na); err != nil {
		return
	}

	query := "INSERT INTO accounts (name, balance, createdAt, updatedAt ) VALUES (?, ?, NOW(), NOW())"

	stmt, _ := tx.Prepare(query)

	// Rollback if error
	if _, err := stmt.Exec(na.Name, na.Balance); err != nil {
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

func MakeTransaction(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		logger.Error("Failed to get param 'id'")
		return
	}
	// Create transaction
	var t TransactionRequest
	t.ID = id
	if err := ctx.ShouldBindJSON(&t); err != nil {
		return
	}

	if err := t.Validate(); err != nil {
		logger.Errorf("Error request for transaction: %s", err.Error())
		return
	}

	// Begin Transaction
	tx, _ := db.Begin()

	queries := struct {
		getBalance    string
		updateBalance string
	}{
		getBalance:    "SELECT id, name, balance FROM accounts WHERE id = ?",
		updateBalance: "UPDATE accounts SET balance = ? WHERE id = ?",
	}

	// First statement
	stmt1, _ := tx.Prepare(queries.getBalance)

	row, err := stmt1.Query(t.ID)

	if err != nil {
		logger.Errorf("Error getting account balance: %s", err.Error())
		tx.Rollback()
		return
	}

	var u schemas.User

	for row.Next() {
		err := row.Scan(&u.ID, &u.Name, &u.Balance)
		if err != nil {
			logger.Errorf("Error while fetching account: %s", err.Error())
			tx.Rollback()
			return
		}
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

	if _, err := stmt2.Exec(newBalance, t.ID); err != nil {
		logger.Errorf("Error getting account balance: %s", err.Error())
		tx.Rollback()
		return
	}

	// TODO: new statement in a new table to create history of transactions

	tx.Commit()

	// Create the response using gin.H
	response := gin.H{
		"success": true,
		"message": "Deposit successful",
		"data": gin.H{
			"account_id":  t.ID,
			"amount":      t.Amount,
			"new_balance": newBalance,
		},
	}

	ctx.JSON(http.StatusOK, response)
}

func MakeWithdrawal(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"response": "Account created successfully."})
}
