package sessionservice

import (
	"template-go/internal/config"
	"template-go/internal/core/domain"
	"template-go/internal/core/ports"
)

type SessionService struct {
	repository            ports.SessionRepository
	config                config.Config
	authenticationService ports.AuthenticationService
}

func New(
	repository ports.SessionRepository,
	config config.Config,
	authenticationService ports.AuthenticationService,
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
) (*domain.Session, *domain.Authentication, error) {

	auth, err := service.authenticationService.Authenticate(username, password)
	if err != nil {
		return nil, nil, err
	}

	session := domain.NewSession(
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
