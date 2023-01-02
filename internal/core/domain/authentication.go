package domain

import "time"

type Authentication struct {
	SessionID             string    `json:"session_id"`
	User                  User      `json:"user"`
	AccessToken           string    `json:"access_token"`
	RefreshToken          string    `json:"refresh_token"`
	AccessTokenExpiresAt  time.Time `json:"access_token_expires_at"`
	RefreshTokenExpiresAt time.Time `json:"refresh_token_expires_at"`
}

func NewAuthentication(
	user User,
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
