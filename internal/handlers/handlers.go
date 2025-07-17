package handlers

import (
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/service"
)

func HtmlHandler(w http.ResponseWriter, r *http.Request) {
	// для корневого эндпоинта / нужно реализовать хендлер, который возвращает HTML из файла index.html.
	data, err := os.ReadFile("index.html")
	if err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

func UploadHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		http.Error(w, "error when parsing form", http.StatusInternalServerError)
		return
	}

	// получаем файл из формы
	file, handler, err := r.FormFile("myFile")
	if err != nil {
		http.Error(w, "error when receiving the file", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	// читаем его
	data, err := io.ReadAll(file)
	if err != nil {
		http.Error(w, "error when reading the file", http.StatusInternalServerError)
	}

	// конвертация
	result, err := service.TextConverter(string(data))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// имя нового файла
	extension := filepath.Ext(handler.Filename)
	time := time.Now().UTC().Format("02-01-2006_15-04-05")
	filename := "converted_" + time + extension

	// создание нового файла
	output, err := os.Create(filename)
	if err != nil {
		http.Error(w, "error when creating the file", http.StatusInternalServerError)
		return
	}
	defer output.Close()

	// запись в новый файл
	_, err = output.Write([]byte(result))
	if err != nil {
		http.Error(w, "error when writing to the file", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(result))
}
