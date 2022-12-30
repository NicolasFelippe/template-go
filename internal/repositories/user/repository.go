package user

import (
	"context"
	"template-go/internal/core/domain"
	db "template-go/internal/sqlc/repositories"
)

type UserRepository struct {
	store db.Store
}

func New(store db.Store) *UserRepository {
	return &UserRepository{
		store: store,
	}
}

func (userConfig UserRepository) CreateUser(user *domain.User) (*domain.User, error) {

	createUserParams := db.CreateUserParams{
		ID:             user.ID,
		Username:       user.Username,
		HashedPassword: user.HashedPassword,
		FullName:       user.FullName,
		Email:          user.Email,
	}

	_, err := userConfig.store.CreateUser(context.Background(), createUserParams)
	if err != nil {
		return nil, err
	}

	return user, nil
}
