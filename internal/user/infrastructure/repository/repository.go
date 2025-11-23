package repository

import (
	"errors"
	"log"

	"gorm.io/gorm"

	"avito_internship_PR/internal/db/models"
	errorPkg "avito_internship_PR/internal/error"
	"avito_internship_PR/internal/user/application/dto/input"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) GetByID(userID string) (*models.User, error) {
	var user models.User
	if err := r.db.Where("id = ?", userID).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) CreateOrUpdate(userDto input.UserInput) (string, error) {
	var existingUser models.User
	user := models.User{
		ID:          userDto.UserId,
		Username:    userDto.UserName,
		IsActive:    userDto.IsActive,
		TeamMembers: userDto.TeamMembers,
	}
	err := r.db.Where("id = ?", userDto.UserId).First(&existingUser).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		r.db.Create(&user)
		return user.ID, nil
	} else if err != nil {
		return "", err
	}

	err = r.db.Save(&user).Error
	if err != nil {
		return "", err
	}

	return user.ID, nil
}

func (r *UserRepository) SetUserActiveStatus(userID string, isActive bool) (*models.User, error) {
	var user models.User
	if err := r.db.Where("id = ?", userID).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errorPkg.NotFound
		}
		log.Printf("ERROR IN USER REPO: %s", err)
		return nil, err
	}
	err := r.db.Model(&user).Update("is_active", isActive).Error
	if err != nil {
		log.Printf("ERROR IN USER REPO: %s", err)
		return nil, err
	}
	err = r.db.Preload("TeamMembers.Team").Where("id = ?", userID).First(&user).Error
	if err != nil {
		return nil, errorPkg.NotFound
	}
	return &user, nil
}
