package routes

import (
	c "github.com/bhanupbalusu/go-restaurantms/controllers"
	"github.com/gin-gonic/gin"
)

func InvoiceRoutes(inr *gin.Engine) {
	inr.GET("/invoices", c.GetInvoices())
	inr.GET("/invoices/:invoice-id", c.GetInvoice())
	inr.POST("/invoices", c.CreateInvoice())
	inr.PATCH("/invoices/:invoice-id", c.UpdateInvoice())
}
