package authentication

import (
	"template-go/internal/core/domain/users"
	"time"
)

type Authentication struct {
	SessionID             string
	User                  users.User
	AccessToken           string
	RefreshToken          string
	AccessTokenExpiresAt  time.Time
	RefreshTokenExpiresAt time.Time
}

func NewAuthentication(
	user users.User,
	accessToken,
	refreshToken,
	sessionId string,
	accessTokenExpiresAt,
	refreshTokenExpiresAt time.Time,
) *Authentication {
	return &Authentication{
		SessionID:             sessionId,
		User:                  user,
		AccessToken:           accessToken,
		RefreshToken:          refreshToken,
		AccessTokenExpiresAt:  accessTokenExpiresAt,
		RefreshTokenExpiresAt: refreshTokenExpiresAt,
	}
}
