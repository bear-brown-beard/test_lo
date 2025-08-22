package config

import (
	"os"
	"together_service/internal/database"
)

// Config конфигурация приложения
type Config struct {
	Database database.Config
	Server   ServerConfig
}

// ServerConfig конфигурация сервера
type ServerConfig struct {
	Port string
}

// Load загружает конфигурацию из переменных окружения
func Load() *Config {
	return &Config{
		Database: database.Config{
			Host:     os.Getenv("DB_HOST"),
			Port:     os.Getenv("DB_PORT"),
			User:     os.Getenv("DB_USER"),
			Password: os.Getenv("DB_PASSWORD"),
			DBName:   os.Getenv("DB_NAME"),
			SSLMode:  os.Getenv("DB_SSLMODE"),
		},
		Server: ServerConfig{
			Port: os.Getenv("SERVER_PORT"),
		},
	}
}
