package user

import (
	"controlF_back/internal/models"
	"controlF_back/internal/utils"
	"fmt"

	"github.com/google/uuid"
)

type UserService struct {
	UserRepository UserRepositoryInterface
	crypt          utils.Crypt
}

func NewUserService(repo UserRepositoryInterface, crypt utils.Crypt) *UserService {
	return &UserService{
		UserRepository: repo,
		crypt:          crypt,
	}
}

func (s *UserService) Create(input UserRegister) (*UserDto, error) {
	hashedPassword, err := s.crypt.Hash(input.Password)
	if err != nil {
		return nil, fmt.Errorf("error generating password: %w", err)
	}

	user := &models.User{
		Name:     input.Name,
		Email:    input.Email,
		Password: hashedPassword,
		Type:     models.UserTypePersonal,
	}

	if err := s.UserRepository.Create(user); err != nil {
		return nil, err
	}

	dto := MapUserResposeDto(user)
	return &dto, nil
}

func (s *UserService) Get(userId uuid.UUID) (*UserDto, error) {
	user, err := s.UserRepository.Get(userId)
	if err != nil {
		return nil, err
	}

	dto := MapUserResposeDto(user)
	return &dto, nil
}

func (s *UserService) Update(userId uuid.UUID, input UserUpdate) (*UserDto, error) {
	user, err := s.UserRepository.Get(userId)
	if err != nil {
		return nil, err
	}

	updates := make(map[string]interface{})
	if input.Name != nil {
		updates["name"] = input.Name
	}
	if input.Type != nil {
		updates["type"] = input.Type
	}
	if input.CompanyID != nil {
		updates["company_id"] = input.CompanyID
	}

	if err := s.UserRepository.Update(user, updates); err != nil {
		return nil, err
	}

	dto := MapUserResposeDto(user)
	return &dto, nil
}

func (s *UserService) UpdatePassword(userId uuid.UUID, input UserUpdatePassword) error {
	user, err := s.UserRepository.Get(userId)
	if err != nil {
		return err
	}

	if err := s.crypt.Check(input.OldPassword, user.Password); err != nil {
		return fmt.Errorf("old password incorrect")
	}

	if input.OldPassword == input.NewPassword {
		return fmt.Errorf("old password and new password cannot be the same")
	}

	hashedPassword, err := s.crypt.Hash(input.NewPassword)
	if err != nil {
		return fmt.Errorf("error generating new password hash: %w", err)
	}

	user.Password = hashedPassword

	if err := s.UserRepository.UpdatePassword(user); err != nil {
		return err
	}

	return nil
}
