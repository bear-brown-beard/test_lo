package di

import (
	"together_service/internal/config"
)

// MainContainer основной контейнер, объединяющий все слои
type MainContainer struct {
	DatabaseContainer   *DatabaseContainer
	RepositoryContainer *RepositoryContainer
	ServiceContainer    *ServiceContainer
}

// NewMainContainer создает новый основной контейнер
func NewMainContainer() (*MainContainer, error) {
	// Создаем контейнер базы данных
	dbContainer, err := NewDatabaseContainer()
	if err != nil {
		return nil, err
	}

	// Создаем контейнер репозиториев
	repoContainer := NewRepositoryContainer(dbContainer.DB)

	// Создаем контейнер сервисов
	serviceContainer := NewServiceContainer(repoContainer.DateRepository)

	return &MainContainer{
		DatabaseContainer:   dbContainer,
		RepositoryContainer: repoContainer,
		ServiceContainer:    serviceContainer,
	}, nil
}

// Close закрывает основной контейнер
func (c *MainContainer) Close() error {
	return c.DatabaseContainer.Close()
}

// GetConfig возвращает конфигурацию
func GetConfig() *config.Config {
	return config.Load()
}
