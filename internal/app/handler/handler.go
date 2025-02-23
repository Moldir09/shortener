package handler

import (
	"github.com/Moldir09/shortener.git/internal/app/service"
	"io"
	"net/http"
	"strings"
)

// Handler структура, содержащая ссылки на сервисы
type Handler struct {
	URLShortenerService service.URLShortener
}

// NewHandler создаёт новый экземпляр Handler с зависимостями
func NewHandler(urlShortenerService service.URLShortener) *Handler {
	return &Handler{
		URLShortenerService: urlShortenerService,
	}
}

// RegisterHandlers регистрирует маршруты для обработки HTTP запросов
func (h *Handler) RegisterHandlers() {
	http.HandleFunc("/", h.handleRequest)
}

// handleRequest диспетчеризует запросы в зависимости от HTTP метода
func (h *Handler) handleRequest(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		h.handleGet(w, r)
	case http.MethodPost:
		h.handlePost(w, r)
	default:
		http.Error(w, "Unsupported HTTP method", http.StatusMethodNotAllowed)
	}
}

// handleGet обрабатывает GET-запросы, извлекая данные
func (h *Handler) handleGet(w http.ResponseWriter, r *http.Request) {
	shortURL := strings.TrimPrefix(r.URL.Path, "/")
	originalURL, err := h.URLShortenerService.ResolveURL(shortURL)
	if err != nil {
		http.Error(w, "URL not found", http.StatusNotFound)
		return
	}
	http.Redirect(w, r, originalURL, http.StatusFound)
}

// handlePost обрабатывает POST-запросы, создавая новые короткие URL
func (h *Handler) handlePost(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusBadRequest)
		return
	}
	originalURL := strings.TrimSpace(string(body))
	shortURL, err := h.URLShortenerService.ShortenURL(originalURL)
	if err != nil {
		http.Error(w, "Failed to shorten URL", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(shortURL))
}
