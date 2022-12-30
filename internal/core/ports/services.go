package ports

import (
	"template-go/internal/core/domain"
)

type UserService interface {
	CreateUser(
		username,
		password,
		fullName,
		email string,
	) (*domain.User, error)
}
