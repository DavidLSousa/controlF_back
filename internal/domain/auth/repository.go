package auth

import (
	"controlF_back/internal/models"
	"errors"
)

type GormAuthRepository struct{}

func (r *GormAuthRepository) login(input LoginRequest) (*LoginResponse, error) {
	user := &models.User{}
	if err := models.DB.Where("email = ?", input.Email).First(user).Error; err != nil {
		return nil, errors.New("invalid credentials")
	}

	if ok, err := models.VerifyPassword(input.Password, user.Password); !ok || err != nil {
		return nil, errors.New("invalid credentials")
	}

	// TODO: Generate JWT token
	token := "dummy-jwt-token"

	return &LoginResponse{Token: token}, nil
}
