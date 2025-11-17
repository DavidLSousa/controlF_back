package user

import (
	"controlF_back/internal/utils"
	"fmt"
)

type UserService struct {
	crypt utils.Crypt
}

func NewUserService(crypt utils.Crypt) *UserService {
	return &UserService{crypt: crypt}
}

func (s *UserService) NewUser(name, email, password string) (*User, error) {
	hashedPassword, err := s.HashPassword(password)
	if err != nil {
		return nil, err
	}

	user := &User{
		Name:     name,
		Email:    email,
		Password: hashedPassword,
		Type:     UserTypePersonal,
	}

	return user, nil
}

func (s *UserService) HashPassword(password string) (string, error) {
	hashedPassword, err := s.crypt.Hash(password)
	if err != nil {
		return "", fmt.Errorf("error generating password: %w", err)
	}

	return hashedPassword, nil
}

func (s *UserService) CheckPassword(oldPassword, hashedPassword string) error {
	return s.crypt.Check(oldPassword, hashedPassword)
}

func (s *UserService) ValidationUpdatePassword(user *User, oldPass, newPass, newPassConfirm string) error {
	if err := s.CheckPassword(oldPass, user.Password); err != nil {
		return fmt.Errorf("old password incorrect")
	}

	if oldPass == newPass {
		return fmt.Errorf("old password and new password cannot be the same")
	}

	hashedNewPassword, err := s.HashPassword(newPass)
	if err != nil {
		return err
	}

	user.Password = hashedNewPassword
	return nil
}

// validaçZoes para update
