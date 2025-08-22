package di

import (
	"user_service/internal/repository"

	"github.com/jmoiron/sqlx"
)

// RepositoryContainer контейнер для репозиториев
type RepositoryContainer struct {
	UserRepository repository.UserRepository
}

func NewRepositoryContainer(db *sqlx.DB) *RepositoryContainer {
	userRepo := repository.NewUserRepository(db)

	return &RepositoryContainer{
		UserRepository: userRepo,
	}
}
