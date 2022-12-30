package uidgen

import "github.com/google/uuid"

const (
	invalidUuid = "00000000-0000-0000-0000-000000000000"
)

type UIDGen interface {
	New() uuid.UUID
	IsValidUuid(str string) bool
}

type uidgen struct{}

func New() UIDGen {
	return &uidgen{}
}

func (u uidgen) New() uuid.UUID {
	return uuid.New()
}

func (u uidgen) IsValidUuid(str string) bool {
	if str == invalidUuid {
		return false
	}
	_, err := uuid.Parse(str)
	return err == nil
}
