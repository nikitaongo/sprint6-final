package server

import (
	"log"
	"net/http"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/handlers"
)

type Server struct {
	Logger *log.Logger  // логгер
	Server *http.Server // http-сервер
}

// NewServer — создаёт и настраивает HTTP-сервер
func NewServer(logger *log.Logger) *Server {
	// Создание роутера
	mux := http.NewServeMux()

	// Регистрация хендлеров
	mux.HandleFunc("/", handlers.HtmlHandler)
	mux.HandleFunc("/upload", handlers.UploadHandler)

	// Настройка http-сервера
	httpSrv := &http.Server{
		Addr:         ":8080",
		Handler:      mux,
		ErrorLog:     logger,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}
	return &Server{
		Logger: logger,
		Server: httpSrv,
	}
}
