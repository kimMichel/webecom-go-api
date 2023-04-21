package controllers

import (
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
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

func Signin(c *gin.Context) {
	db := database.GetDatabase()
	// Get the user email and pass off req body
	var user models.User

	if c.ShouldBind(&user) != nil {
		c.JSON(400, gin.H{
			"error": "cannot bind JSON",
		})
		return
	}
	// Look up requested user
	db.First(&user, "email = ?", user.Email)
	if user.ID == 0 {
		c.JSON(400, gin.H{
			"error": "invalid email or password",
		})
		return
	}

	// Compare sent in pass with saved user pass hash
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(user.Password))
	if err != nil {
		c.JSON(400, gin.H{
			"error": "invalid email or password",
		})
		return
	}
	// Generate jwt token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 2).Unix(),
	})

	tokenString, err := token.SignedString(os.Getenv("SECRET"))
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Failed to create token",
		})
		return
	}

	// Send it back
	c.JSON(http.StatusOK, gin.H{
		"token": tokenString,
	})
}
