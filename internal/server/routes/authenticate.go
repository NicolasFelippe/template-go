package routes

import (
	"github.com/gin-gonic/gin"
	"template-go/internal/config"
	"template-go/internal/core/services/authenticationservice"
	"template-go/internal/core/services/sessionservice"
	"template-go/internal/core/services/userservice"
	"template-go/internal/handlers/authhandler"
	"template-go/internal/repositories/sessionrepo"
	"template-go/internal/repositories/userrepo"
	db "template-go/internal/sqlc/repositories"
	"template-go/pkg/crypto"
	"template-go/pkg/makertoken"
	"template-go/pkg/uidgen"
)

func InitAuthRoutes(route *gin.Engine, store db.Store, tokenMaker makertoken.Maker, config config.Config) {
	newUuidGen := uidgen.New()
	newCrypto := crypto.New()

	userRepository := userrepo.New(store, newUuidGen)
	userService := userservice.New(userRepository, newUuidGen, newCrypto)

	sessionRepository := sessionrepo.New(store, newUuidGen)

	authService := authenticationservice.New(tokenMaker, newUuidGen, newCrypto, config, userService)

	service := sessionservice.New(sessionRepository, config, authService)
	handler := authhandler.NewAuthHTTPHandler(service)

	groupRoute := route.Group("/api/v1")
	groupRoute.POST("/authenticate", handler.Login)
}
