package service

import (
	"context"
	"event_service/internal/models"
	"event_service/internal/repository"
	"time"
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

func (s *DateService) GetAllDateEvents() ([]models.DateEvent, error) {
	ctx := context.Background()
	return s.repo.GetAll(ctx)
}

func (s *DateService) DeleteDateEvent(id int) error {
	ctx := context.Background()
	return s.repo.Delete(ctx, id)
}
