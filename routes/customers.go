package routes

import (
	"github.com/Hugovarellaa/golang-crud/controllers"
	"github.com/gin-gonic/gin"
)

func CustomersRoutes(routes *gin.Engine) {
	routes.POST("/customers", controllers.CreateCustomer)
	routes.GET("/customers", controllers.GetAllCustomer)
	routes.GET("/customers/:id", controllers.GetCustomerById)
	routes.PUT("/customers/:id", controllers.UpdateCustomerById)
}
