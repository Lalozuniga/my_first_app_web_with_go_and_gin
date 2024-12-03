package routes

import (
	"github.com/Lalozuniga/my_first_app_with_go_and_gin/controllers"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	router.Static("/public", "./public")
	router.LoadHTMLGlob("views/*")

	router.GET("/", func(c *gin.Context) {
		c.HTML(200, "products.html", nil)
	})

	router.GET("/api/products", controllers.GetItems)
	router.POST("/api/products", controllers.CreateItem)
}
