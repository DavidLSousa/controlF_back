package auth

type AuthRepositoryInterface interface {
	login(input LoginRequest) (*LoginResponse, error)
	register(input RegisterRequest) error
}

func NewAuthRepository() AuthRepositoryInterface {
	return &GormAuthRepository{}
}

func InitAuthService() *AuthController {
	repo := NewAuthRepository()
	service := NewAuthService(repo)
	controller := NewAuthController(*service)

	return controller
}
