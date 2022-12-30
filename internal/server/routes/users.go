package routes

import (
	"github.com/gin-gonic/gin"
	"template-go/internal/core/services/user.service"
	"template-go/internal/handlers/users.handler"
	"template-go/internal/repositories/user"
	db "template-go/internal/sqlc/repositories"
	"template-go/pkg/crypto"
	"template-go/pkg/uidgen"
)

func InitUserRoutes(route *gin.Engine, store db.Store) {
	repository := user.New(store)
	service := user_service.New(repository, uidgen.New(), crypto.New())
	handler := users_handler.NewUserHTTPHandler(service)

	groupRoute := route.Group("/api/v1")
	groupRoute.POST("/users", handler.CreateUser)
}
