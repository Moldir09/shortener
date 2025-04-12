package service

import (
	"crypto/rand"
	"encoding/base64"
	"github.com/Moldir09/shortener.git/internal/app/storage"
)

type URLShortener interface {
	ShortenURL(originalURL string) (string, error)
	ResolveURL(shortURL string) (string, error)
}

type Service struct {
	store   storage.Storage
	baseURL string
}

func NewURLShortenerService(store storage.Storage, baseURL string) URLShortener {
	return &Service{
		store:   store,
		baseURL: baseURL,
	}
}

func (s *Service) ShortenURL(originalURL string) (string, error) {
	shortURL := generateShortURL()
	if err := s.store.Save(shortURL, originalURL); err != nil {
		return "", err
	}
	return s.baseURL + "/" + shortURL, nil
}

func (s *Service) ResolveURL(shortURL string) (string, error) {
	return s.store.Load(shortURL)
}

func generateShortURL() string {
	b := make([]byte, 6) // Генерируем 6 байт случайных данных
	_, err := rand.Read(b)
	if err != nil {
		panic(err)
	}
	return base64.URLEncoding.EncodeToString(b)[:8] // Кодируем в base64 и обрезаем до 8 символов
}
