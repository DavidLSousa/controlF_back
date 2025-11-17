package auth

import (
	"controlF_back/internal/token"

	"github.com/google/uuid"
)

type AuthService struct{}

func NewAuthService() *AuthService {
	return &AuthService{}
}

// Precisa ser ajustado:
/*
Lógica de domínio ✅
	Definir roles do usuário
	Criar metadata do token
Mas depende de infraestrutura ❌
	token.CreateAccessToken() - geração de JWT (infra)
	uuid.New() - geração de UUID (infra)
*/

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

	// refreshToken, err := token.CreateRefreshToken(meta)
	// if err != nil {
	// 	return "", "", err
	// }

	return accessToken, nil
}

// func (d *AuthService) setRefreshTokenCookie(c *gin.Context, refreshToken string) {
// 	token.SetRefreshTokenCookie(c, refreshToken)
// }
