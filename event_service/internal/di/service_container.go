package di

import (
	"event_service/internal/repository"
	"event_service/internal/service"
)

type ServiceContainer struct {
	DateService *service.DateService
}

func NewServiceContainer(dateRepo repository.DateRepositoryInterface) *ServiceContainer {
	dateService := service.NewDateService(dateRepo)

	return &ServiceContainer{
		DateService: dateService,
	}
}
