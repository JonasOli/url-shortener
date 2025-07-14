package repositories

import (
	"context"
	"time"
	"url-shortener/internal/models"
	"url-shortener/pkg/database"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type URLRepository struct {
}

func NewURLRepository() *URLRepository {
	return &URLRepository{}
}

func (r *URLRepository) Create(url *models.URL) error {
	collection := database.Client.Database("local").Collection("urls")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	indexModel := mongo.IndexModel{
		Keys:    bson.D{{Key: "short_code", Value: 1}}, // 1 for ascending order
		Options: options.Index().SetUnique(true).SetName("short_code_unique_idx"),
	}

	_, err := collection.Indexes().CreateOne(ctx, indexModel)

	if err != nil {
		return err
	}

	_, err = collection.InsertOne(ctx, url)

	if err != nil {
		return err
	}

	return nil
}
