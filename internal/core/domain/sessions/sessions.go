package sessions

import (
	"template-go/internal/core/domain/users"
	"time"
)

type Session struct {
	ID           string
	UserID       string
	RefreshToken string
	UserAgent    string
	ClientIp     string
	IsBlocked    bool
	ExpiresAt    time.Time
	User         users.User
}

func NewSession(
	id,
	userAgent,
	refreshToken,
	userID,
	clientIp string,
	expiresAt time.Time,
	user users.User) *Session {
	return &Session{
		ID:           id,
		UserAgent:    userAgent,
		ClientIp:     clientIp,
		RefreshToken: refreshToken,
		IsBlocked:    false,
		UserID:       userID,
		ExpiresAt:    expiresAt,
		User:         user,
	}
}
