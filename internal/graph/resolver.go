package graph

import (
	"template-go/internal/core/domain/users"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	UserService users.UserService
}
