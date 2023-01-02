package ports

import "template-go/internal/core/domain"

type UserRepository interface {
	CreateUser(user *domain.User) (*domain.User, error)
	Users(limit, offset *int) ([]*domain.User, error)
	GetUserByUsername(username string) (*domain.User, error)
}

type SessionRepository interface {
	CreateSession(session *domain.Session) (*domain.Session, error)
	//GetSession(id string) (*domain.Session, error)
}
