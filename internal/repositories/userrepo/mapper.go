package userrepo

import (
	"github.com/google/uuid"
	domain "template-go/internal/core/domain/users"
	db "template-go/internal/sqlc/repositories"
)

func toDBModel(entity *domain.User, uuid *uuid.UUID) *db.CreateUserParams {
	return &db.CreateUserParams{
		ID:             *uuid,
		Username:       entity.Username,
		HashedPassword: entity.HashedPassword,
		FullName:       entity.FullName,
		Email:          entity.Email,
	}
}

func toDomainModel(entity *db.User) *domain.User {
	return &domain.User{
		ID:                entity.ID.String(),
		Username:          entity.Username,
		FullName:          entity.FullName,
		Email:             entity.Email,
		HashedPassword:    entity.HashedPassword,
		CreatedAt:         entity.CreatedAt,
		PasswordChangedAt: entity.PasswordChangedAt,
	}
}

func listToDomainModel(entities []db.User) []domain.User {
	var users = make([]domain.User, len(entities))

	for i, element := range entities {
		users[i] = *toDomainModel(&element)
	}
	return users
}
