package config

import (
	"os"
	"strconv"
)

// Config конфигурация приложения
type Config struct {
	Database DatabaseConfig
	Server   ServerConfig
}

// DatabaseConfig конфигурация базы данных
type DatabaseConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	DBName   string
	SSLMode  string
}

// ServerConfig конфигурация сервера
type ServerConfig struct {
	Port string
}

// Load загружает конфигурацию из переменных окружения
func Load() *Config {
	port, _ := strconv.Atoi(getEnv("DB_PORT", "5432"))

	return &Config{
		Database: DatabaseConfig{
			Host:     getEnv("DB_HOST", "localhost"),
			Port:     port,
			User:     getEnv("DB_USER", "together_user"),
			Password: getEnv("DB_PASSWORD", "jj0545bk"),
			DBName:   getEnv("DB_NAME", "together_db"),
			SSLMode:  getEnv("DB_SSLMODE", "disable"),
		},
		Server: ServerConfig{
			Port: getEnv("SERVER_PORT", "8080"),
		},
	}
}

// getEnv получает значение переменной окружения или возвращает значение по умолчанию
func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
