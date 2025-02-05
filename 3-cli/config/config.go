package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Key         string
	ContentType string
}

func NewConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		panic("Не удалось найти env файл")
	}
	key := os.Getenv("KEY")
	contentType := os.Getenv("CONTENT_TYPE")
	if key == "" {
		panic("Не передан параметр KEY в переменные окружения")
	}
	if contentType == "" {
		panic("Не передан параметр CONTENT-TYPE в переменные окружения")
	}
	return &Config{
		Key:         key,
		ContentType: contentType,
	}
}
