package auth

import "github.com/google/uuid"

type AuthUser struct {
	ID       uuid.UUID
	Email    string
	Password string
}

// TableName diz ao GORM qual tabela usar
func (AuthUser) TableName() string {
	return "users"
}
