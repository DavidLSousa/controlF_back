package auth

import (
	"controlF_back/internal/database"
	"controlF_back/internal/utils"
)

// Initializer
func InitAuthService() *AuthController {
	crypt := utils.NewBcrypt()
	repo := NewAuthRepository(database.DB)
	service := NewAuthService()

	useCases := NewAuthUseCases(*repo, *service, crypt)

	controller := NewAuthHandler(*useCases)

	return controller
}
