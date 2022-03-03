package routes

import (
	c "github.com/bhanupbalusu/go-restaurantms/controllers"
	"github.com/gin-gonic/gin"
)

func OrderItemRoutes(inr *gin.Engine) {
	inr.GET("/orderItems", c.GetOrderItems())
	inr.GET("/orderItems/:orderItem-id", c.GetOrderItem())
	inr.GET("/orderItems-order/:orderItem-id", c.GetOrderItemsByOrder())
	inr.POST("/orderItems", c.CreateOrderItem())
	inr.PATCH("/orderItems/:orderItem-id", c.UpdateOrderItem())
}
