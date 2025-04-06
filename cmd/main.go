package main

import (
	"github.com/Moldir09/shortener.git/internal/app/handler"
	"github.com/Moldir09/shortener.git/internal/app/service"
	"github.com/Moldir09/shortener.git/internal/app/storage"
	"github.com/gin-gonic/gin"
)

func main() {
	storage := storage.NewInMemoryURLStore()           // Создаем хранилище
	service := service.NewURLShortenerService(storage) // Создаем сервис
	handler := handler.NewHandler(service)             // Создаем обработчик

	r := gin.Default()
	handler.RegisterRoutes(r)

	if err := r.Run(":8080"); err != nil {
		panic(err)
	}
}
