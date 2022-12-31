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
	ListUsers(limit, offset *int) ([]*domain.User, error)
}

type AuthService interface {
	Authenticate(
		username,
		password string,
	) (*domain.User, error)
	ListUsers(limit, offset *int) ([]*domain.User, error)
}
