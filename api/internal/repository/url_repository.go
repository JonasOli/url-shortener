package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/jonasOli/url-shortener/api/internal/model"
	"github.com/redis/go-redis/v9"
)

type URLRepository struct {
	db    *sql.DB
	redis *redis.Client
}

func NewURLRepository(db *sql.DB, redis *redis.Client) *URLRepository {
	return &URLRepository{db, redis}
}

func (r *URLRepository) SaveURL(url model.URL) error {
	_, err := r.db.Exec("INSERT INTO urls (id, original, short) VALUES ($1, $2, $3)", url.ID, url.Original, url.Short)

	return err
}

func (r *URLRepository) GetURL(short string) (string, error) {
	ctx := context.Background()

	cachedURL, err := r.redis.Get(ctx, short).Result()
	if err == nil {
		return cachedURL, nil
	}

	var original string
	err = r.db.QueryRow("SELECT original FROM urls WHERE short=$1", short).Scan(&original)
	if err != nil {
		return "", errors.New("URL not found")
	}

	r.redis.Set(ctx, short, original, 0)

	return original, nil
}
