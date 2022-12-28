package ports

import (
	"template-go/internal/handlers/usershandler"
)

type UserService interface {
	CreateUser(user usershandler.RequestUserDTO) (usershandler.ResponseCreateUserDTO, error)
}
