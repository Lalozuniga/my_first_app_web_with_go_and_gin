package main

import (
    "github.com/gin-gonic/gin"
)

func main() {
    r := gin.Default() // Crea una nueva instancia del servidor con configuraciones predeterminadas
    r.GET("/ping", func(c *gin.Context) {
        c.JSON(200, gin.H{
            "message": "pong",
        })
    })
    r.Run() // Inicia el servidor en el puerto 8080
}
