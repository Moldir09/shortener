package storage

import (
	"fmt"
	"sync"
)

type Storage interface {
	Save(shortURL, originalURL string) error
	Load(shortURL string) (string, error)
	GetByOriginal(original string) (string, bool)
}

type InMemoryURLStore struct {
	originalURLs map[string]string
	shortURLs    map[string]string
	mu           sync.Mutex
}

func NewInMemoryURLStore() *InMemoryURLStore {
	return &InMemoryURLStore{
		originalURLs: make(map[string]string),
		shortURLs:    make(map[string]string),
	}
}

func (s *InMemoryURLStore) Save(shortURL, originalURL string) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.originalURLs[shortURL] = originalURL
	s.shortURLs[originalURL] = shortURL
	return nil
}

func (s *InMemoryURLStore) Load(shortURL string) (string, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	originalURL, exists := s.originalURLs[shortURL]
	if !exists {
		return "", fmt.Errorf("URL не найден")
	}
	return originalURL, nil
}

func (s *InMemoryURLStore) GetByOriginal(original string) (string, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()

	short, ok := s.shortURLs[original]
	return short, ok
}
