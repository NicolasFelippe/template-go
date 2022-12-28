package uidgen

import "github.com/google/uuid"

type UIDGen interface {
	New() uuid.UUID
}

type uidgen struct{}

func New() UIDGen {
	return &uidgen{}
}

func (u uidgen) New() uuid.UUID {
	return uuid.New()
}
