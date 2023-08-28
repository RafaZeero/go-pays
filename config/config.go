package config

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var (
	db     *sql.DB
	logger *Logger
)

func Init() error {
	var err error

	//Initialize DB
	db, err = InitDB()

	if err != nil {
		return fmt.Errorf("Error init DB: %v", err)
	}

	return nil
}

func GetDB() *sql.DB {
	return db
}

func GetLogger(p string) *Logger {
	logger = NewLogger(p)
	return logger
}
