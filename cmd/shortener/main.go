package main

import (
	"github.com/Moldir09/shortener.git/internal/app/handler"
	"github.com/Moldir09/shortener.git/internal/app/service"
	"github.com/Moldir09/shortener.git/internal/app/storage"
	"github.com/Moldir09/shortener.git/internal/config"
	"github.com/gin-gonic/gin"
)

func main() {

	cfg := config.NewConfig()

	storage := storage.NewInMemoryURLStore()
	service := service.NewURLShortenerService(storage, cfg.BaseURL) // Создаем сервис
	handler := handler.NewHandler(service)

	r := gin.Default()

	handler.RegisterRoutes(r)
	if err := r.Run(cfg.ServerAddress); err != nil {
		panic(err)
	}

}
