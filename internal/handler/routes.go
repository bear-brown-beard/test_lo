package handler

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func SetupRoutes(dateHandler *DateHandler) *chi.Mux {
	r := chi.NewRouter()

	// Добавляем middleware для логирования
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// Маршруты для работы с датами
	r.Route("/api", func(r chi.Router) {
		r.Route("/dates", func(r chi.Router) {
			r.Post("/", dateHandler.CreateDateEventHandler)       // POST /api/dates
			r.Get("/all", dateHandler.GetAllDateEventsHandler)    // GET /api/dates/all
			r.Delete("/{id}", dateHandler.DeleteDateEventHandler) // DELETE /api/dates/{id}
		})
	})

	return r
}
