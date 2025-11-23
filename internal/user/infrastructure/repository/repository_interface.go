package repository

import (
	"avito_internship_PR/internal/db/models"
	"avito_internship_PR/internal/user/application/dto/input"
)

type IUserRepository interface {
	GetByID(userID string) (*models.User, error)
	CreateOrUpdate(userDto input.UserInput) (string, error)
	SetUserActiveStatus(userID string, isActive bool) (*models.User, error)
}
