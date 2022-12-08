package controllers

import (
	"github.com/Hugovarellaa/golang-crud/config"
	"github.com/Hugovarellaa/golang-crud/models"
	"github.com/gin-gonic/gin"
)

func CreateCustomer(c *gin.Context) {
	var body struct {
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		Email     string `json:"email"`
		Password  string `json:"password"`
	}

	c.Bind(&body)

	customers := models.Customers{
		FirstName: body.FirstName,
		LastName:  body.LastName,
		Email:     body.Email,
		Password:  body.Password,
	}

	result := config.DB.Create(&customers)

	if result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(200, gin.H{
		"post": customers,
	})
}

func GetAllCustomer(c *gin.Context) {
	customers := []models.Customers{}
	config.DB.Find(&customers)

	c.JSON(200, &customers)
}

func GetCustomerById(c *gin.Context) {
	id := c.Param("id")

	var customers models.Customers
	config.DB.Find(&customers, id)

	c.JSON(200, customers)
}

func UpdateCustomerById(c *gin.Context) {
	id := c.Param("id")

	var body struct {
		Id        uint64 `json:"id"`
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		Email     string `json:"email"`
		Password  string `json:"password"`
	}

	c.Bind(&body)

	var customers models.Customers
	config.DB.First(&customers, id)

	config.DB.Model(&customers).Updates(models.Customers{
		FirstName: body.FirstName,
		LastName:  body.LastName,
		Email:     body.Email,
		Password:  body.Password,
	})

	c.JSON(200, &customers)
}
