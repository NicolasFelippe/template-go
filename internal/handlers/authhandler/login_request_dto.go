package authhandler

import (
	"template-go/internal/handlers/usershandler"
	"time"
)

type LoginUserRequest struct {
	Username string `json:"username" binding:"required,alphanum"`
	Password string `json:"password" binding:"required,min=6"`
}

type LoginUserResponse struct {
	ID                    string                             `json:"session_id"`
	AccessToken           string                             `json:"access_token"`
	AccessTokenExpiresAt  time.Time                          `json:"access_token_expires_at"`
	RefreshToken          string                             `json:"refresh_token"`
	RefreshTokenExpiresAt time.Time                          `json:"refresh_token_expires_at"`
	User                  usershandler.ResponseCreateUserDTO `json:"user"`
}
