package handler

import (
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
