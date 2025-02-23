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

	mux := http.NewServeMux() // Создаем Router
	handler.RegisterHandlers(mux)

	if err := http.ListenAndServe(":8080", mux); err != nil {
		panic(err)
	}
}
