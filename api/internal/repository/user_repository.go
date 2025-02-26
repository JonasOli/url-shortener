package repository

import (
	"database/sql"

	"github.com/jonasOli/url-shortener/api/internal/model"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db}
}

func (r *UserRepository) CreateUser(user model.User) error {
	_, err := r.db.Exec(
		"INSERT INTO users (name, email, password, salt) VALUES ($1, $2, $3, $4)",
		user.Name, user.Email, user.Password, user.Salt,
	)

	return err
}

func (r *UserRepository) GetUser(email string) (model.User, error) {
	var user model.User

	err := r.db.QueryRow("SELECT id, password, salt FROM users WHERE email=$1", email).Scan(&user.ID, &user.Password, &user.Salt)

	if err != nil {
		return model.User{}, err
	}

	return user, nil
}
