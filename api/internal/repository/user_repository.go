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
	_, err := r.db.Exec("INSERT INTO users (name, password) VALUES ($1, $2)", user.Name, user.Password)

	return err
}

func (r *UserRepository) GetUser(name string) (model.User, error) {
	var user model.User

	err := r.db.QueryRow("SELECT id, password FROM users WHERE name=$1", name).Scan(&user.ID, &user.Password)

	if err != nil {
		return model.User{}, err
	}

	return user, nil
}
