package database

import (
	"context"
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

// Database интерфейс для работы с базой данных
type Database interface {
	Connect(ctx context.Context) error
	Close() error
	Ping(ctx context.Context) error
	GetDB() *sqlx.DB
}

// PostgresDB реализация интерфейса Database для PostgreSQL
type PostgresDB struct {
	db *sqlx.DB
}

// Config конфигурация для подключения к PostgreSQL
type Config struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
	SSLMode  string
}

// NewPostgresDB создает новый экземпляр PostgreSQL базы данных
func NewPostgresDB(config Config) (*PostgresDB, error) {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
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

// GetDB возвращает соединение с базой данных для использования в репозитории
func (p *PostgresDB) GetDB() *sqlx.DB {
	return p.db
}
