package user

import (
	"time"

	"github.com/google/uuid"
)

type UserType uint8

const (
	UserTypePersonal      UserType = 1 // Usuário de conta pessoal
	UserTypeCompanyAdmin  UserType = 2 // Usuário admin de uma empresa
	UserTypeCompanyMember UserType = 3 // Usuário membro de uma empresa
)

type User struct {
	ID        uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time

	Name     string   `gorm:"type:varchar(255);not null"`
	Email    string   `gorm:"type:varchar(255);not null;unique"`
	Password string   `gorm:"type:varchar(255);not null"`
	Type     UserType `gorm:"not null"`

	CompanyID *uuid.UUID `gorm:"type:uuid;index"` // Ponteiro para permitir valor nulo (NULL)
	// []Transactions
	// []Categories
	// []Summaries
	// []PaymentMethods
}
