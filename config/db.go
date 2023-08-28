package config

import (
	"database/sql"
	"fmt"
)

// Create DB and Connect
func InitDB() (*sql.DB, error) {
	logger := GetLogger("DB")

	// DB Config
	dbDriver := "mysql"
	dbUser := "root"
	dbPass := "password"
	dbName := "go_pays_db"
	dbAdress := "tcp(localhost:3306)"

	// Create the data source name (DSN)
	dsn := fmt.Sprintf("%s:%s@%s/%s", dbUser, dbPass, dbAdress, dbName)

	// Open a database connection
	db, err := sql.Open(dbDriver, dsn)
	if err != nil {
		logger.Errorf("Failed to open db connection: %v", err.Error())
		return nil, err
	}
	return db, nil
}
