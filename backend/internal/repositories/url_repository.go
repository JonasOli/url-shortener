package repositories

import (
	"context"
	"time"
	"url-shortener/internal/models"
	"url-shortener/pkg/database"

	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type URLRepository struct {
	redis *redis.Client
}

type OriginalURLOnly struct {
	OriginalURL string `bson:"original_url"`
}

func NewURLRepository(redis *redis.Client) *URLRepository {
	return &URLRepository{redis: redis}
}

func (r *URLRepository) Create(url *models.URL) error {
	collection := database.Client.Database("local").Collection("urls")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	indexModel := mongo.IndexModel{
		Keys:    bson.D{{Key: "short_code", Value: 1}},
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

func (r *URLRepository) FindByShortCode(shortCode string) (string, error) {
	ctx := context.Background()

	cachedURL, err := r.redis.Get(ctx, shortCode).Result()

	if err == nil {
		return cachedURL, nil
	}

	collection := database.Client.Database("local").Collection("urls")

	_, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var result OriginalURLOnly

	err = collection.FindOne(context.TODO(), bson.M{"short_code": shortCode}).Decode(&result)

	if err != nil {
		return "", err
	}

	err = r.redis.Set(ctx, shortCode, result.OriginalURL, 60*time.Minute).Err()

	if err != nil {
		return "", err
	}

	return result.OriginalURL, nil
}
