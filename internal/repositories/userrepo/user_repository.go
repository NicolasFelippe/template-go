package userrepo

import (
	"context"
	"github.com/lib/pq"
	domainErrors "template-go/internal/core/domain/errors"
	domain "template-go/internal/core/domain/users"
	"template-go/internal/logger"
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

	uuid, err := userRepository.uidGen.IsValidUuid(user.ID)
	if err != nil {
		return nil, err
	}

	createUserParams := toDBModel(user, uuid)

	result, err := userRepository.store.CreateUser(context.Background(), *createUserParams)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			switch pqErr.Code.Name() {
			case "unique_violation":
				appErr := domainErrors.NewAppError(pqErr, domainErrors.RepositoryError)
				return nil, appErr
			}
		}
		appErr := domainErrors.NewAppError(err, domainErrors.RepositoryError)
		return nil, appErr
	}
	return toDomainModel(&result), nil
}

func (userRepository UserRepository) ListUsersByPagination(limit, offset *int) ([]domain.User, error) {
	var listUsersParams = db.ListUsersParams{
		Limit:  int32(*limit),
		Offset: int32(*offset),
	}
	result, err := userRepository.store.ListUsers(context.Background(), listUsersParams)
	if err != nil {
		appErr := domainErrors.NewAppError(err, domainErrors.RepositoryError)
		return nil, appErr
	}

	return listToDomainModel(result), nil
}

func (userRepository UserRepository) GetUserByUsername(username string) (*domain.User, error) {
	result, err := userRepository.store.GetUser(context.Background(), username)

	if pqErr, ok := err.(*pq.Error); ok {
		test := pqErr.Code.Name()
		logger.Logger.Info(test)
		//switch pqErr.Code.Name() {
		//case "unique_violation":
		//	appErr := domainErrors.NewAppError(pqErr, domainErrors.RepositoryError)
		//	return nil, appErr
		//}
		appErr := domainErrors.NewAppErrorWithType(domainErrors.NotFound)
		return nil, appErr
	}

	return toDomainModel(&result), nil
}
