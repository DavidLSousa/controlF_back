package user

import (
	"controlF_back/internal/shared"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgconn"
	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{DB: db}
}

func (r *UserRepository) Create(u *User) error {
	if u.ID == uuid.Nil {
		u.ID = uuid.New()
	}

	// TODO: Criar junto as categorias basicas - todas com status inactive (cafeteria, jantar, transporte, contas de casa, investimentos) na mesma tx

	if err := r.DB.Create(&u).Error; err != nil {
		var pgErr *pgconn.PgError

		if errors.As(err, &pgErr) && pgErr.Code == shared.CodeErrorEmailAlreadyExists {
			return errors.New(shared.ErrMessage("email", shared.IsExists))
		}

		return fmt.Errorf("error creating user: %w", err)
	}

	return nil
}

func (r *UserRepository) Update(u *User, updates map[string]interface{}) error {
	if len(updates) == 0 {
		return nil
	}

	return r.DB.Model(u).Updates(updates).Error
}

func (r *UserRepository) UpdatePassword(u *User) error {
	return r.DB.Model(u).Update("password", u.Password).Error
}

func (r *UserRepository) Get(id uuid.UUID) (*User, error) {
	if id == uuid.Nil {
		return nil, gorm.ErrInvalidData
	}

	u := &User{}
	err := r.DB.First(u, "id = ?", id).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New(shared.ErrMessage("user", shared.NotFound))
		}
		return nil, fmt.Errorf("error getting user: %w", err)
	}

	return u, nil
}

// func (u *User) GetUserWithRoles(id uuid.UUID) (*User, error) {
// 	if id == uuid.Nil {
// 		return nil, gorm.ErrInvalidData
// 	}

// 	err := DB.Preload("RoleType").First(u, "id = ?", id).Error

// 	if err != nil {
// 		return nil, err
// 	}

// 	return u, nil
// }
