package repository

import (
	"context"
	"event_service/internal/models"
	"fmt"

	"github.com/jmoiron/sqlx"
)

// DateRepositoryInterface интерфейс для работы с датами
type DateRepositoryInterface interface {
	Create(ctx context.Context, event models.DateEvent) error
	GetAll(ctx context.Context) ([]models.DateEvent, error)
	Delete(ctx context.Context, id int) error
}

// DateRepository репозиторий для работы с датами
type DateRepository struct {
	db *sqlx.DB
}

// NewDateRepository создает новый репозиторий дат
func NewDateRepository(db *sqlx.DB) *DateRepository {
	return &DateRepository{
		db: db,
	}
}

// Create создает новое событие
func (r *DateRepository) Create(ctx context.Context, event models.DateEvent) error {
	query := `
		INSERT INTO date_events (date, description, created_at)
		VALUES ($1, $2, $3)
	`

	_, err := r.db.ExecContext(ctx, query,
		event.Date,
		event.Description,
		event.CreatedAt,
	)

	if err != nil {
		return fmt.Errorf("failed to create date event: %w", err)
	}

	return nil
}

// GetAll получает все события
func (r *DateRepository) GetAll(ctx context.Context) ([]models.DateEvent, error) {
	query := `
		SELECT id, date, description, created_at
		FROM date_events
		ORDER BY date DESC
	`

	var events []models.DateEvent
	err := r.db.SelectContext(ctx, &events, query)
	if err != nil {
		return nil, fmt.Errorf("failed to get all date events: %w", err)
	}

	return events, nil
}

// Delete удаляет событие по ID
func (r *DateRepository) Delete(ctx context.Context, id int) error {
	query := `DELETE FROM date_events WHERE id = $1`

	_, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return fmt.Errorf("failed to delete date event: %w", err)
	}

	return nil
}
