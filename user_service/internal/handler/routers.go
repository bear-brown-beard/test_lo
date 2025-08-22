package handler

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func SetupRoutes(userHandler *UserHandler) *chi.Mux {
	r := chi.NewRouter()

	// Добавляем middleware для логирования
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// Маршруты для работы с пользователями
	r.Route("/api", func(r chi.Router) {
		r.Route("/users", func(r chi.Router) {
			r.Post("/", userHandler.CreateUser)       // POST /api/users
			r.Get("/", userHandler.GetAllUsers)       // GET /api/users
			r.Get("/{id}", userHandler.GetUserByID)   // GET /api/users/{id}
			r.Delete("/{id}", userHandler.DeleteUser) // DELETE /api/users/{id}
		})
	})

	return r
}
