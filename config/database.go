package config

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoClient *mongo.Client

func ConnectDB() {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017") // Cambia si tu configuración es diferente
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Conectar al cliente MongoDB
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatalf("Error conectando a MongoDB: %v", err)
	}

	// Probar la conexión
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatalf("No se pudo conectar a MongoDB: %v", err)
	}

	fmt.Println("Conexión exitosa a MongoDB")
	MongoClient = client
}

func GetCollection(collectionName string) *mongo.Collection {
	// Verifica que el cliente esté inicializado
	if MongoClient == nil {
		log.Fatal("El cliente de MongoDB no está inicializado. Asegúrate de llamar a ConnectDB() antes de usar GetCollection.")
	}
	return MongoClient.Database("inventory_db").Collection(collectionName) // Cambia "inventory_db" por el nombre de tu base de datos
}
