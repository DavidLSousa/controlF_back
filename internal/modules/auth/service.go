package auth

import (
	"controlF_back/internal/token"

	"github.com/google/uuid"
)

type AuthService struct{}

func NewAuthService() *AuthService {
	return &AuthService{}
}

func (d *AuthService) generateAuthTokens(authUser AuthUser) (string, error) {
	jti := uuid.New()
	roles := []token.RoleType{token.RoleTypeUser}

	meta := token.TokenMeta{
		UserId:   authUser.ID,
		UserName: authUser.Email,
		Roles:    roles,
		Jti:      jti,
	}

	accessToken, err := token.CreateAccessToken(meta)
	if err != nil {
		return "", err
	}

	return accessToken, nil
}
