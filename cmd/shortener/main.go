package main

import (
	"github.com/Moldir09/shortener.git/internal/app/handler"
	"github.com/Moldir09/shortener.git/internal/app/middleware"
	"github.com/Moldir09/shortener.git/internal/app/service"
	"github.com/Moldir09/shortener.git/internal/app/storage"
	"github.com/Moldir09/shortener.git/internal/config"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"log"
)

func main() {

	cfg := config.NewConfig()

	storage := storage.NewInMemoryURLStore()
	service := service.NewURLShortenerService(storage, cfg.BaseURL) // Создаем сервис
	handler := handler.NewHandler(service)

	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatalf("failed to create logger : %v", err)
	}
	defer logger.Sync()

	sugar := logger.Sugar()
	r := gin.Default()
	r.Use(middleware.WithLogging(sugar))

	handler.RegisterRoutes(r)
	if err := r.Run(cfg.ServerAddress); err != nil {
		panic(err)
	}
}
