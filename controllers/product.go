package controllers

import (
	"github.com/gin-gonic/gin"
	"go-atomic/main/helpers"
	"go-atomic/main/models"
	"go-atomic/main/repositories"
)

func GetAllProduct(c *gin.Context) {
	var products []models.Product
	err := repositories.GetAllProduct(&products)

	if err != nil {
		helpers.RespondJSON(c, 404, products)
	} else {
		helpers.RespondJSON(c, 200, products)
	}
}

func GetProductById(c *gin.Context) {
	product := helpers.GetProduct()
	helpers.RespondJSON(c, 200, product)
}

func UpdateProduct(c *gin.Context) {
	helpers.DecQuantity()
	product := helpers.GetProduct()
	helpers.RespondJSON(c, 200, product)
}
