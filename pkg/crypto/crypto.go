package crypto

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

type Crypto interface {
	HashPassword(password string) (string, error)
	CheckPassword(password string, hashedPassword string) error
}

type crypto struct{}

func New() Crypto {
	return &crypto{}
}

func (u crypto) HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("failed to has crypto: %w", err)
	}
	return string(hashedPassword), nil
}

func (u crypto) CheckPassword(password string, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
