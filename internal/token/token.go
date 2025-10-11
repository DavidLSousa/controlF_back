package token

import (
	"context"
	"controlF_back/internal/kv"
	"controlF_back/internal/utils"
	"encoding/json"
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

	err = SaveToken(meta)
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
	err = SaveToken(meta)
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

func SaveToken(meta TokenMeta) error {
	key := meta.GetKey()
	value, err := json.Marshal(meta)
	if err != nil {
		return err
	}
	return kv.Client.Set(context.Background(), key, value, meta.Ttl).Err()
}

func ListTokens(tokenType TokenType, userId string) ([]TokenMeta, error) {
	key := fmt.Sprintf("%s%s:%s:*", keyPrefix, userId, tokenType)
	keys, err := kv.Client.Keys(context.Background(), key).Result()
	if err != nil {
		return nil, err
	}

	tokens := make([]TokenMeta, 0)
	for _, key := range keys {
		value, err := kv.Client.Get(context.Background(), key).Bytes()
		var result TokenMeta
		if err != nil {
			return nil, err
		}
		err = json.Unmarshal(value, &result)
		if err != nil {
			log.Error().Err(err).Msg("Error parsing token")
		}
		tokens = append(tokens, result)
	}

	return tokens, nil
}

func GetToken(tokenType TokenType, userId string, jti string) (TokenMeta, error) {
	key := fmt.Sprintf("%s%s:%s:%s", keyPrefix, userId, tokenType, jti)
	value, err := kv.Client.Get(context.Background(), key).Bytes()
	var result TokenMeta
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(value, &result)
	if err != nil {
		log.Error().Err(err).Msg("Error parsing token")
	}
	return result, nil
}

func CleanAll(pattern string) error {
	keys, err := kv.Client.Keys(context.Background(), pattern).Result()
	if err != nil {
		return err
	}

	for _, key := range keys {
		err := Clean(key)
		if err != nil {
			log.Error().Err(err).Msg("Error cleaning token")
		}
	}

	return kv.Client.Del(context.Background(), pattern).Err()
}

func Clean(key string) error {
	return kv.Client.Del(context.Background(), key).Err()
}
