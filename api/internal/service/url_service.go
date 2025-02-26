package service

import (
	"github.com/google/uuid"
	"github.com/jonasOli/url-shortener/api/internal/model"
	"github.com/jonasOli/url-shortener/api/internal/repository"
)

type URLService struct {
	repo *repository.URLRepository
}

func NewURLService(repo *repository.URLRepository) *URLService {
	return &URLService{repo}
}

func (s *URLService) ShortenURL(original string, user_id int) (string, error) {
	short_url := uuid.New().String()[:8]

	url := model.URL{
		Original: original,
		Short:    short_url,
	}

	err := s.repo.SaveURL(url, user_id)

	if err != nil {
		return "", err
	}

	return short_url, nil
}

func (s *URLService) GetOriginalURL(short_code string) (string, error) {
	original_url, err := s.repo.GetURL(short_code)

	if err != nil {
		return "", err
	}

	return original_url, nil
}
