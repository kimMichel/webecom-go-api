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
