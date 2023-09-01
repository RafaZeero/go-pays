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
