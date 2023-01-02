package server

import (
	"github.com/gin-gonic/gin"
	"template-go/internal/config"
	errorsHandlers "template-go/internal/handlers/errors"
	"template-go/internal/server/middlewares"
	"template-go/internal/server/routes"
	db "template-go/internal/sqlc/repositories"
	"template-go/pkg/makertoken"
)

type Server struct {
	config     config.Config
	store      db.Store
	tokenMaker makertoken.Maker
	router     *gin.Engine
}

func NewServer(config config.Config, store db.Store, tokenMaker makertoken.Maker) (*Server, error) {
	router := SetupRouter(store, tokenMaker, config)

	server := &Server{
		config:     config,
		store:      store,
		router:     router,
		tokenMaker: tokenMaker,
	}
	return server, nil
}

func (srv *Server) Start() error {
	return srv.router.Run(srv.config.HTTPServerAddress)
}

func SetupRouter(store db.Store, tokenMaker makertoken.Maker, config config.Config) *gin.Engine {
	routerDefault := gin.Default()

	routerDefault.Use(errorsHandlers.Handler)
	authRoutes := routerDefault.Group("/api/v1").Use(middlewares.AuthMiddleware(tokenMaker))

	routes.InitUserRoutes(routerDefault, authRoutes, store)
	//routes.InitGraphQlRoutes(router, authRoutes, store)
	//routes.InitAuthRoutes(router, authRoutes, store, tokenMaker, config)
	return routerDefault
}
