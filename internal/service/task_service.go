package service

import (
	"crypto/rand"
	"encoding/hex"
	"test_lo/internal/models"
	"time"
)

type TaskRepository interface {
	Create(task models.Task) error
	GetByID(id string) (models.Task, error)
	GetAll(status string) ([]models.Task, error)
}

type TaskService struct {
	repo TaskRepository
}

func NewTaskService(repo TaskRepository) *TaskService {
	return &TaskService{repo: repo}
}

func (s *TaskService) CreateTask(title, description string) (models.Task, error) {
	newTask := models.Task{
		ID:          generateID(),
		Title:       title,
		Description: description,
		Status:      "pending",
		CreatedAt:   time.Now(),
	}
	if err := s.repo.Create(newTask); err != nil {
		return models.Task{}, err
	}
	return newTask, nil
}

func generateID() string {
	bytes := make([]byte, 16)
	rand.Read(bytes)
	return hex.EncodeToString(bytes)
}

func (s *TaskService) GetTaskByID(id string) (models.Task, error) {
	return s.repo.GetByID(id)
}

func (s *TaskService) GetAllTasks(status string) ([]models.Task, error) {
	return s.repo.GetAll(status)
}
