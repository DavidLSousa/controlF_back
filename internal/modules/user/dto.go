package user

import (
	"github.com/google/uuid"
)

type UserRegisterDto struct {
	Name            string `json:"name" binding:"required,min=5,max=100"`
	Email           string `json:"email" binding:"required,email"`
	Password        string `json:"password" binding:"required,password"`
	PasswordConfirm string `json:"passwordConfirm" binding:"required,password,eqfield=Password"`
}

type UserUpdateDto struct {
	Name      *string    `json:"name,omitempty"`
	Type      *string    `json:"type,omitempty"`
	CompanyID *uuid.UUID `json:"companyId,omitempty"`
}

type UserUpdatePasswordDto struct {
	NewPassword        string `json:"newPassword" binding:"required,password"`
	NewPasswordConfirm string `json:"newPasswordConfirm" binding:"required,password,eqfield=NewPassword"`
	OldPassword        string `json:"oldPassword" binding:"required,password"`
}

type UserDto struct {
	Id        string   `json:"id,omitempty"`
	Name      string   `json:"name"`
	Email     string   `json:"email,omitempty"`
	Type      UserType `json:"user_type"`
	CompanyId string   `json:"company_id,omitempty"`
}

func MapUserResposeDto(user *User) UserDto {
	dto := UserDto{
		Name:  user.Name,
		Email: user.Email,
		Type:  user.Type,
	}
	if user.CompanyID != nil {
		dto.CompanyId = user.CompanyID.String()
	}
	return dto
}

func NewUserDto(user *User) UserDto {
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
