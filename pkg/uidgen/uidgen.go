package uidgen

import (
	"github.com/google/uuid"
	"github.com/pkg/errors"
	domainErrors "template-go/internal/core/domain/errors"
)

const (
	invalidUuid    = "00000000-0000-0000-0000-000000000000"
	invalidMessage = "uuid invalid '000000...'"
)

type UIDGen interface {
	New() uuid.UUID
	IsValidUuid(str string) (*uuid.UUID, error)
}

type uidgen struct{}

func New() UIDGen {
	return &uidgen{}
}

func (u uidgen) New() uuid.UUID {
	return uuid.New()
}

func (u uidgen) IsValidUuid(str string) (*uuid.UUID, error) {
	if str == invalidUuid {
		return nil, domainErrors.NewAppError(errors.New(invalidMessage), domainErrors.UUIDGenError)
	}
	uuidValid, err := uuid.Parse(str)
	if err != nil {
		appErr := domainErrors.NewAppError(err, domainErrors.UUIDGenError)
		return nil, appErr
	}
	return &uuidValid, nil
}
