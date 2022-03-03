package routes

import (
	c "github.com/bhanupbalusu/go-restaurantms/controllers"
	"github.com/gin-gonic/gin"
)

func MenuRoutes(inr *gin.Engine) {
	inr.GET("/menus", c.GetMenus())
	inr.GET("/menus/:menu-id", c.GetMenu())
	inr.POST("/menus", c.CreateMenu())
	inr.PATCH("/menus/:menu-id", c.UpdateMenu())
}
