package userservice

import (
	"template-go/internal/core/domain"
	"template-go/internal/core/ports"
	"template-go/pkg/uidgen"
)

type service struct {
	userRepository ports.UserRepository
	uidGen         uidgen.UIDGen
}

func New(userRepository ports.UserRepository, uidGen uidgen.UIDGen) *service {
	return &service{
		userRepository: userRepository,
		uidGen:         uidGen,
	}
}

func (srv *service) CreateUser(
	username string,
	password string,
	fullName string,
	email string,
) (domain.User, error) {

	userDomain := domain.NewUser(
		srv.uidGen.New(),
		username,
		password,
		fullName,
		email,
	)

	result, err := srv.userRepository.CreateUser(userDomain)

	if err != nil {
		return domain.User{}, err
	}

	return result, nil
}
