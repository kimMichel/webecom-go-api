package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/kimMichel/webecom-go-api/controllers"
)

func ConfigRoutes(router *gin.Engine) *gin.Engine {
	main := router.Group("api/v1")
	{
		products := main.Group("products")
		{
			products.GET("/", controllers.GetProducts)
			products.POST("/", controllers.PostProduct)
		}
		users := main.Group("user")
		{
			users.POST("/signup", controllers.Signup)
		}
	}
	return router
}
