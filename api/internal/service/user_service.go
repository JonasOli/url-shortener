package service

import (
	"crypto/rsa"
	"errors"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/jonasOli/url-shortener/api/internal/model"
	"github.com/jonasOli/url-shortener/api/internal/repository"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{repo}
}

func (s *UserService) CreateUser(name string, password string) error {
	trimmedName := strings.TrimSpace(name)
	trimmedPassword := strings.TrimSpace(password)

	if trimmedName == "" || trimmedPassword == "" {
		return errors.New("Name or password cannot be empty!")
	}

	hashedPassword, err := _HashPassword(trimmedPassword)

	if err != nil {
		return err
	}

	user := model.User{
		Name:     trimmedName,
		Password: hashedPassword,
	}

	err = s.repo.CreateUser(user)

	return err
}

func (s *UserService) Login(name string, password string, privateKey *rsa.PrivateKey) (string, *fiber.Error) {
	user, err := s.repo.GetUser(name)

	if err != nil {
		return "", fiber.ErrInternalServerError
	}

	if validPassword := _VerifyPassword(password, user.Password); !validPassword {
		return "", fiber.NewError(400, "Invalid password")
	}

	claims := jwt.MapClaims{
		"name":  name,
		"admin": true,
		"exp":   time.Now().Add(time.Hour * 72).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)

	t, err := token.SignedString(privateKey)

	if err != nil {
		return "", fiber.ErrInternalServerError
	}

	return t, nil
}

func _HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func _VerifyPassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
