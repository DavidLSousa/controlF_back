package user

import (
	"github.com/google/uuid"
)

type UserUseCases struct {
	repo    UserRepository
	service UserService
}

func NewUserUseCases(repo UserRepository, service UserService) *UserUseCases {
	return &UserUseCases{
		repo:    repo,
		service: service,
	}
}

func (u *UserUseCases) Create(input UserRegisterDto) (*UserDto, error) {
	user, err := u.service.NewUser(input.Name, input.Email, input.Password)
	if err != nil {
		return nil, err
	}

	if err := u.repo.Create(user); err != nil {
		return nil, err
	}

	dto := MapUserResposeDto(user)
	return &dto, nil
}

func (u *UserUseCases) Get(userId uuid.UUID) (*UserDto, error) {
	user, err := u.repo.Get(userId)
	if err != nil {
		return nil, err
	}

	dto := MapUserResposeDto(user)
	return &dto, nil
}

func (u *UserUseCases) Update(userId uuid.UUID, input UserUpdateDto) (*UserDto, error) {
	user, err := u.repo.Get(userId)
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

	if err := u.repo.Update(user, updates); err != nil {
		return nil, err
	}

	dto := MapUserResposeDto(user)
	return &dto, nil
}

func (u *UserUseCases) UpdatePassword(userId uuid.UUID, input UserUpdatePasswordDto) error {
	user, err := u.repo.Get(userId)
	if err != nil {
		return err
	}

	if err := u.service.ValidationUpdatePassword(user, input.OldPassword, input.NewPassword, input.NewPasswordConfirm); err != nil {
		return err
	}

	if err := u.repo.UpdatePassword(user); err != nil {
		return err
	}

	return nil
}
