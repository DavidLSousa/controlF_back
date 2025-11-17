package user

import (
	"controlF_back/internal/database"
	"controlF_back/internal/utils"
)

func InitUserService() *UserController {
	crypt := utils.NewBcrypt()
	repo := NewUserRepository(database.DB)
	service := NewUserService(crypt)

	useCases := NewUserUseCases(*repo, *service)

	handler := NewUserController(*useCases)

	return handler
}
