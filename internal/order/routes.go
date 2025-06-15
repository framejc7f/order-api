package order

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})
	group := r.Group("/orders")
	group.GET("/", GetOrders)
	group.POST("/", CreateOrder)
}
