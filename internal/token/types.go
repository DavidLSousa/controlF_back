package token

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

const keyPrefix = "jwt:"

type RoleType string

const (
	RoleTypeUser          RoleType = "user"
	RoleTypeComplete      RoleType = "complete"
	RoleTypeVerified      RoleType = "verified"
	RoleTypeAdministrator RoleType = "administrator"
	RoleTypeBetaTester    RoleType = "beta_tester"
)

type JwtCustomClaims struct {
	Name  string     `json:"name,omitempty"`
	Roles []RoleType `json:"roles,omitempty"`
	jwt.RegisteredClaims
}

type TokenType string

const (
	AccessTokenType  TokenType = "access_token"
	RefreshTokenType TokenType = "refresh_token"
)

type TokenMeta struct {
	Type      TokenType     `json:"type"`
	UserAgent string        `json:"user_agent"`
	Ip        string        `json:"ip"`
	Name      string        `json:"name"`
	UserId    uuid.UUID     `json:"user"`
	UserName  string        `json:"user_name"`
	Roles     []RoleType    `json:"roles"`
	Jti       uuid.UUID     `json:"jti"`
	Exp       time.Time     `json:"exp"`
	Ttl       time.Duration `json:"ttl"`
}

func (meta *TokenMeta) SetTtl(ttl time.Duration) {
	meta.Ttl = ttl
	if meta.Ttl > 0 {
		meta.Exp = time.Now().Add(ttl)
	}
}

func (meta *TokenMeta) GetKey() string {
	return fmt.Sprintf("%s%s:%s:%s", keyPrefix, meta.UserId, meta.Type, meta.Jti)
}
