package services

import (
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
	}

	return s.urlRepository.Create(url)
}
