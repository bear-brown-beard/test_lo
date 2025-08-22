package di

import (
	"fmt"
	"log"

	"event_service/internal/config"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

// DatabaseContainer контейнер для базы данных
type DatabaseContainer struct {
	DB *sqlx.DB
}

// NewDatabaseContainer создает новый контейнер для базы данных
func NewDatabaseContainer() (*DatabaseContainer, error) {
	// Загружаем конфигурацию
	cfg := config.Load()

	// Формируем строку подключения
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName, cfg.DBSSLMode)

	// Подключаемся к базе данных
	db, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	// Проверяем соединение
	if err := db.Ping(); err != nil {
		db.Close()
		return nil, fmt.Errorf("database ping failed: %w", err)
	}

	log.Println("Подключение к БД установлено")

	return &DatabaseContainer{
		DB: db,
	}, nil
}

func (c *DatabaseContainer) Close() error {
	return c.DB.Close()
}
