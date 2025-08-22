package service

import (
	"context"
	"time"
	"together_service/internal/models"
	"together_service/internal/repository"
)

type DateService struct {
	repo repository.DateRepositoryInterface
}

func NewDateService(repo repository.DateRepositoryInterface) *DateService {
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

func (s *DateService) GetAllDateEvents() (map[string][]models.DateEvent, error) {
	ctx := context.Background()
	return s.repo.GetAll(ctx)
}

func (s *DateService) DeleteDateEvent(id int) error {
	ctx := context.Background()
	return s.repo.Delete(ctx, id)
}
