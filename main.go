package main

import (
	"github.com/RafaZeero/go-pays/config"
	"github.com/RafaZeero/go-pays/router"
)

var (
	logger *config.Logger
)

func main() {
	logger = config.GetLogger("main")

	// Initialize Configs
	err := config.Init()

	if err != nil {
		logger.Errorf("Config init error: %v", err.Error())
		return
	}

	// Initialize Router
	router.InitRouter()
}

// #region insert
// func insert(db *sql.DB) {
// 	// Init transaction
// 	tx, _ := db.Begin()
// 	// Statement
// 	stmt, _ := tx.Prepare("insert into accounts (name, balance) values (?,?)")
// 	stmt.Exec("RafaZeero3", 1456)
// 	stmt.Exec("Luffy3", 2789)
// 	_, err := stmt.Exec("goku2", 34770)
// 	// Rollback if an error occur
// 	if err != nil {
// 		tx.Rollback()
// 		log.Fatal(err)
// 	}
// 	// Commit transaction
// 	tx.Commit()
// }
// #endregion insert

// #region select
// func selectFromDB(db *sql.DB) {
// 	rows, _ := db.Query("select id, name, balance from accounts where id > ?", 2004)
// 	defer rows.Close()
// 	for rows.Next() {
// 		var u User
// 		rows.Scan(&u.ID, &u.Name, &u.Balance)
// 		fmt.Println(u)
// 	}
// }

// #endregion select
