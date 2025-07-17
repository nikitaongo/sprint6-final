package main

import (
	"log"
	"os"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/server"
)

func main() {
	logger := log.New(os.Stdout, "serv: ", log.LstdFlags)

	// Создание сервера через ваш пакет
	serv := server.NewServer(logger)

	// Запуск сервера
	logger.Println("starting server on port:", serv.Server.Addr)
	err := serv.Server.ListenAndServe()
	if err != nil {
		logger.Fatal("error when starting server:", err)
	}
}
