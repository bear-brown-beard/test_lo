package di

import (
	"user_service/internal/repository"
	"user_service/internal/service"
)

type ServiceContainer struct {
	UserService service.UserService
}

func NewServiceContainer(userRepo repository.UserRepository) *ServiceContainer {
	userService := service.NewUserService(userRepo)

	return &ServiceContainer{
		UserService: userService,
	}
}
