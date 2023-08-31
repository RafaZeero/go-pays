package schemas

import (
	"fmt"
	"time"
)

type Account struct {
	User
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type User struct {
	ID      int64   `json:"id"`
	Name    string  `json:"name"`
	Balance float64 `json:"balance"`
}

// v=Value, t=Type
func (u *User) Validate() error {
	if u.ID <= 0 || u.Name == "" {
		return fmt.Errorf("user not found")
	}
	return nil
}
