package schemas

import "time"

type Account struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	Balance   float64   `json:"balance"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type User struct {
	ID      int64  `json:"id"`
	Name    string `json:"name"`
	Balance string `json:"balance"`
}
