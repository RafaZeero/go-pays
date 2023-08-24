package config

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func DbConn() (db *sql.DB) {
	dbDriver := "mysql"
	dbUser := "root"
	dbPass := "password"
	dbName := "go_pays_db"
	// Create the data source name (DSN)
	dsn := fmt.Sprintf("%s:%s@tcp(localhost:3306)/%s", dbUser, dbPass, dbName)

	// Open a database connection
	db, err := sql.Open(dbDriver, dsn)
	if err != nil {
		panic(err.Error())
	}
	return db
}
