package handler

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

type Routes struct {
	taskHandler *TaskHandler
}

func BuildRoutes(taskHandler *TaskHandler) *Routes {
	return &Routes{
		taskHandler: taskHandler,
	}
}

func (r *Routes) SetupRoutes() http.Handler {
	router := chi.NewRouter()

	router.Route("/tasks", func(router chi.Router) {
		router.Post("/", r.taskHandler.CreateTaskHandler)
		router.Get("/", r.taskHandler.GetAllTasksHandler)
		router.Get("/{id}", r.taskHandler.GetTaskByIDHandler)
	})

	return router
}
