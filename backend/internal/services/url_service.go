package services

import (
	"time"
	"url-shortener/internal/models"
	"url-shortener/internal/repositories"

	"github.com/google/uuid"
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
	shortCode := uuid.New().String()[:8]

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
