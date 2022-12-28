package ports

import (
	"template-go/internal/core/domain"
)

type UserService interface {
	CreateUser(
		username string,
		password string,
		fullName string,
		email string,
	) (domain.User, error)
}
