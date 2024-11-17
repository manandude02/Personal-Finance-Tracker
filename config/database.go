package config

import (
	"context"
	"log"
	"os"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB *mongo.Database

func ConnectDatabase() {
	// Load MongoDB URI from environment variable or fallback
	err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }

	uri := os.Getenv("MONGODB_URI")
	if uri == "" {
		log.Fatal("MONGODB_URI is blank")
	}

	// Set MongoDB client options
	clientOptions := options.Client().ApplyURI(uri)

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}

	// Ping the MongoDB server
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatalf("Failed to ping MongoDB: %v", err)
	}

	// Access the database
	DB = client.Database("personal-finance-tracker") // Replace with your DB name
	log.Println("Connected to MongoDB!")
}
