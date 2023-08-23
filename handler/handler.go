package handler

import (
	"time"
)

// Models
type Account struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Balance   float64   `json:"balance"`
	CreatedAt time.Time `json:"createdAt"`
}

var Accounts = []Account{
	{ID: "1", Name: "User 01", Balance: 10, CreatedAt: time.Now()},
	{ID: "2", Name: "User 02", Balance: 5000, CreatedAt: time.Now()},
	{ID: "3", Name: "User 03", Balance: -10, CreatedAt: time.Now()},
}
