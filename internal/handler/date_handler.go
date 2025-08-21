package handler

import (
	"encoding/json"
	"net/http"
	"together_service/internal/models"
)

type DateService interface {
	SaveDateEvent(event models.DateEvent) error
	GetDateEvents(person1Name, person2Name string) ([]models.DateEvent, error)
	GetAllDateEvents() (map[string][]models.DateEvent, error)
}

type DateHandler struct {
	service DateService
}

func NewDateHandler(service DateService) *DateHandler {
	return &DateHandler{
		service: service,
	}
}

// CreateDateEventHandler обрабатывает POST запрос для создания нового события
func (h *DateHandler) CreateDateEventHandler(w http.ResponseWriter, r *http.Request) {
	var event models.DateEvent
	if err := json.NewDecoder(r.Body).Decode(&event); err != nil {
		http.Error(w, "Неверный формат JSON", http.StatusBadRequest)
		return
	}

	// Проверяем обязательные поля
	if event.Person1Name == "" || event.Person2Name == "" || event.Description == "" {
		http.Error(w, "Все поля обязательны для заполнения", http.StatusBadRequest)
		return
	}

	// Сохраняем событие
	if err := h.service.SaveDateEvent(event); err != nil {
		http.Error(w, "Ошибка сохранения события", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Событие успешно сохранено",
	})
}

// GetDateEventsHandler обрабатывает GET запрос для получения событий пары
func (h *DateHandler) GetDateEventsHandler(w http.ResponseWriter, r *http.Request) {
	person1Name := r.URL.Query().Get("person1")
	person2Name := r.URL.Query().Get("person2")

	if person1Name == "" || person2Name == "" {
		http.Error(w, "Необходимо указать имена обоих людей", http.StatusBadRequest)
		return
	}

	events, err := h.service.GetDateEvents(person1Name, person2Name)
	if err != nil {
		http.Error(w, "Ошибка получения событий", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(events)
}

// GetAllDateEventsHandler обрабатывает GET запрос для получения всех событий
func (h *DateHandler) GetAllDateEventsHandler(w http.ResponseWriter, r *http.Request) {
	events, err := h.service.GetAllDateEvents()
	if err != nil {
		http.Error(w, "Ошибка получения всех событий", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(events)
}
