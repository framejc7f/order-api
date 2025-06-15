package order

import (
	"net/http"

	"github.com/framejc7f/order-api/internal/database"
	"github.com/framejc7f/order-api/models"
	"github.com/gin-gonic/gin"
)

const getOrdersQuery = `SELECT id, title, amount FROM orders`

func GetOrders(c *gin.Context) {
	rows, err := database.DB.Query(getOrdersQuery)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "DB error"})
		return
	}
	defer rows.Close()

	var orders []models.Order
	for rows.Next() {
		var order models.Order
		if err := rows.Scan(&order.ID, &order.Title, &order.Amount); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Scan error"})
			return
		}
		orders = append(orders, order)
	}

	c.JSON(http.StatusOK, orders)
}

const createOrderQuery = `INSERT INTO orders (title, amount) VALUES ($1, $2) RETURNING id`

func CreateOrder(c *gin.Context) {
	var newOrder models.Order
	if err := c.ShouldBindJSON(&newOrder); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := database.DB.QueryRow(
		createOrderQuery,
		newOrder.Title, newOrder.Amount,
	).Scan(&newOrder.ID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Insert failed"})
		return
	}

	c.JSON(http.StatusCreated, newOrder)
}
