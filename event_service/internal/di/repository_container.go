package di

import (
	"event_service/internal/repository"

	"github.com/jmoiron/sqlx"
)

// RepositoryContainer контейнер для репозиториев
type RepositoryContainer struct {
	DateRepository *repository.DateRepository
}

func NewRepositoryContainer(db *sqlx.DB) *RepositoryContainer {
	dateRepo := repository.NewDateRepository(db)

	return &RepositoryContainer{
		DateRepository: dateRepo,
	}
}
