package models

import (
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID        uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time

	Name     string   `gorm:"type:varchar(255);not null"`
	Email    string   `gorm:"type:varchar(255);not null;unique"`
	Password string   `gorm:"type:varchar(255);not null"`
	Type     UserType `gorm:"not null"`

	/* Relacionamento com Empresa (opcional, apenas para usuários de empresa) */
	CompanyID      *uuid.UUID      // Usamos ponteiro para permitir valor nulo (NULL)
	Company        *Company        `gorm:"foreignKey:CompanyID"`
	Transactions   []Transaction   `gorm:"foreignKey:UserID"`
	Categories     []Category      `gorm:"foreignKey:UserID"`
	Summaries      []Summary       `gorm:"foreignKey:UserID"`
	PaymentMethods []PaymentMethod `gorm:"foreignKey:UserID"`
}

func (u *User) Save() error {
	err := DB.Create(&u).Error
	if err != nil {
		return err
	}
	// Criar junto as categoras basicas - todas com status inactive (cafeteria, jantar, transporte, contas de casa, investimentos)
	return nil
}

func (u *User) Update(updates map[string]interface{}) error {
	return DB.
		Model(&u).
		Updates(updates).Error
}

func GetUserWithRoles(id uuid.UUID) (*User, error) {
	var err error
	u := User{}
	err = DB.
		Preload("RoleType").
		Take(&u, "id = ?", id).Error

	if err != nil {
		return nil, err
	}

	return &u, nil
}

func GetUser(id uuid.UUID) (*User, error) {
	var err error
	f := User{}
	err = DB.Take(&f, "id = ?", id).Error

	if err != nil {
		return nil, err
	}

	return &f, nil
}

func ListUser(userId uuid.UUID, pagination *Pagination) ([]User, error) {
	var err error
	var u []User

	if pagination == nil {
		pagination = DefaultPagination()
	}
	err = DB.
		Scopes(pagination.GetScope).
		Find(&u, "user_id = ?", userId).Error

	if err != nil {
		return nil, err
	}

	return u, nil
}

func VerifyPassword(password, hashedPassword string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		if err == bcrypt.ErrMismatchedHashAndPassword {
			return false, nil
		} else {
			return false, err
		}
	}

	return true, nil
}
