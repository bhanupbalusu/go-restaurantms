package routes

import (
	c "github.com/bhanupbalusu/go-restaurantms/controllers"
	"github.com/gin-gonic/gin"
)

func OrderRoutes(inr *gin.Engine) {
	inr.GET("/orders", c.GetOrders())
	inr.GET("/orders/:order-id", c.GetOrder())
	inr.POST("/orders", c.CreateOrder())
	inr.PATCH("/orders/:order-id", c.UpdateOrder())
}
