package userservice

import (
	"errors"
	"template-go/internal/core/domain"
	"template-go/internal/core/ports"
	"template-go/internal/handlers/usershandler"
	uidgen "template-go/pkg"
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

func (srv *service) CreateUser(userRequest usershandler.RequestUserDTO) (usershandler.ResponseCreateUserDTO, error) {
	userDomain := domain.NewUser(
		srv.uidGen.New(),
		userRequest.Username,
		userRequest.Password,
		userRequest.FullName,
		userRequest.Email,
	)
	result, err := srv.userRepository.CreateUser(userDomain)

	if err != nil {
		return usershandler.ResponseCreateUserDTO{}, errors.New("error on create user")
	}

	rsp := usershandler.ResponseCreateUserDTO{
		FullName:          result.FullName,
		Email:             result.Email,
		Username:          result.Username,
		CreatedAt:         result.CreatedAt,
		PasswordChangedAt: result.PasswordChangedAt,
	}

	return rsp, nil
}
