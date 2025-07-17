package service

import (
	"fmt"
	"strings"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

func TextConverter(input string) (string, error) {
	if len(input) == 0 {
		return "", fmt.Errorf("string is empty: \"%s\"", input)
	}

	//строка - это текст, если в наличии любой символ кроме знаков азбуки Морзе
	isText := strings.ContainsAny(input, "АБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯабвгдеёжзийклмнопрстуфхцчшщъыьэюя1234567890,:?`'()“”\"\\/")
	if isText {
		return morse.ToMorse(input), nil
	}
	return morse.ToText(input), nil
}
