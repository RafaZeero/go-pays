package handler

import (
	"database/sql"
	"time"

	"github.com/RafaZeero/go-pays/config"
)

var (
	logger *config.Logger
	db     *sql.DB
)

func InitHandler() {
	logger = config.GetLogger("handler")
	db = config.GetDB()
}

func ParseDBDate(dbDate []uint8) (time.Time, error) {
	dateString := string(dbDate)
	parsedTime, err := time.Parse("2006-01-02 15:04:05", dateString)
	if err != nil {
		return time.Time{}, err
	}
	return parsedTime, nil
}
