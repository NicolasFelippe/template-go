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
	GetUserByUsername(username string) (*domain.User, error)
}

type SessionService interface {
	CreateSession(
		username,
		password,
		userAgent,
		clientIp string,
	) (*domain.Session, *domain.Authentication, error)
	//GetSession(id string) (*domain.Session, error)
}

type AuthenticationService interface {
	Authenticate(
		username,
		password string,
	) (*domain.Authentication, error)
}
