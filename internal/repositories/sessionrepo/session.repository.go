package sessionrepo

import (
	"context"
	"errors"
	"template-go/internal/core/domain"
	db "template-go/internal/sqlc/repositories"
	"template-go/pkg/uidgen"
)

type SessionRepository struct {
	store  db.Store
	uidGen uidgen.UIDGen
}

func New(store db.Store, uidGen uidgen.UIDGen) *SessionRepository {
	return &SessionRepository{
		store:  store,
		uidGen: uidGen,
	}
}

func (repository *SessionRepository) CreateSession(session *domain.Session) (*domain.Session, error) {

	uuid, isValid := repository.uidGen.IsValidUuid(session.ID)
	if !isValid {
		return nil, errors.New("UUID invalid")
	}

	uid, isValid := repository.uidGen.IsValidUuid(session.UserID)
	if !isValid {
		return nil, errors.New("UUID invalid")
	}

	arg := db.CreateSessionParams{
		ID:           *uuid,
		ClientIp:     session.ClientIp,
		UserID:       *uid,
		UserAgent:    session.UserAgent,
		RefreshToken: session.RefreshToken,
		ExpiresAt:    session.ExpiresAt,
		IsBlocked:    session.IsBlocked,
	}
	createSession, err := repository.store.CreateSession(context.Background(), arg)
	if err != nil {
		return nil, err
	}
	rsp := domain.NewSession(
		createSession.ID.String(),
		createSession.UserAgent,
		createSession.RefreshToken,
		createSession.UserID.String(),
		createSession.ClientIp,
		createSession.ExpiresAt,
		session.User,
	)

	return rsp, nil
}
