package repositories

import (
	"context"
	"log"
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
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)

	defer cancel()

	cachedURL, err := r.redis.Get(ctx, shortCode).Result()

	if err == nil {
		return cachedURL, nil
	}

	collection := database.Client.Database("local").Collection("urls")

	var result OriginalURLOnly

	err = collection.FindOne(ctx, bson.M{"short_code": shortCode}).Decode(&result)

	if err != nil {
		return "", err
	}

	err = r.redis.Set(ctx, shortCode, result.OriginalURL, 60*time.Minute).Err()

	if err != nil {
		return "", err
	}

	return result.OriginalURL, nil
}

func (r *URLRepository) List() ([]models.URL, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)

	defer cancel()

	collection := database.Client.Database("local").Collection("urls")

	cursor, err := collection.Find(ctx, bson.M{})

	if err != nil {
		log.Fatal(err)
	}

	defer cursor.Close(ctx)

	var urls []models.URL

	if err = cursor.All(ctx, &urls); err != nil {
		return nil, err
	}

	return urls, nil
}
