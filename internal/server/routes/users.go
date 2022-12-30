package routes

import (
	"github.com/gin-gonic/gin"
	"template-go/internal/config"
	"template-go/internal/core/services/userservice"
	"template-go/internal/handlers/usershandler"
	"template-go/internal/repositories/userrepository"
	db "template-go/internal/sqlc/repositories"
	"template-go/pkg/password"
	"template-go/pkg/uidgen"
)

func InitUserRoutes(projectConfiguration config.Config, route *gin.Engine, store *db.SQLStore) {
	repository := userrepository.New(store)
	service := userservice.New(repository, uidgen.New(), password.New())
	handler := usershandler.NewUserHTTPHandler(service)

	groupRoute := route.Group("/api/v1")
	groupRoute.POST("/users", handler.CreateUser)
}
