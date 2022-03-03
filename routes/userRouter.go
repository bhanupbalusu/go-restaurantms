package routes

import (
	c "github.com/bhanupbalusu/go-restaurantms/controllers"
	"github.com/gin-gonic/gin"
)

func UserRoutes(inr *gin.Engine) {
	inr.GET("/users", c.GetUsers())
	inr.GET("/users/:user-id", c.GetUser())
	inr.POST("/users/sign-up", c.SignUp())
	inr.POST("/users/login", c.Login())
}
