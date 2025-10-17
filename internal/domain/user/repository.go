package user

import (
	"controlF_back/internal/models"

	"github.com/google/uuid"
)

type GormUserRepository struct{}

func (r *GormUserRepository) Create(user *models.User) error {
	return user.Save()
}

func (r *GormUserRepository) Get(userId uuid.UUID) (*models.User, error) {
	return models.GetUser(userId)
}

func (r *GormUserRepository) Update(user *models.User, updates map[string]interface{}) error {
	return user.Update(updates)
}

func (r *GormUserRepository) UpdatePassword(user *models.User) error {
	return user.Save()
}
