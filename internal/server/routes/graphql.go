package routes

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
	"template-go/internal/core/services/userservice"
	"template-go/internal/graph"
	"template-go/internal/repositories/userrepo"
	db "template-go/internal/sqlc/repositories"
	"template-go/pkg/crypto"
	"template-go/pkg/uidgen"
)

func InitGraphQlRoutes(route *gin.Engine, store db.Store) {
	groupRoute := route.Group("/graphql/v1")
	groupRoute.POST("/query", graphqlHandler(&store))
	groupRoute.GET("/", playgroundHandler())
}

// Defining the Graphql handler
func graphqlHandler(store *db.Store) gin.HandlerFunc {
	// NewExecutableSchema and Config are in the generated.go file
	// Resolver is in the resolver.go file
	repository := userrepo.New(*store, uidgen.New())
	service := userservice.New(repository, uidgen.New(), crypto.New())

	h := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{UserService: service}}))
	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

// Defining the Playground handler
func playgroundHandler() gin.HandlerFunc {
	h := playground.Handler("GraphQL", "/graphql/v1/query")
	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}
