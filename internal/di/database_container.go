package di

import (
	"context"
	"together_service/internal/config"
	"together_service/internal/database"
)

// DatabaseContainer контейнер для базы данных
type DatabaseContainer struct {
	DB database.Database
}

// NewDatabaseContainer создает новый контейнер для базы данных
func NewDatabaseContainer() (*DatabaseContainer, error) {
	// Загружаем конфигурацию
	cfg := config.Load()

	// Создаем конфигурацию базы данных
	dbConfig := database.Config{
		Host:     cfg.Database.Host,
		Port:     cfg.Database.Port,
		User:     cfg.Database.User,
		Password: cfg.Database.Password,
		DBName:   cfg.Database.DBName,
		SSLMode:  cfg.Database.SSLMode,
	}

	// Подключаемся к базе данных
	db, err := database.NewPostgresDB(dbConfig)
	if err != nil {
		return nil, err
	}

	// Проверяем соединение
	ctx := context.Background()
	if err := db.Ping(ctx); err != nil {
		db.Close()
		return nil, err
	}

	return &DatabaseContainer{
		DB: db,
	}, nil
}

// Close закрывает контейнер базы данных
func (c *DatabaseContainer) Close() error {
	return c.DB.Close()
}
