package router

import (
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

// Models
type account struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Balance   float64   `json:"balance"`
	CreatedAt time.Time `json:"createdAt"`
}

var accounts = []account{
	{ID: "1", Name: "User 01", Balance: 10, CreatedAt: time.Now()},
	{ID: "2", Name: "User 02", Balance: 5000, CreatedAt: time.Now()},
	{ID: "3", Name: "User 03", Balance: -10, CreatedAt: time.Now()},
}

func Initialize() {
	// Define port
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	// Initialize Router
	r := gin.Default()

	// Initialize Routes
	InitializeRoutes(r)

	// Run server
	r.Run(":" + port)
}
