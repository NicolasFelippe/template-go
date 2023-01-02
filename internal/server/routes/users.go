package routes

import (
	"github.com/gin-gonic/gin"
	"template-go/internal/core/services/userservice"
	"template-go/internal/handlers/usershandler"
	"template-go/internal/repositories/userrepo"
	db "template-go/internal/sqlc/repositories"
	"template-go/pkg/crypto"
	"template-go/pkg/uidgen"
)

func InitUserRoutes(route *gin.Engine, store db.Store) {
	repository := userrepo.New(store, uidgen.New())
	service := userservice.New(repository, uidgen.New(), crypto.New())
	handler := usershandler.NewUserHTTPHandler(service)

	groupRoute := route.Group("/api/v1")
	groupRoute.POST("/users", handler.CreateUser)
	groupRoute.GET("/users", handler.ListUsers)
}
