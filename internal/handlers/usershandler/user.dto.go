package usershandler

import (
	"github.com/google/uuid"
	"template-go/internal/core/domain"
	"time"
)

type ResponseCreateUserDTO struct {
	ID                uuid.UUID `json:"id"`
	Username          string    `json:"username"`
	FullName          string    `json:"full_name"`
	Email             string    `json:"email"`
	PasswordChangedAt time.Time `json:"password_changed_at"`
	CreatedAt         time.Time `json:"created_at"`
}

type RequestUserDTO struct {
	Username string `json:"username" binding:"required,alphanum"`
	Password string `json:"crypto" binding:"required,min=6"`
	FullName string `json:"full_name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
}

func (td *ResponseCreateUserDTO) FromDomain(user *domain.User) {
	if td == nil {
		td = &ResponseCreateUserDTO{}
	}

	td.ID = user.ID
	td.Username = user.Username
	td.Email = user.Email
	td.FullName = user.FullName
	td.CreatedAt = user.CreatedAt
	td.PasswordChangedAt = user.PasswordChangedAt
}
