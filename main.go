package main

import (
	"github.com/Lalozuniga/my_first_app_with_go_and_gin/config"
	"github.com/Lalozuniga/my_first_app_with_go_and_gin/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	// Conectar a MongoDB
	config.ConnectDB()

	// Crear el router
	router := gin.Default()

	// Cargar rutas
	routes.SetupRoutes(router)

	// Iniciar el servidor
	router.Run(":8080")
}
