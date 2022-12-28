package domain

import (
	"github.com/google/uuid"
	"time"
)

type User struct {
	ID                uuid.UUID `json:"id"`
	Username          string    `json:"username"`
	HashedPassword    string    `json:"hashed_password"`
	FullName          string    `json:"full_name"`
	Email             string    `json:"email"`
	PasswordChangedAt time.Time `json:"password_changed_at"`
	CreatedAt         time.Time `json:"created_at"`
}

func NewUser(id uuid.UUID, username string, hashedPassword string, fullName string, email string) User {
	return User{
		ID:                id,
		Username:          username,
		HashedPassword:    hashedPassword,
		FullName:          fullName,
		Email:             email,
		CreatedAt:         time.Now(),
		PasswordChangedAt: time.Date(0001, 01, 01, 00, 00, 00, 0, nil),
	}
}
