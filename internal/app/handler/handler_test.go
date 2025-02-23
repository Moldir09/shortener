package handler

import (
	"bytes"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
)

var mockStore = make(map[string]string)

type mockService struct{}

func (m *mockService) ShortenURL(originalURL string) (string, error) {
	return "short123", nil
}

func (m *mockService) ResolveURL(shortURL string) (string, error) {
	if shortURL == "short123" {
		return "https://test.com", nil
	}
	return "", errors.New("URL not found")
}

func newTestHandler() *Handler {
	return NewHandler(&mockService{})
}

func TestHandlePost(t *testing.T) {
	h := newTestHandler()

	req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBufferString("https://test.com"))
	w := httptest.NewRecorder()

	// Вызываем хендлер
	h.handlePost(w, req)

	// Проверяем статус-код
	res := w.Result()
	defer res.Body.Close()

	if res.StatusCode != http.StatusCreated {
		t.Errorf("Expected status %d, got %d", http.StatusCreated, res.StatusCode)
	}
}

func TestHandleGet(t *testing.T) {
	h := newTestHandler()
	mockStore["short123"] = "https://test.com"

	req := httptest.NewRequest(http.MethodGet, "/short123", nil)
	w := httptest.NewRecorder()

	h.handleGet(w, req)
	res := w.Result()
	defer res.Body.Close()
	if res.StatusCode != http.StatusFound {
		t.Errorf("Expected status %d, got %d", http.StatusFound, res.StatusCode)
	}

	location := res.Header.Get("Location")
	if location != "https://test.com" {
		t.Errorf("Expected Location header %s, got %s", "https://example.com", location)
	}
}
