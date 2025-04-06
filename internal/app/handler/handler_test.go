package handler

import (
	"errors"
	"github.com/Moldir09/shortener.git/internal/app/service"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

type mockService struct{}

func (m *mockService) ShortenURL(originalURL string) (string, error) {
	return "short123", nil
}

func (m *mockService) ResolveURL(shortURL string) (string, error) {
	if shortURL == "short123" {
		return "https://practicum.yandex.kz", nil
	}
	return "", errors.New("URL not found")
}

func TestHandler_handlePost(t *testing.T) {
	gin.SetMode(gin.TestMode)
	type fields struct {
		URLShortenerService service.URLShortener
	}
	tests := []struct {
		name     string
		fields   fields
		setupReq func() *http.Request
		wantCode int
		wantBody string
	}{
		{
			name: "valid POST request",
			fields: fields{
				URLShortenerService: &mockService{},
			},
			setupReq: func() *http.Request {
				body := strings.NewReader("https://test.com")
				req := httptest.NewRequest(http.MethodPost, "/", body)
				req.Header.Set("Content-Type", "text/plain")
				return req
			},
			wantCode: http.StatusCreated,
			wantBody: "short123",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &Handler{
				URLShortenerService: tt.fields.URLShortenerService,
			}

			req := tt.setupReq()
			rr := httptest.NewRecorder()

			c, _ := gin.CreateTestContext(rr)
			c.Request = req

			h.handlePost(c)

			if rr.Code != tt.wantCode {
				t.Errorf("expected status %d, got %d", tt.wantCode, rr.Code)
			}

			if strings.TrimSpace(rr.Body.String()) != tt.wantBody {
				t.Errorf("expected body '%s', got '%s'", tt.wantBody, rr.Body.String())
			}

		})
	}
}

func TestHandler_handleGet(t *testing.T) {
	gin.SetMode(gin.TestMode)
	type fields struct {
		URLShortenerService service.URLShortener
	}
	tests := []struct {
		name       string
		fields     fields
		setupReq   func(shortParam string) *http.Request
		wantCode   int
		shortParam string
	}{
		{
			name:   "valid GET request",
			fields: fields{URLShortenerService: &mockService{}},
			setupReq: func(shortParam string) *http.Request {
				return httptest.NewRequest(http.MethodGet, "/"+shortParam, nil)

			},
			wantCode:   http.StatusTemporaryRedirect,
			shortParam: "short123",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &Handler{
				URLShortenerService: tt.fields.URLShortenerService,
			}

			req := tt.setupReq(tt.shortParam)
			rr := httptest.NewRecorder()

			c, _ := gin.CreateTestContext(rr)
			c.Request = req
			c.Params = gin.Params{gin.Param{Key: "short", Value: tt.shortParam}}

			h.handleGet(c)

			require.Equal(t, tt.wantCode, rr.Code)
			require.Equal(t, "https://practicum.yandex.kz", rr.Header().Get("Location"))

		})
	}
}
