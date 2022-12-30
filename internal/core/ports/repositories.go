package ports

import "template-go/internal/core/domain"

type UserRepository interface {
	CreateUser(user *domain.User) (*domain.User, error)
}
