package auth

type AuthService struct {
	AuthRepository AuthRepositoryInterface
}

func NewAuthService(repo AuthRepositoryInterface) *AuthService {
	return &AuthService{
		AuthRepository: repo,
	}
}

func (s *AuthService) login(input LoginRequest) *LoginResponse {
	response, err := s.AuthRepository.login(input)
	if err != nil {
		return nil
	}
	return response
}

func (s *AuthService) logout() *LoginResponse {
	return nil
}
