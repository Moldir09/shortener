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

func (h *Handler) RegisterHandlers(mux *http.ServeMux) {
	mux.HandleFunc("/", h.handleRequest)
}

func (h *Handler) handleRequest(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		h.handlePost(w, r)
	case http.MethodGet:
		h.handleGet(w, r)
	default:
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
}

func (h *Handler) handleGet(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}

	shortURL := strings.TrimPrefix(r.URL.Path, "/")
	if shortURL == "" {
		http.Error(w, "URL not found", http.StatusNotFound)
		return
	}
	originalURL, err := h.URLShortenerService.ResolveURL(shortURL)
	if err != nil {
		http.Error(w, "URL not found", http.StatusNotFound)
		return
	}
	g
	http.Redirect(w, r, originalURL, http.StatusTemporaryRedirect)
}

func (h *Handler) handlePost(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

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
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(shortURL))
}
