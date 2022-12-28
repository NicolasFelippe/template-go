package userrepository

import (
	"context"
	"fmt"
	"template-go/internal/core/domain"
	log "template-go/internal/logger"
	db "template-go/internal/sqlc/repositories"
)

type UserConfig struct {
	store db.SQLStore
}

func New(store *db.SQLStore) *UserConfig {
	return &UserConfig{
		store: *store,
	}
}

func (userConfig UserConfig) CreateUser(user domain.User) (domain.User, error) {

	createUserParams := db.CreateUserParams{
		ID:             user.ID,
		Username:       user.Username,
		HashedPassword: user.HashedPassword,
		FullName:       user.FullName,
		Email:          user.Email,
	}

	createUser, err := userConfig.store.CreateUser(context.Background(), createUserParams)
	if err != nil {
		log.Logger.Error(fmt.Sprintf("Got an error receiving messages: %v", err))
	}
	rsp := domain.User{
		Username: createUser.Username,
		FullName: createUser.FullName,
		Email:    createUser.Email,
	}
	return rsp, nil
}
