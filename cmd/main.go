package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"github.com/framejc7f/order-api/internal/order"
)

func main() {
	r := gin.Default()

	order.RegisterRoutes(r)

	log.Println("Server running on http://localhost:8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
