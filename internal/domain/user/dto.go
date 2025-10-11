package user

import (
	"controlF_back/internal/models"

	"github.com/google/uuid"
)

type UserRegister struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

type UserUpdate struct {
	Name      *string    `json:"name,omitempty"`
	Email     *string    `json:"email,omitempty"`
	Type      *string    `json:"type,omitempty"`
	CompanyID *uuid.UUID `json:"companyId,omitempty"`
}

type UserUpdatePassword struct {
	NewPassword        string `json:"newPassword" binding:"required,min=6"`
	NewPasswordConfirm string `json:"newPasswordConfirm" binding:"required,min=6"`
	OldPassword        string `json:"oldPassword" binding:"required,min=6"`
}

type UserDto struct {
	Id        string          `json:"id"`
	Name      string          `json:"name"`
	Email     string          `json:"email"`
	Type      models.UserType `json:"user_type"`
	CompanyId string          `json:"company_id,omitempty"`
}

func NewUserDto(user *models.User) UserDto {
	dto := UserDto{
		Id:    user.ID.String(),
		Name:  user.Name,
		Email: user.Email,
		Type:  user.Type,
	}
	if user.CompanyID != nil {
		dto.CompanyId = user.CompanyID.String()
	}
	return dto
}
