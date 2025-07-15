package services

import (
	"math/rand"
	"time"
	"url-shortener/internal/models"
	"url-shortener/internal/repositories"

	"github.com/redis/go-redis/v9"
)

type URLService struct {
	urlRepository *repositories.URLRepository
	redis         *redis.Client
}

func NewURLService(redis *redis.Client) *URLService {
	repository := repositories.NewURLRepository(redis)

	return &URLService{urlRepository: repository, redis: redis}
}

func (s *URLService) CreateShortURL(originalURL string) error {
	shortCode := generateShortCode()

	url := &models.URL{
		OriginalURL: originalURL,
		ShortCode:   shortCode,
		CreatedAt:   time.Now(),
	}

	return s.urlRepository.Create(url)
}

func (s *URLService) FindByShortCode(shortCode string) (string, error) {
	return s.urlRepository.FindByShortCode(shortCode)
}

func generateShortCode() string {
	const base62Alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	const keyLength = 8

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	var shortKey = make([]byte, keyLength)

	for i := range shortKey {
		shortKey[i] = base62Alphabet[r.Intn(len(base62Alphabet))]
	}

	return string(shortKey)
}
