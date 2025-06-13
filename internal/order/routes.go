package order

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	group := r.Group("/orders")
	group.GET("/", GetOrders)
	group.POST("/", CreateOrder)
}
