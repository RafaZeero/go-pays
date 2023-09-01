package config

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

// Create DB and Connect
func InitDB() (*sql.DB, error) {
	logger := GetLogger("DB")
	// Load environment variables from the .env file
	err := godotenv.Load(".env")
	if err != nil {
		logger.Error("Error loading .env file")
	}
	// DB Config
	dbDriver := "mysql"
	// * For Production
	dsn := os.Getenv("DSN")

	if dsn == "" {
		// * For Development
		dbUser := os.Getenv("DB_USER")
		dbPass := os.Getenv("DB_PASSWORD")
		dbName := os.Getenv("DB_NAME")
		dbAdress := os.Getenv("DB_ADDRESS")

		// Create the data source name (DSN)
		dsn = fmt.Sprintf("%s:%s@%s/%s", dbUser, dbPass, dbAdress, dbName)
	}

	// Open a database connection
	db, err := sql.Open(dbDriver, dsn)
	if err != nil {
		logger.Errorf("Failed to open db connection: %v", err.Error())
		return nil, err
	}
	return db, nil
}
