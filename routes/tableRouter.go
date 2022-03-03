package routes

import (
	c "github.com/bhanupbalusu/go-restaurantms/controllers"
	"github.com/gin-gonic/gin"
)

func TableRoutes(inr *gin.Engine) {
	inr.GET("/tables", c.GetTables())
	inr.GET("/tables/:table-id", c.GetTable())
	inr.POST("/tables", c.CreateTable())
	inr.PATCH("/tables/:table-id", c.UpdateTable())
}
