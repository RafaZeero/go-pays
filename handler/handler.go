package handler

import (
	"database/sql"

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
