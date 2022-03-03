package main

import (
	"fmt"
	"os"

	"github.com/bhanupbalusu/go-restaurantms/database"
	"github.com/bhanupbalusu/go-restaurantms/middleware"
	"github.com/bhanupbalusu/go-restaurantms/routes"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

var foodCollection *mongo.Collection = database.OpenCollection(database.Client, "food")

func main() {
	fmt.Println("Restautrant Management System")
	port := os.Getenv("PORT")

	if port == "" {
		port = "8000"
	}

	router := gin.New()
	router.Use(gin.Logger())

	routes.UserRoutes(router)
	router.Use(middleware.Authentication())

	routes.FoodRoutes(router)
	routes.MenuRoutes(router)
	routes.InvoiceRoutes(router)
	routes.OrderRoutes(router)
	routes.OrderItemRoutes(router)
	routes.TableRoutes(router)

	router.Run(":" + port)
}
