package routes

import (
    "github.com/gin-gonic/gin"
    "app-web/controllers"
)

func SetupRouter() *gin.Engine {
    r := gin.Default()

    // Ruta para la pantalla de inicio
    r.GET("/", controllers.Home)

    return r
}