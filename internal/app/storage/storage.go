package storage

import (
	"fmt"
	"sync"
)

type Storage interface {
	Save(shortURL, originalURL string) error
	Load(shortURL string) (string, error)
}

type InMemoryURLStore struct {
	store map[string]string
	mu    sync.Mutex
}

func NewInMemoryURLStore() *InMemoryURLStore {
	return &InMemoryURLStore{
		store: make(map[string]string),
	}
}

func (s *InMemoryURLStore) Save(shortURL, originalURL string) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.store[shortURL] = originalURL
	return nil
}

func (s *InMemoryURLStore) Load(shortURL string) (string, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	originalURL, exists := s.store[shortURL]
	if !exists {
		return "", fmt.Errorf("URL не найден")
	}
	return originalURL, nil
}
