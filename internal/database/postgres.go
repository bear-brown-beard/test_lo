package database

import (
	"context"
	"fmt"
	"sort"
	"strings"
	"time"

	"together_service/internal/models"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

// PostgresDB реализация интерфейса Database для PostgreSQL
type PostgresDB struct {
	db *sqlx.DB
}

// Config конфигурация для подключения к PostgreSQL
type Config struct {
	Host     string
	Port     int
	User     string
	Password string
	DBName   string
	SSLMode  string
}

// NewPostgresDB создает новый экземпляр PostgreSQL базы данных
func NewPostgresDB(config Config) (*PostgresDB, error) {
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		config.Host, config.Port, config.User, config.Password, config.DBName, config.SSLMode)

	db, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	// Настройка пула соединений
	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(25)
	db.SetConnMaxLifetime(5 * time.Minute)

	return &PostgresDB{db: db}, nil
}

// Connect подключается к базе данных
func (p *PostgresDB) Connect(ctx context.Context) error {
	return p.db.PingContext(ctx)
}

// Close закрывает соединение с базой данных
func (p *PostgresDB) Close() error {
	return p.db.Close()
}

// Ping проверяет соединение с базой данных
func (p *PostgresDB) Ping(ctx context.Context) error {
	return p.db.PingContext(ctx)
}

// CreateDateEvent создает новое событие
func (p *PostgresDB) CreateDateEvent(ctx context.Context, event models.DateEvent) error {
	query := `
		INSERT INTO date_events (person1_name, person2_name, date, description, created_at)
		VALUES ($1, $2, $3, $4, $5)
	`

	_, err := p.db.ExecContext(ctx, query,
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

// GetDateEvents получает события для конкретной пары
func (p *PostgresDB) GetDateEvents(ctx context.Context, person1Name, person2Name string) ([]models.DateEvent, error) {
	query := `
		SELECT id, person1_name, person2_name, date, description, created_at
		FROM date_events
		WHERE (person1_name = $1 AND person2_name = $2) OR (person1_name = $2 AND person2_name = $1)
		ORDER BY date ASC
	`

	var events []models.DateEvent
	err := p.db.SelectContext(ctx, &events, query, person1Name, person2Name)
	if err != nil {
		return nil, fmt.Errorf("failed to get date events: %w", err)
	}

	return events, nil
}

// GetAllDateEvents получает все события, сгруппированные по парам
func (p *PostgresDB) GetAllDateEvents(ctx context.Context) (map[string][]models.DateEvent, error) {
	query := `
		SELECT id, person1_name, person2_name, date, description, created_at
		FROM date_events
		ORDER BY person1_name, person2_name, date
	`

	var events []models.DateEvent
	err := p.db.SelectContext(ctx, &events, query)
	if err != nil {
		return nil, fmt.Errorf("failed to get all date events: %w", err)
	}

	// Группируем события по парам
	result := make(map[string][]models.DateEvent)
	for _, event := range events {
		key := p.generatePairKey(event.Person1Name, event.Person2Name)
		result[key] = append(result[key], event)
	}

	return result, nil
}

// GetDateEventByID получает событие по ID
func (p *PostgresDB) GetDateEventByID(ctx context.Context, id int) (*models.DateEvent, error) {
	query := `
		SELECT id, person1_name, person2_name, date, description, created_at
		FROM date_events
		WHERE id = $1
	`

	var event models.DateEvent
	err := p.db.GetContext(ctx, &event, query, id)
	if err != nil {
		return nil, fmt.Errorf("failed to get date event by id: %w", err)
	}

	return &event, nil
}

// generatePairKey генерирует ключ для пары людей
func (p *PostgresDB) generatePairKey(person1, person2 string) string {
	names := []string{person1, person2}
	sort.Strings(names)
	return strings.Join(names, "_")
}
