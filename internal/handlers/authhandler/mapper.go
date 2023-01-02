package authhandler

import (
	"template-go/internal/core/domain/authentication"
	"template-go/internal/core/domain/sessions"
	"template-go/internal/handlers/usershandler"
)

func toResponseModel(
	auth *authentication.Authentication,
	session *sessions.Session,
) *LoginResponse {
	return &LoginResponse{
		User: usershandler.UserResponse{
			ID:                auth.User.ID,
			Email:             auth.User.Email,
			PasswordChangedAt: auth.User.PasswordChangedAt,
			FullName:          auth.User.FullName,
			Username:          auth.User.Username,
			CreatedAt:         auth.User.CreatedAt,
		},
		ID:                    session.ID,
		AccessToken:           auth.AccessToken,
		RefreshToken:          session.RefreshToken,
		AccessTokenExpiresAt:  auth.AccessTokenExpiresAt,
		RefreshTokenExpiresAt: auth.RefreshTokenExpiresAt,
	}

}
