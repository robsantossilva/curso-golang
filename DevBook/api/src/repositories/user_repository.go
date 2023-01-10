package repositories

import (
	"api/src/models"
	"database/sql"
)

type userRepository struct {
	DB *sql.DB
}

func NewUserRepository(db *sql.DB) *userRepository {
	return &userRepository{
		DB: db,
	}
}

func (r *userRepository) Create(user models.User) (uint64, error) {
	return 0, nil
}
