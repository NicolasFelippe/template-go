package authenticationservice

import (
	"template-go/internal/config"
	"template-go/internal/core/domain/authentication"
	"template-go/internal/core/domain/users"
	"template-go/pkg/crypto"
	"template-go/pkg/makertoken"
	"template-go/pkg/uidgen"
)

type AuthenticationService struct {
	makerToken  makertoken.Maker
	uidGen      uidgen.UIDGen
	crypto      crypto.Crypto
	userService users.UserService
	config      config.Config
}

func New(
	makerToken makertoken.Maker,
	uidGen uidgen.UIDGen,
	crypto crypto.Crypto,
	config config.Config,
	userService users.UserService,
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
) (*authentication.Authentication, error) {

	user, err := service.userService.GetUserByUsername(username)
	if err != nil {
		return nil, err
	}

	err = service.crypto.CheckPassword(password, user.HashedPassword)
	if err != nil {
		return nil, err
	}

	accessToken, accessPayload, err := service.makerToken.CreateToken(user.Username, service.config.AccessTokenDuration)
	if err != nil {
		return nil, err
	}

	refreshToken, refreshPayload, err := service.makerToken.CreateToken(
		user.Username,
		service.config.RefreshTokenDuration,
	)
	if err != nil {
		return nil, err
	}

	newAuth := authentication.NewAuthentication(
		*user,
		accessToken,
		refreshToken,
		service.uidGen.New().String(),
		accessPayload.ExpiredAt,
		refreshPayload.ExpiredAt,
	)
	return newAuth, nil
}
