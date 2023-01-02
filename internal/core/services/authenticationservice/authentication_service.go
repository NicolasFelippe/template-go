package authenticationservice

import (
	"template-go/internal/config"
	"template-go/internal/core/domain"
	"template-go/internal/core/ports"
	"template-go/pkg/crypto"
	"template-go/pkg/makertoken"
	"template-go/pkg/uidgen"
)

type AuthenticationService struct {
	makerToken  makertoken.Maker
	uidGen      uidgen.UIDGen
	crypto      crypto.Crypto
	userService ports.UserService
	config      config.Config
}

func New(
	makerToken makertoken.Maker,
	uidGen uidgen.UIDGen,
	crypto crypto.Crypto,
	config config.Config,
	userService ports.UserService,
) *AuthenticationService {
	return &AuthenticationService{
		makerToken,
		uidGen,
		crypto,
		userService,
		config,
	}
}

func (service *AuthenticationService) Authenticate(
	username,
	password string,
) (*domain.Authentication, error) {

	user, err := service.userService.GetUserByUsername(username)
	if err != nil {
		return nil, err
	}

	err = service.crypto.CheckPassword(password, user.HashedPassword)
	if err != nil {
		return nil, err
	}

	accessToken, accessPayload, err := service.makerToken.CreateToken(user.Username, service.config.AccessTokenDuration)

	refreshToken, refreshPayload, err := service.makerToken.CreateToken(
		user.Username,
		service.config.RefreshTokenDuration,
	)
	newAuth := domain.NewAuthentication(
		*user,
		accessToken,
		refreshToken,
		service.uidGen.New().String(),
		accessPayload.ExpiredAt,
		refreshPayload.ExpiredAt,
	)
	return newAuth, nil
}
