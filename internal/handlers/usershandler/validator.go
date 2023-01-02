package usershandler

import (
	"github.com/gin-gonic/gin"
	"template-go/internal/core/domain/users"
)

type UserValidator struct {
	Username string `json:"username" binding:"required,alphanum"`
	Password string `json:"password" binding:"required,min=6"`
	FullName string `json:"full_name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
}

type PageValidator struct {
	PageID   int `form:"page_id" binding:"required,min=1"`
	PageSize int `form:"page_size" binding:"required,min=5,max=10"`
}

func BindJson(c *gin.Context) (*users.User, error) {
	var json UserValidator
	if err := c.ShouldBindJSON(&json); err != nil {
		return nil, err
	}

	user := &users.User{
		Username:       json.Username,
		Email:          json.Email,
		FullName:       json.FullName,
		HashedPassword: json.Password,
	}

	return user, nil
}
