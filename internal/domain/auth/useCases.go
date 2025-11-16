package auth

import (
	"controlF_back/internal/utils"
	"fmt"
)

type AuthUseCases struct {
	repo    AuthRepository
	service AuthService
	crypt   utils.Crypt
}

func NewAuthUseCases(
	repo AuthRepository,
	service AuthService,
	crypt utils.Crypt) *AuthUseCases {
	return &AuthUseCases{
		crypt:   crypt,
		service: service,
		repo:    repo,
	}
}

func (s *AuthUseCases) Login(input LoginRequestDto) (*LoginResponseDto, error) {
	user, err := s.repo.GetUserByEmail(input.Email)
	if err != nil {
		return nil, fmt.Errorf("error getting user: %w", err)
	}

	if err := s.crypt.Check(input.Password, user.Password); err != nil {
		return nil, fmt.Errorf("error checking password: invalid credentials")
	}

	accessToken, err := s.service.generateAuthTokens(*user)
	if err != nil {
		return nil, fmt.Errorf("error generating tokens: %w", err)
	}

	return &LoginResponseDto{Token: accessToken}, nil
}

func (s *AuthUseCases) Logout() *LoginResponseDto {
	return nil
}
