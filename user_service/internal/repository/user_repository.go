package repository

import (
	"user_service/internal/models"

	"github.com/jmoiron/sqlx"
)

type UserRepository interface {
	Create(user *models.User) (int64, error)
	GetByID(id int64) (*models.User, error)
	GetAll() ([]*models.User, error)
	Delete(id int64) error
}

type userRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) Create(user *models.User) (int64, error) {
	query := `INSERT INTO users (first_name, last_name) VALUES ($1, $2) RETURNING id`
	var id int64
	err := r.db.QueryRow(query, user.FirstName, user.LastName).Scan(&id)
	return id, err
}

func (r *userRepository) GetByID(id int64) (*models.User, error) {
	query := `SELECT id, first_name, last_name FROM users WHERE id = $1`

	var user models.User
	err := r.db.Get(&user, query, id)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *userRepository) GetAll() ([]*models.User, error) {
	query := `SELECT id, first_name, last_name FROM users`

	var users []*models.User
	err := r.db.Select(&users, query)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (r *userRepository) Delete(id int64) error {
	query := `DELETE FROM users WHERE id = $1`
	_, err := r.db.Exec(query, id)
	return err
}
