package controllers

import (
    "github.com/gin-gonic/gin"
    "net/http"
)

// Controlador para la pantalla de inicio
func Home(c *gin.Context) {
    c.HTML(http.StatusOK, "index.html", gin.H{
        "title": "Pantalla de Inicio",
    })
}