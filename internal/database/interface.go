package database

import (
	"context"
	"together_service/internal/models"
)

// Database интерфейс для работы с базой данных
type Database interface {
	// Подключение и управление
	Connect(ctx context.Context) error
	Close() error
	Ping(ctx context.Context) error

	// CRUD операции для дат
	CreateDateEvent(ctx context.Context, event models.DateEvent) error
	GetDateEvents(ctx context.Context, person1Name, person2Name string) ([]models.DateEvent, error)
	GetAllDateEvents(ctx context.Context) (map[string][]models.DateEvent, error)
	GetDateEventByID(ctx context.Context, id int) (*models.DateEvent, error)
}
