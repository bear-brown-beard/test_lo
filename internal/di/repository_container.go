package di

import (
	"together_service/internal/database"
	"together_service/internal/repository"
)

// RepositoryContainer контейнер для репозиториев
type RepositoryContainer struct {
	DateRepository *repository.DateRepository
}

// NewRepositoryContainer создает новый контейнер для репозиториев
func NewRepositoryContainer(db database.Database) *RepositoryContainer {
	// Создаем репозиторий дат
	dateRepo := repository.NewDateRepository(db)

	return &RepositoryContainer{
		DateRepository: dateRepo,
	}
}
