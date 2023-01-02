package domain

import "time"

type Session struct {
	ID           string    `json:"id"`
	UserID       string    `json:"user_id"`
	RefreshToken string    `json:"refresh_token"`
	UserAgent    string    `json:"user_agent"`
	ClientIp     string    `json:"client_ip"`
	IsBlocked    bool      `json:"is_blocked"`
	ExpiresAt    time.Time `json:"expires_at"`
	User         User      `json:"user"`
}

func NewSession(
	id,
	userAgent,
	refreshToken,
	userID,
	clientIp string,
	expiresAt time.Time,
	user User) *Session {
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
