package sessionrepo

import (
	"context"
	"template-go/internal/core/domain/sessions"
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

func (repository *SessionRepository) CreateSession(session *sessions.Session) (*sessions.Session, error) {

	uuid, err := repository.uidGen.IsValidUuid(session.ID)
	if err != nil {
		return nil, err
	}

	uid, err := repository.uidGen.IsValidUuid(session.UserID)
	if err != nil {
		return nil, err
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
	rsp := sessions.NewSession(
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
