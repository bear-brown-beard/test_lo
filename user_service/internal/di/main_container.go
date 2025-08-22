package di

import (
	"user_service/internal/config"
)

type MainContainer struct {
	DatabaseContainer   *DatabaseContainer
	RepositoryContainer *RepositoryContainer
	ServiceContainer    *ServiceContainer
}

func NewMainContainer() (*MainContainer, error) {
	dbContainer, err := NewDatabaseContainer()
	if err != nil {
		return nil, err
	}

	repoContainer := NewRepositoryContainer(dbContainer.DB)
	serviceContainer := NewServiceContainer(repoContainer.UserRepository)

	return &MainContainer{
		DatabaseContainer:   dbContainer,
		RepositoryContainer: repoContainer,
		ServiceContainer:    serviceContainer,
	}, nil
}

func (c *MainContainer) Close() error {
	return c.DatabaseContainer.Close()
}

func GetConfig() *config.Config {
	return config.NewConfig()
}
