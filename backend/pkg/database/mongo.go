package database

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client

func InitMongo(uri string) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	clientOptions := options.Client().ApplyURI(uri)

	var err error
	Client, err = mongo.Connect(ctx, clientOptions)

	if err != nil {
		log.Fatalf("MongoDB connection error: %v", err)
	}

	// Check the connection
	err = Client.Ping(ctx, nil)
	if err != nil {
		log.Fatalf("MongoDB ping error: %v", err)
	}

	log.Println("✅ Connected to MongoDB")
}
