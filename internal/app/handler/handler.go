package handler

import (
	"encoding/json"
	"github.com/Moldir09/shortener.git/internal/app/service"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"strings"
)

type Handler struct {
	URLShortenerService service.URLShortener
}

func NewHandler(urlShortenerService service.URLShortener) *Handler {
	return &Handler{
		URLShortenerService: urlShortenerService,
	}
}

func (h *Handler) RegisterRoutes(r *gin.Engine) {
	r.GET("/:short", h.handleGet)
	r.POST("/", h.handlePost)
	r.POST("api/shorten", h.GetShortenURL)
}

func (h *Handler) GetShortenURL(c *gin.Context) {
	var req Response

	if err := json.NewDecoder(c.Request.Body).Decode(&req); err != nil {
		c.Writer.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(c.Writer).Encode(gin.H{"error": "invalid JSON"})
		return
	}

	result, err := h.URLShortenerService.ShortenURL(req.URL)
	if err != nil {
		c.Writer.WriteHeader(http.StatusInternalServerError)
		_ = json.NewEncoder(c.Writer).Encode(gin.H{"error": err.Error()})
		return
	}

	c.Writer.Header().Set("Content-Type", "application/json")
	c.Writer.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(c.Writer).Encode(Result{Result: result})
}

func (h *Handler) handleGet(c *gin.Context) {

	shortURL := c.Param("short")

	originalURL, err := h.URLShortenerService.ResolveURL(shortURL)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "URL not found"})
		return
	}

	c.Redirect(http.StatusTemporaryRedirect, originalURL)
}

func (h *Handler) handlePost(c *gin.Context) {

	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read request body",
		})
		return
	}

	originalURL := strings.TrimSpace(string(body))
	shortURL, err := h.URLShortenerService.ShortenURL(originalURL)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to shorten URL"})
		return
	}

	c.Data(http.StatusCreated, "text/plain", []byte(shortURL))
}
