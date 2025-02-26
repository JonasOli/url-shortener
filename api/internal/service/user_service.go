package service

import (
	"crypto/rand"
	"encoding/base64"
	"errors"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
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

func (s *UserService) Signup(name string, email string, password string) (string, error) {
	trimmedName := strings.TrimSpace(name)
	trimmedEmail := strings.TrimSpace(email)
	trimmedPassword := strings.TrimSpace(password)

	if trimmedName == "" || trimmedPassword == "" || trimmedEmail == "" {
		return "", errors.New("Name or password cannot be empty!")
	}

	salt, err := generateSalt()

	if err != nil {
		return "", err
	}

	hashedPassword, err := hashPassword(strings.Join([]string{trimmedPassword, salt}, ""))

	if err != nil {
		return "", err
	}

	user := model.User{
		Name:     trimmedName,
		Email:    trimmedEmail,
		Password: hashedPassword,
		Salt:     salt,
	}

	user_id, err := s.repo.CreateUser(user)

	if err != nil {
		log.Infof("Error on CreateUser: %s", err)

		return "", fiber.NewError(500, "Invalid password")
	}

	session_key, err := s.repo.CreateSessionId(user_id)

	return session_key, err
}

func (s *UserService) Signin(email string, password string) (string, *fiber.Error) {
	user, err := s.repo.GetUser(email)

	if err != nil {
		return "", fiber.ErrInternalServerError
	}

	if validPassword := verifyPassword(strings.Join([]string{password, user.Salt}, ""), user.Password); !validPassword {
		return "", fiber.NewError(400, "Invalid password")
	}

	session_key, err := s.repo.CreateSessionId(user.ID)

	if err != nil {
		return "", fiber.ErrInternalServerError
	}

	return session_key, nil
}

func (s *UserService) Signout(session_key string) error {
	err := s.repo.DeleteUserSession(session_key)

	return err
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)

	return string(bytes), err
}

func verifyPassword(password string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))

	return err == nil
}

func generateSalt() (string, error) {
	bytes := make([]byte, 16)
	_, err := rand.Read(bytes)

	if err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(bytes), nil
}
