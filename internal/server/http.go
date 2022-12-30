package server

import (
	"github.com/gin-gonic/gin"
	"template-go/internal/config"
	"template-go/internal/server/routes"
	db "template-go/internal/sqlc/repositories"
)

type Server struct {
	config config.Config
	store  db.Store
	//tokenMaker token.Maker
	router *gin.Engine
}

func NewServer(projectConfiguration config.Config, store db.Store) (*Server, error) {
	router := SetupRouter(store)

	server := &Server{
		config: projectConfiguration,
		store:  store,
		router: router,
	}
	return server, nil
}

func (srv *Server) Start() error {
	return srv.router.Run(srv.config.HTTPServerAddress)
}

func SetupRouter(store db.Store) *gin.Engine {
	router := gin.Default()
	routes.InitUserRoutes(router, store)

	return router
}
