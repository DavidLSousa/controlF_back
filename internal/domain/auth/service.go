package auth

type AuthService struct {
}

func NewAuthService() *AuthService {
	return &AuthService{}
}

// Garante que AuthService implementa IAuthService
var _ IAuthService = (*AuthService)(nil)

func (s *AuthService) login(input LoginRequest) *LoginResponse {
	return nil
}

func (s *AuthService) logout() *LoginResponse {
	return nil
}
