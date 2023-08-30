package router

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func InitRouter() {
	// Define port
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	// Initialize Router
	r := gin.Default()

	// Initialize Routes
	InitRoutes(r)

	// Run server
	log.Fatal(r.Run(":" + port))
}
