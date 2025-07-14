package services

import (
	"time"
	"url-shortener/internal/models"
	"url-shortener/internal/repositories"

	"github.com/google/uuid"
)

type URLService struct {
	urlRepository *repositories.URLRepository
}

func NewURLService(url_repository *repositories.URLRepository) *URLService {
	return &URLService{urlRepository: url_repository}
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

func (s *URLService) FindByShortCode(shortCode string) (models.URL, error) {
	return s.urlRepository.FindByShortCode(shortCode)
}
