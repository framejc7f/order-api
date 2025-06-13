package order

import (
	"net/http"

	"github.com/framejc7f/order-api/models"
	"github.com/gin-gonic/gin"
)

var orders = []models.Order{
	{ID: 1, Title: "Тестовый заказ", Amount: 2},
}

func GetOrders(c *gin.Context) {
	c.JSON(http.StatusOK, orders)
}

func CreateOrder(c *gin.Context) {
	var newOrder models.Order
	if err := c.ShouldBindJSON(&newOrder); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	newOrder.ID = len(orders) + 1
	orders = append(orders, newOrder)
	c.JSON(http.StatusCreated, newOrder)
}
