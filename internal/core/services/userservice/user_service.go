package userservice

import (
	"template-go/internal/core/domain/users"
	"template-go/pkg/crypto"
	"template-go/pkg/uidgen"
)

type UserService struct {
	userRepository users.UserRepository
	uidGen         uidgen.UIDGen
	crypto         crypto.Crypto
}

func New(
	userRepository users.UserRepository,
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
	user *users.User,
) (*users.User, error) {

	hashedPassword, err := srv.crypto.HashPassword(user.HashedPassword)
	if err != nil {
		return nil, err
	}
	user.HashedPassword = hashedPassword
	user.ID = srv.uidGen.New().String()

	result, err := srv.userRepository.CreateUser(user)

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (srv *UserService) ListUsersByPagination(limit, offset *int) ([]users.User, error) {
	return srv.userRepository.ListUsersByPagination(limit, offset)
}

func (srv *UserService) GetUserByUsername(username string) (*users.User, error) {
	return srv.userRepository.GetUserByUsername(username)
}
