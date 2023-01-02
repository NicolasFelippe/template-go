package authhandler

import (
	"template-go/internal/handlers/usershandler"
	"time"
)

type LoginResponse struct {
	ID                    string                    `json:"session_id"`
	AccessToken           string                    `json:"access_token"`
	AccessTokenExpiresAt  time.Time                 `json:"access_token_expires_at"`
	RefreshToken          string                    `json:"refresh_token"`
	RefreshTokenExpiresAt time.Time                 `json:"refresh_token_expires_at"`
	User                  usershandler.UserResponse `json:"user"`
}
