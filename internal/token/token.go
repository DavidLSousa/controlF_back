package token

import (
	"controlF_back/internal/utils"
	"errors"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/rs/zerolog/log"
)

func CreateAccessTokenTTL(meta TokenMeta) (accessToken string, err error) {
	meta.Type = AccessTokenType
	var expiresAt *jwt.NumericDate
	if meta.Ttl > 0 {
		expiresAt = jwt.NewNumericDate(meta.Exp)
	}

	claims := &JwtCustomClaims{
		Name:  meta.UserName,
		Roles: meta.Roles,
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   meta.UserId.String(),
			ExpiresAt: expiresAt,
			ID:        meta.Jti.String(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte(utils.GetEnv("JWT_KEY", "")))
	if err != nil {
		return "", err
	}

	return t, err
}

func CreateAccessToken(meta TokenMeta) (accessToken string, err error) {
	expiry, err := strconv.Atoi(os.Getenv("JWT_LIFESPAN"))
	if err != nil {
		return "", err
	}

	meta.SetTtl(time.Minute * time.Duration(expiry))
	return CreateAccessTokenTTL(meta)
}

func CreateRefreshToken(meta TokenMeta) (refreshToken string, err error) {
	meta.Type = RefreshTokenType
	expiry, err := strconv.Atoi(os.Getenv("JWT_REFRESH_LIFESPAN"))
	if err != nil {
		return "", err
	}
	meta.SetTtl(time.Hour * time.Duration(expiry))

	claims := &JwtCustomClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   meta.UserId.String(),
			ExpiresAt: jwt.NewNumericDate(meta.Exp),
			ID:        meta.Jti.String(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte(utils.GetEnv("JWT_KEY", "")))
	if err != nil {
		return "", err
	}

	return t, err
}

func SetRefreshTokenCookie(c *gin.Context, token string) error {
	expiry, err := strconv.Atoi(os.Getenv("JWT_REFRESH_LIFESPAN"))
	if err != nil {
		return err
	}
	maxAge := expiry * 60 * 60
	c.SetSameSite(http.SameSiteNoneMode)
	domain := c.Request.Host
	log.Debug().Str("domain", domain).Msg("SetRefreshTokenCookie")
	c.SetCookie("refresh_token", token, maxAge, "/", domain, true, true)

	return nil
}

func GetRefreshTokenCookie(c *gin.Context) (string, error) {
	return c.Cookie("refresh_token")
}

func IsAuthorized(requestToken string) (*JwtCustomClaims, error) {
	token, err := ExtractToken(requestToken)
	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("invalid token")
	}

	claims, err := ExtractClaims(token)
	if err != nil {
		return nil, errors.New("invalid token in claims")
	}

	return claims, nil
}

func ExtractToken(requestToken string) (*jwt.Token, error) {
	token, err := jwt.ParseWithClaims(requestToken, &JwtCustomClaims{}, func(token *jwt.Token) (interface{}, error) {

		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("método de assinatura inesperado: %v", token.Header["alg"])
		}

		return []byte(utils.GetEnv("JWT_KEY", "")), nil
	})

	if err != nil {
		return nil, err
	}

	return token, nil
}

func ExtractClaims(token *jwt.Token) (*JwtCustomClaims, error) {
	claims, ok := token.Claims.(*JwtCustomClaims)

	if !ok {
		return nil, errors.New("invalid token")
	}
	return claims, nil
}
