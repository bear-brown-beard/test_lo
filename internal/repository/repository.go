package repository

import (
	"fmt"
	"sync"
	"test_lo/internal/models"
)

type TaskRepository interface {
	Create(task models.Task) error
	GetByID(id string) (models.Task, error)
	GetAll(status string) ([]models.Task, error)
}

type InMemoryRepository struct {
	tasks map[string]models.Task
	mu    sync.RWMutex
}

func NewInMemoryRepository() *InMemoryRepository {
	return &InMemoryRepository{
		tasks: make(map[string]models.Task),
	}
}

func (r *InMemoryRepository) Create(task models.Task) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.tasks[task.ID]; exists {
		return fmt.Errorf("task with ID %s already exists", task.ID)
	}
	r.tasks[task.ID] = task
	return nil
}

func (r *InMemoryRepository) GetByID(id string) (models.Task, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	task, exists := r.tasks[id]
	if !exists {
		return models.Task{}, fmt.Errorf("task with ID %s not found", id)
	}
	return task, nil
}

func (r *InMemoryRepository) GetAll(status string) ([]models.Task, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	var result []models.Task
	for _, task := range r.tasks {
		if status == "" || task.Status == status {
			result = append(result, task)
		}
	}
	return result, nil
}
