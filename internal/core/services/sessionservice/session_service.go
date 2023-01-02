package sessionservice

import (
	"template-go/internal/config"
	"template-go/internal/core/domain/authentication"
	"template-go/internal/core/domain/sessions"
)

type SessionService struct {
	repository            sessions.SessionRepository
	config                config.Config
	authenticationService authentication.AuthenticationService
}

func New(
	repository sessions.SessionRepository,
	config config.Config,
	authenticationService authentication.AuthenticationService,
) *SessionService {
	return &SessionService{
		repository,
		config,
		authenticationService,
	}
}

func (service *SessionService) CreateSession(
	username,
	password,
	userAgent,
	clientIp string,
) (*sessions.Session, *authentication.Authentication, error) {

	auth, err := service.authenticationService.Authenticate(username, password)
	if err != nil {
		return nil, nil, err
	}

	session := sessions.NewSession(
		auth.SessionID,
		userAgent,
		auth.RefreshToken,
		auth.User.ID,
		clientIp,
		auth.RefreshTokenExpiresAt,
		auth.User,
	)

	createSession, err := service.repository.CreateSession(session)
	if err != nil {
		return nil, nil, err
	}

	return createSession, auth, nil
}

//func (service *SessionService) GetSession(id string) (*domain.Session, error) {
//	return service.repository.GetSession(id)
//}
