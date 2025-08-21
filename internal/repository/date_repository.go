package repository

import (
	"context"
	"together_service/internal/database"
	"together_service/internal/models"
)

// DateRepository репозиторий для работы с датами
type DateRepository struct {
	db database.Database
}

// NewDateRepository создает новый репозиторий дат
func NewDateRepository(db database.Database) *DateRepository {
	return &DateRepository{
		db: db,
	}
}

// Create создает новое событие
func (r *DateRepository) Create(ctx context.Context, event models.DateEvent) error {
	return r.db.CreateDateEvent(ctx, event)
}

// GetByPair получает события для конкретной пары
func (r *DateRepository) GetByPair(ctx context.Context, person1Name, person2Name string) ([]models.DateEvent, error) {
	return r.db.GetDateEvents(ctx, person1Name, person2Name)
}

// GetAll получает все события, сгруппированные по парам
func (r *DateRepository) GetAll(ctx context.Context) (map[string][]models.DateEvent, error) {
	return r.db.GetAllDateEvents(ctx)
}

// GetByID получает событие по ID
func (r *DateRepository) GetByID(ctx context.Context, id int) (*models.DateEvent, error) {
	return r.db.GetDateEventByID(ctx, id)
}
