package user

import (
	"controlF_back/internal/models"
	"fmt"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	UserRepository UserRepositoryInterface
}

func NewUserService(repo UserRepositoryInterface) *UserService {
	return &UserService{
		UserRepository: repo,
	}
}

func (s *UserService) Create(input UserRegister) (*UserDto, error) {
	// Nao deve usar o bcrypt diretaemnte
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("erro ao gerar hash da nova senha: %w", err)
	}

	user := &models.User{
		Name:     input.Name,
		Email:    input.Email,
		Password: string(hashedPassword),
		Type:     models.UserTypePersonal,
	}

	if err := s.UserRepository.Create(user); err != nil {
		return nil, err
	}

	dto := NewUserDto(user)
	return &dto, nil
}

func (s *UserService) Get(userId uuid.UUID) (*UserDto, error) {
	user, err := models.GetUser(userId)
	if err != nil {
		return nil, err
	}

	dto := NewUserDto(user)
	return &dto, nil
}

func (s *UserService) Update(userId uuid.UUID, input UserUpdate) (*UserDto, error) {
	user, err := models.GetUser(userId)
	if err != nil {
		return nil, err
	}

	updates := make(map[string]interface{})
	if input.Name != nil {
		updates["name"] = input.Name
	}
	if input.Email != nil {
		updates["email"] = input.Email
	}
	if input.Type != nil {
		updates["type"] = input.Type
	}
	if input.CompanyID != nil {
		updates["company_id"] = input.CompanyID
	}

	if err := user.Update(updates); err != nil {
		return nil, err
	}

	dto := NewUserDto(user)
	return &dto, nil
}

func (s *UserService) UpdatePassword(userId uuid.UUID, input UserUpdatePassword) (*UserDto, error) {
	user, err := models.GetUser(userId)
	if err != nil {
		return nil, err
	}

	// Verifica se as novas senhas conferem
	if input.NewPassword != input.NewPasswordConfirm {
		return nil, fmt.Errorf("senhas não correspondem")
	}

	// Verifica se a senha antiga confere com a que está no banco
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.OldPassword)); err != nil {
		return nil, fmt.Errorf("senha antiga incorreta")
	}

	// Gera o hash da nova senha
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.NewPassword), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("erro ao gerar hash da nova senha: %w", err)
	}

	// Atualiza a senha do usuário
	user.Password = string(hashedPassword)
	if err := user.Save(); err != nil { // assumindo que Save persiste no banco
		return nil, fmt.Errorf("erro ao atualizar senha: %w", err)
	}

	dto := NewUserDto(user)
	return &dto, nil
}
