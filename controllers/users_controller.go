package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/kimMichel/webecom-go-api/database"
	"github.com/kimMichel/webecom-go-api/models"
	"golang.org/x/crypto/bcrypt"
)

func Signup(c *gin.Context) {
	db := database.GetDatabase()

	// Get the user off req body
	var user models.User

	if c.ShouldBind(&user) != nil {
		c.JSON(400, gin.H{
			"error": "cannot bind JSON",
		})
		return
	}

	// Hash the password
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot hash password",
		})
		return
	}
	// Create the user
	user.Password = string(hash)
	err = db.Create(&user).Error
	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot create user: " + err.Error(),
		})
		return
	}
	// Respond
	c.Status(204)
}
