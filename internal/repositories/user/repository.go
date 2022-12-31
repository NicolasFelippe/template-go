package user

import (
	"context"
	"errors"
	"template-go/internal/core/domain"
	db "template-go/internal/sqlc/repositories"
	"template-go/pkg/uidgen"
)

type UserRepository struct {
	store  db.Store
	uidGen uidgen.UIDGen
}

func New(store db.Store, uidGen uidgen.UIDGen) *UserRepository {
	return &UserRepository{
		store:  store,
		uidGen: uidGen,
	}
}

func (userRepository UserRepository) CreateUser(user *domain.User) (*domain.User, error) {

	uuid, isValid := userRepository.uidGen.IsValidUuid(user.ID)
	if !isValid {
		return nil, errors.New("UUID invalid")
	}

	createUserParams := db.CreateUserParams{
		ID:             *uuid,
		Username:       user.Username,
		HashedPassword: user.HashedPassword,
		FullName:       user.FullName,
		Email:          user.Email,
	}

	_, err := userRepository.store.CreateUser(context.Background(), createUserParams)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (userRepository UserRepository) Users(limit, offset *int) ([]*domain.User, error) {
	var listUsersParams = db.ListUsersParams{
		Limit:  int32(*limit),
		Offset: int32(*offset),
	}
	result, err := userRepository.store.ListUsers(context.Background(), listUsersParams)
	if err != nil {
		return nil, err
	}

	var users []*domain.User
	for _, user := range result {
		users = append(users, &domain.User{
			ID:                user.ID.String(),
			FullName:          user.FullName,
			Username:          user.Username,
			Email:             user.Email,
			PasswordChangedAt: user.PasswordChangedAt,
			CreatedAt:         user.CreatedAt,
		})
	}
	return users, nil
}
