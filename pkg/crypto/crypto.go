package crypto

import (
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
	domainErrors "template-go/internal/core/domain/errors"
)

const (
	hashError = "Error in create hash password"
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
		return "", domainErrors.NewAppError(errors.Wrap(err, hashError), domainErrors.HashError)
	}
	return string(hashedPassword), nil
}

func (u crypto) CheckPassword(password string, hashedPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		return domainErrors.NewAppError(err, domainErrors.HashError)
	}
	return nil
}
