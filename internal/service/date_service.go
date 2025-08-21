package service

import (
	"context"
	"sort"
	"time"
	"together_service/internal/models"
	"together_service/internal/repository"
)

type DateService struct {
	repo *repository.DateRepository
}

func NewDateService(repo *repository.DateRepository) *DateService {
	return &DateService{
		repo: repo,
	}
}

func (s *DateService) SaveDateEvent(event models.DateEvent) error {
	ctx := context.Background()

	// Устанавливаем время создания
	event.CreatedAt = time.Now()

	// Сохраняем в базу данных
	return s.repo.Create(ctx, event)
}

func (s *DateService) GetDateEvents(person1Name, person2Name string) ([]models.DateEvent, error) {
	ctx := context.Background()
	return s.repo.GetByPair(ctx, person1Name, person2Name)
}

func (s *DateService) GetAllDateEvents() (map[string][]models.DateEvent, error) {
	ctx := context.Background()
	return s.repo.GetAll(ctx)
}

func (s *DateService) GetDateEventByID(id int) (*models.DateEvent, error) {
	ctx := context.Background()
	return s.repo.GetByID(ctx, id)
}

// Вспомогательные методы для совместимости с существующим кодом
func (s *DateService) GetSortedDateEvents(person1Name, person2Name string) ([]models.DateEvent, error) {
	events, err := s.GetDateEvents(person1Name, person2Name)
	if err != nil {
		return nil, err
	}

	// Сортируем по дате (хотя база данных уже возвращает отсортированные)
	sort.Slice(events, func(i, j int) bool {
		return events[i].Date.Before(events[j].Date)
	})

	return events, nil
}
