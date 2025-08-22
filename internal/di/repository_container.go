package di

import (
	"together_service/internal/database"
	"together_service/internal/repository"
)

// RepositoryContainer контейнер для репозиториев
type RepositoryContainer struct {
	DateRepository *repository.DateRepository
}

func NewRepositoryContainer(db database.Database) *RepositoryContainer {
	dateRepo := repository.NewDateRepository(db)

	return &RepositoryContainer{
		DateRepository: dateRepo,
	}
}
