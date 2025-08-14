package di

import (
	"test_lo/internal/handler"
	"test_lo/internal/logger"
	"test_lo/internal/repository"
	"test_lo/internal/service"
)

type Container struct {
	Logger        *logger.Logger
	TaskHandler   *handler.TaskHandler
	Routes        *handler.Routes
	ServerService *service.ServerService
}

func CreateContainer() (*Container, error) {

	logger := logger.NewLogger()
	repo := repository.NewInMemoryRepository()
	taskService := service.NewTaskService(repo)
	taskHandler := handler.MakeTaskHandler(taskService, logger)
	routes := handler.BuildRoutes(taskHandler)
	serverService := service.BuildServerService(logger, routes)

	return &Container{
		Logger:        logger,
		TaskHandler:   taskHandler,
		Routes:        routes,
		ServerService: serverService,
	}, nil
}
