package auth

import (
	"gorm.io/gorm"
)

type AuthRepository struct {
	DB *gorm.DB
}

func NewAuthRepository(db *gorm.DB) *AuthRepository {
	return &AuthRepository{DB: db}
}

func (r *AuthRepository) GetUserByEmail(email string) (*AuthUser, error) {
	user := &AuthUser{}
	err := r.DB.First(user, "email = ?", email).Error
	return user, err
}
