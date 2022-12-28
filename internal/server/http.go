package server

import (
	"github.com/gin-gonic/gin"
	"template-go/internal/config"
	"template-go/internal/server/routes"
	db "template-go/internal/sqlc/repositories"
)

func Server(projectConfiguration config.Config, store *db.SQLStore) error {
	router := SetupRouter(projectConfiguration, store)

	err := router.Run(projectConfiguration.HTTPServerAddress)
	if err != nil {
		return err
	}

	return nil
}

func SetupRouter(projectConfiguration config.Config, store *db.SQLStore) *gin.Engine {
	router := gin.Default()
	routes.InitUserRoutes(projectConfiguration, router, store)

	return router
}
