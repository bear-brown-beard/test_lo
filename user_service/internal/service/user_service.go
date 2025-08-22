package service

import (
	"user_service/internal/models"
	"user_service/internal/repository"
)

type UserService interface {
	CreateUser(user *models.User) (int64, error)
	GetUserByID(id int64) (*models.User, error)
	GetAllUsers() ([]*models.User, error)
	DeleteUser(id int64) error
}

type userService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{repo: repo}
}

func (s *userService) CreateUser(user *models.User) (int64, error) {
	return s.repo.Create(user)
}

func (s *userService) GetUserByID(id int64) (*models.User, error) {
	return s.repo.GetByID(id)
}

func (s *userService) GetAllUsers() ([]*models.User, error) {
	return s.repo.GetAll()
}

func (s *userService) DeleteUser(id int64) error {
	return s.repo.Delete(id)
}
