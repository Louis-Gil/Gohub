package repository

import (
	"server/domain/models"
)

type UserRepository interface {
	Create(user *models.User) error
	GetByID(userID string) (*models.User, error)
}
