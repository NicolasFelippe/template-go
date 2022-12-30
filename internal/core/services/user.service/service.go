package user_service

import (
	"template-go/internal/core/domain"
	"template-go/internal/core/ports"
	"template-go/pkg/crypto"
	"template-go/pkg/uidgen"
)

type service struct {
	userRepository ports.UserRepository
	uidGen         uidgen.UIDGen
	crypto         crypto.Crypto
}

func New(
	userRepository ports.UserRepository,
	uidGen uidgen.UIDGen,
	crypto crypto.Crypto,
) *service {
	return &service{
		userRepository: userRepository,
		uidGen:         uidGen,
		crypto:         crypto,
	}
}

func (srv *service) CreateUser(
	username string,
	password string,
	fullName string,
	email string,
) (*domain.User, error) {

	hashedPassword, err := srv.crypto.HashPassword(password)
	if err != nil {
		return nil, err
	}

	userDomain := domain.NewUser(
		srv.uidGen.New(),
		username,
		hashedPassword,
		fullName,
		email,
	)

	result, err := srv.userRepository.CreateUser(userDomain)

	if err != nil {
		return nil, err
	}

	return result, nil
}
