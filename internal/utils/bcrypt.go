package utils

import (
	b "golang.org/x/crypto/bcrypt"
)

type Crypt interface {
	Hash(password string) (string, error)
	Check(password, hash string) error
}

type crypt struct{}

func NewBcrypt() Crypt {
	return &crypt{}
}

// Hash implementa a função da interface para gerar um hash bcrypt.
func (s *crypt) Hash(password string) (string, error) {
	hashedPassword, err := b.GenerateFromPassword([]byte(password), b.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

// Check implementa a função da interface para comparar a senha com o hash.
func (s *crypt) Check(password, hash string) error {
	return b.CompareHashAndPassword([]byte(hash), []byte(password))
}
