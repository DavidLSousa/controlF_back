package models

import (
	"controlF_back/internal/shared"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgconn"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	ID        uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time

	Name     string   `gorm:"type:varchar(255);not null"`
	Email    string   `gorm:"type:varchar(255);not null;unique"`
	Password string   `gorm:"type:varchar(255);not null"`
	Type     UserType `gorm:"not null"`

	CompanyID      *uuid.UUID      // Ponteiro para permitir valor nulo (NULL)
	Company        *Company        `gorm:"foreignKey:CompanyID"`
	Transactions   []Transaction   `gorm:"foreignKey:UserID"`
	Categories     []Category      `gorm:"foreignKey:UserID"`
	Summaries      []Summary       `gorm:"foreignKey:UserID"`
	PaymentMethods []PaymentMethod `gorm:"foreignKey:UserID"`
}

func (u *User) Save() error {
	if u.ID == uuid.Nil {
		u.ID = uuid.New()
	}

	// TODO: Criar junto as categorias basicas - todas com status inactive (cafeteria, jantar, transporte, contas de casa, investimentos) na mesma tx

	if err := DB.Create(&u).Error; err != nil {
		var pgErr *pgconn.PgError

		if errors.As(err, &pgErr) && pgErr.Code == shared.CodeErrorEmailAlreadyExists {
			return errors.New(shared.ErrMessage("email", shared.IsExists))
		}

		return fmt.Errorf("error creating user: %w", err)
	}

	return nil
}

func (u *User) Update(updates map[string]interface{}) error {
	if len(updates) == 0 {
		return nil
	}
	return DB.Model(u).Updates(updates).Error
}

func GetUserWithRoles(id uuid.UUID) (*User, error) {
	if id == uuid.Nil {
		return nil, gorm.ErrInvalidData
	}

	u := &User{}
	err := DB.Preload("RoleType").First(u, "id = ?", id).Error

	if err != nil {
		return nil, err
	}

	return u, nil
}

func GetUser(id uuid.UUID) (*User, error) {
	if id == uuid.Nil {
		return nil, gorm.ErrInvalidData
	}

	u := &User{}
	err := DB.First(u, "id = ?", id).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New(shared.ErrMessage("user", shared.NotFound))
		}
		return nil, fmt.Errorf("error getting user: %w", err)
	}

	return u, nil
}

// func ListUser(userId uuid.UUID, pagination *Pagination) ([]User, error) {
// 	if pagination == nil {
// 		pagination = DefaultPagination()
// 	}

// 	var users []User
// 	err := DB.
// 		Scopes(pagination.GetScope).
// 		Where("company_id = ?", userId).
// 		Find(&users).Error

// 	if err != nil {
// 		return nil, err
// 	}

// 	return users, nil
// }

func VerifyPassword(password, hashedPassword string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		if err == bcrypt.ErrMismatchedHashAndPassword {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}
