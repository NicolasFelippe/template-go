package usershandler

import "template-go/internal/core/domain/users"

func toResponseModel(user *users.User) *UserResponse {
	return &UserResponse{
		ID:                user.ID,
		FullName:          user.FullName,
		Email:             user.Email,
		Username:          user.Username,
		CreatedAt:         user.CreatedAt,
		PasswordChangedAt: user.PasswordChangedAt,
	}
}
