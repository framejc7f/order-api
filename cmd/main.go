package main

import (
	"log"

	"github.com/framejc7f/order-api/internal/database"
	"github.com/framejc7f/order-api/internal/order"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func main() {
	err := database.Init()
	if err != nil {
		log.Fatalf("Failed to connect to DB: %v", err)
	}

	r := gin.Default()
	order.RegisterRoutes(r)

	log.Println("Server running on http://localhost:8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
