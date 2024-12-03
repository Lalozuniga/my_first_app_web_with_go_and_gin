package controllers

import (
	"context"
	"net/http"
	"time"

	"github.com/Lalozuniga/my_first_app_with_go_and_gin/config"
	"github.com/Lalozuniga/my_first_app_with_go_and_gin/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var collection *mongo.Collection = config.GetCollection("inventory")

func GetItems(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var items []models.InventoryItem
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var item models.InventoryItem
		cursor.Decode(&item)
		items = append(items, item)
	}

	c.JSON(http.StatusOK, items)
}
func CreateItem(c *gin.Context) {
	var item models.InventoryItem
	if err := c.BindJSON(&item); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := collection.InsertOne(ctx, item)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Producto agregado exitosamente"})
}
