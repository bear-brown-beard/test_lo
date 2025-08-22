package handler

import (
	"encoding/json"
	"net/http"
	"strconv"
	"together_service/internal/models"

	"github.com/go-chi/chi/v5"
)

type DateService interface {
	SaveDateEvent(event models.DateEvent) error
	GetAllDateEvents() (map[string][]models.DateEvent, error)
	DeleteDateEvent(id int) error
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

// DeleteDateEventHandler обрабатывает DELETE запрос для удаления события
func (h *DateHandler) DeleteDateEventHandler(w http.ResponseWriter, r *http.Request) {
	// Получаем ID из URL path
	idStr := chi.URLParam(r, "id")
	if idStr == "" {
		http.Error(w, "Необходимо указать ID события", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Неверный формат ID", http.StatusBadRequest)
		return
	}

	// Удаляем событие
	if err := h.service.DeleteDateEvent(id); err != nil {
		http.Error(w, "Ошибка удаления события", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Событие успешно удалено",
	})
}
