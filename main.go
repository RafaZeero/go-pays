package main

import (
	"fmt"

	"github.com/RafaZeero/go-pays/config"
	// "github.com/RafaZeero/go-pays/router"
)

type User struct {
	Name    string `json:"name"`
	Balance string `json:"balance"`
}

func main() {
	db := config.DbConn()

	defer db.Close()
	// router.Initialize()
	res, err := db.Query("SELECT * FROM account")

	if err != nil {
		panic(err.Error())
	}

	for res.Next() {
		var user User

		err := res.Scan(&user.Name, &user.Balance)

		if err != nil {
			panic(err.Error())
		}

		fmt.Println(user)
	}

	defer res.Close()
}
