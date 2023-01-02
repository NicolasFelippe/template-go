package sessions

import "template-go/internal/core/domain/authentication"

type SessionService interface {
	CreateSession(
		username,
		password,
		userAgent,
		clientIp string,
	) (*Session, *authentication.Authentication, error)
	//GetSession(id string) (*domain.Session, error)
}
