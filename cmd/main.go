package main

import (
	"github.com/Moldir09/shortener.git/internal/app/handler"
	"github.com/Moldir09/shortener.git/internal/app/service"
	"github.com/Moldir09/shortener.git/internal/app/storage"
	"net/http"
)

func main() {
	storage := storage.NewInMemoryURLStore()           // Создаем хранилище
	service := service.NewURLShortenerService(storage) // Создаем сервис
	handler := handler.NewHandler(service)             // Создаем обработчик

	handler.RegisterHandlers() // Регистрируем обработчики

	// Запуск сервера
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
