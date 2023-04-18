package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/kimMichel/webecom-go-api/database"
	"github.com/kimMichel/webecom-go-api/models"
)

func GetProducts(c *gin.Context) {
	db := database.GetDatabase()

	var products []models.Product
	err := db.Find(&products).Error

	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot find product: " + err.Error(),
		})
		return
	}

	c.JSON(200, products)
}

func PostProduct(c *gin.Context) {
	db := database.GetDatabase()

	var product models.Product

	err := c.ShouldBind(&product)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot bind JSON: " + err.Error(),
		})
		return
	}

	err = db.Create(&product).Error
	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot create product: " + err.Error(),
		})
		return
	}
	c.JSON(200, product)
}
