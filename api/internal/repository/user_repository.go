package repository

import (
	"context"
	"crypto/rand"
	"database/sql"
	"encoding/base64"
	"time"

	"github.com/jonasOli/url-shortener/api/internal/model"
	"github.com/redis/go-redis/v9"
)

type UserRepository struct {
	db    *sql.DB
	redis *redis.Client
}

func NewUserRepository(db *sql.DB, redis *redis.Client) *UserRepository {
	return &UserRepository{db, redis}
}

func (r *UserRepository) CreateUser(user model.User) (int, error) {
	var user_id int

	err := r.db.QueryRow(
		"INSERT INTO users (name, email, password, salt) VALUES ($1, $2, $3, $4) returning id",
		user.Name, user.Email, user.Password, user.Salt,
	).Scan(&user_id)

	if err != nil {
		return -1, err
	}

	return user_id, nil
}

func (r *UserRepository) GetUser(email string) (model.User, error) {
	var user model.User

	err := r.db.QueryRow("SELECT id, password, salt FROM users WHERE email=$1", email).Scan(&user.ID, &user.Password, &user.Salt)

	if err != nil {
		return model.User{}, err
	}

	return user, nil
}

func (r *UserRepository) CreateSessionId(user_id int) (string, error) {
	session_key, err := generateSessionKey()

	if err != nil {
		return "", err
	}
	ctx := context.Background()

	err = r.redis.Set(ctx, session_key, user_id, time.Hour).Err()

	return session_key, err
}

func (r *UserRepository) DeleteUserSession(session_key string) error {
	ctx := context.Background()
	err := r.redis.Del(ctx, session_key).Err()

	return err
}

func generateSessionKey() (string, error) {
	bytes := make([]byte, 128)
	_, err := rand.Read(bytes)

	if err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(bytes), nil
}
