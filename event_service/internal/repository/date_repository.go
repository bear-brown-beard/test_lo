package repository

import (
	"context"
	"event_service/internal/database"
	"event_service/internal/models"
	"fmt"
	"sort"
	"strings"
)

// DateRepositoryInterface интерфейс для работы с датами
type DateRepositoryInterface interface {
	Create(ctx context.Context, event models.DateEvent) error
	GetAll(ctx context.Context) (map[string][]models.DateEvent, error)
	Delete(ctx context.Context, id int) error
}

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
	// Получаем соединение с базой данных
	postgresDB, ok := r.db.(*database.PostgresDB)
	if !ok {
		return fmt.Errorf("unsupported database type")
	}

	db := postgresDB.GetDB()

	query := `
		INSERT INTO date_events (person1_name, person2_name, date, description, created_at)
		VALUES ($1, $2, $3, $4, $5)
	`

	_, err := db.ExecContext(ctx, query,
		event.Person1Name,
		event.Person2Name,
		event.Date,
		event.Description,
		event.CreatedAt,
	)

	if err != nil {
		return fmt.Errorf("failed to create date event: %w", err)
	}

	return nil
}

// GetAll получает все события, сгруппированные по парам
func (r *DateRepository) GetAll(ctx context.Context) (map[string][]models.DateEvent, error) {
	// Получаем соединение с базой данных
	postgresDB, ok := r.db.(*database.PostgresDB)
	if !ok {
		return nil, fmt.Errorf("unsupported database type")
	}

	db := postgresDB.GetDB()

	query := `
		SELECT id, person1_name, person2_name, date, description, created_at
		FROM date_events
		ORDER BY person1_name, person2_name, date
	`

	var events []models.DateEvent
	err := db.SelectContext(ctx, &events, query)
	if err != nil {
		return nil, fmt.Errorf("failed to get all date events: %w", err)
	}

	// Группируем события по парам
	result := make(map[string][]models.DateEvent)
	for _, event := range events {
		key := r.generatePairKey(event.Person1Name, event.Person2Name)
		result[key] = append(result[key], event)
	}

	return result, nil
}

// Delete удаляет событие по ID
func (r *DateRepository) Delete(ctx context.Context, id int) error {
	// Получаем соединение с базой данных
	postgresDB, ok := r.db.(*database.PostgresDB)
	if !ok {
		return fmt.Errorf("unsupported database type")
	}

	db := postgresDB.GetDB()

	query := `
		DELETE FROM date_events
		WHERE id = $1
	`

	result, err := db.ExecContext(ctx, query, id)
	if err != nil {
		return fmt.Errorf("failed to delete date event: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("date event with id %d not found", id)
	}

	return nil
}

// generatePairKey генерирует ключ для пары людей
func (r *DateRepository) generatePairKey(person1, person2 string) string {
	names := []string{person1, person2}
	sort.Strings(names)
	return strings.Join(names, "_")
}
