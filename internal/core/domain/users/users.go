package users

import (
	"time"
)

type User struct {
	ID                string
	Username          string
	HashedPassword    string
	FullName          string
	Email             string
	PasswordChangedAt time.Time
	CreatedAt         time.Time
}
