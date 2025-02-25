package repository

import (
	"context"
	"database/sql"

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

func (r *URLRepository) SaveURL(url model.URL, userName string) error {
	_, err := r.db.Exec("INSERT INTO urls (original, short_code, created_by) VALUES ($1, $2, (select id from users where name = $3))", url.Original, url.Short, userName)

	return err
}

func (r *URLRepository) GetURL(short_code string) (string, error) {
	ctx := context.Background()

	cachedURL, err := r.redis.Get(ctx, short_code).Result()
	if err == nil {
		return cachedURL, nil
	}

	var original string
	err = r.db.QueryRow("SELECT original FROM urls WHERE short_code=$1", short_code).Scan(&original)
	if err != nil {
		return "", err
	}

	_, err = r.db.Exec("UPDATE urls set visit_count = visit_count + 1 where short_code=$1", short_code)

	if err != nil {
		return "", err
	}

	r.redis.Set(ctx, short_code, original, 0)

	return original, nil
}
