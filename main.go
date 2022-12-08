package main

import (
	"github.com/Hugovarellaa/golang-crud/config"
	"github.com/Hugovarellaa/golang-crud/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	route := gin.Default()
	config.ConnectToDB()
	routes.CustomersRoutes(route)
	route.Run(":8080") // listen and serve on 0.0.0.0:8080
}
