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

func (userConfig UserRepository) Users(limit, offset *int) ([]*domain.User, error) {
	var listUsersParams = db.ListUsersParams{
		Limit:  int32(*limit),
		Offset: int32(*offset),
	}
	result, err := userConfig.store.ListUsers(context.Background(), listUsersParams)
	if err != nil {
		return nil, err
	}

	var users []*domain.User
	for _, user := range result {
		users = append(users, &domain.User{
			ID:                user.ID,
			FullName:          user.FullName,
			Username:          user.Username,
			Email:             user.Email,
			PasswordChangedAt: user.PasswordChangedAt,
			CreatedAt:         user.CreatedAt,
		})
	}
	return users, nil
}
