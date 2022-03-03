package routes

import (
	c "github.com/bhanupbalusu/go-restaurantms/controllers"
	"github.com/gin-gonic/gin"
)

func FoodRoutes(inr *gin.Engine) {
	inr.GET("/foods", c.GetFoods())
	inr.GET("/foods/:food_id", c.GetFood())
	inr.POST("/foods", c.CreateFood())
	inr.PATCH("/foods/:food_id", c.UpdateFood())
}
