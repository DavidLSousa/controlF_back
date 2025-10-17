package user

import (
	"controlF_back/internal/models"

	"github.com/google/uuid"
)

type UserRepositoryInterface interface {
	Create(user *models.User) error
	Get(userId uuid.UUID) (*models.User, error)
	Update(user *models.User, updates map[string]interface{}) error
	UpdatePassword(user *models.User) error
}

func NewUserRepository() UserRepositoryInterface {
	return &GormUserRepository{}
}

func InitUserService() *UserController {
	repo := NewUserRepository()
	service := NewUserService(repo)
	controller := NewUserController(*service)

	return controller
}
