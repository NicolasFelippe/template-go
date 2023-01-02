package userservice

import (
	"template-go/internal/core/domain"
	"template-go/internal/core/ports"
	"template-go/pkg/crypto"
	"template-go/pkg/uidgen"
)

type UserService struct {
	userRepository ports.UserRepository
	uidGen         uidgen.UIDGen
	crypto         crypto.Crypto
}

func New(
	userRepository ports.UserRepository,
	uidGen uidgen.UIDGen,
	crypto crypto.Crypto,
) *UserService {
	return &UserService{
		userRepository: userRepository,
		uidGen:         uidGen,
		crypto:         crypto,
	}
}

func (srv *UserService) CreateUser(
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
		srv.uidGen.New().String(),
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

func (srv *UserService) ListUsers(limit, offset *int) ([]*domain.User, error) {
	return srv.userRepository.Users(limit, offset)
}

func (srv *UserService) GetUserByUsername(username string) (*domain.User, error) {
	return srv.userRepository.GetUserByUsername(username)
}
