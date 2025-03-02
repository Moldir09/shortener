package main

import (
	"github.com/Moldir09/shortener.git/internal/app/handler"
	"github.com/Moldir09/shortener.git/internal/app/service"
	"github.com/Moldir09/shortener.git/internal/app/storage"
	"net/http"
)

func main() {
	storage := storage.NewInMemoryURLStore()
	service := service.NewURLShortenerService(storage) // Создаем сервис
	handler := handler.NewHandler(service)             // Создаем обработчик

	mux := http.NewServeMux()     // Создаем Router
	handler.RegisterHandlers(mux) // Запуск сервера

	if err := http.ListenAndServe(":8080", mux); err != nil {
		panic(err)
	}
}
