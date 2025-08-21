package di

import (
	"together_service/internal/repository"
	"together_service/internal/service"
)

// ServiceContainer контейнер для сервисов
type ServiceContainer struct {
	DateService *service.DateService
}

// NewServiceContainer создает новый контейнер для сервисов
func NewServiceContainer(dateRepo *repository.DateRepository) *ServiceContainer {
	// Создаем сервис дат
	dateService := service.NewDateService(dateRepo)

	return &ServiceContainer{
		DateService: dateService,
	}
}
