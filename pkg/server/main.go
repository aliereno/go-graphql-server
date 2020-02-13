package server

import (
	"github.com/aliereno/go-graphql-server/internal/handlers"
	log "github.com/aliereno/go-graphql-server/internal/logger"
	"github.com/aliereno/go-graphql-server/internal/orm"
	"github.com/aliereno/go-graphql-server/pkg/middlewares"
	"github.com/gin-gonic/gin"
)

// Run spins up the server
func Run(orm *orm.ORM) {

	endpoint := "http://" + host + ":" + port

	r := gin.Default()
	r.Use(middlewares.GinContextToContextMiddleware())
	// Handlers
	// Simple keep-alive/ping handler
	r.GET("/ping", handlers.Ping())
	// GraphQL handlers
	// Playground handler
	if isPgEnabled {
		r.GET(gqlPgPath, handlers.PlaygroundHandler(gqlPath))
		log.Info("GraphQL Playground @ " + endpoint + gqlPgPath)
	}
	// Pass in the ORM instance to the GraphqlHandler
	r.POST(gqlPath, handlers.GraphqlHandler(orm))
	// Run the server
	// Inform the user where the server is listening
	log.Info("GraphQL @ " + endpoint + gqlPath)
	log.Info("Running @ " + endpoint)
	// Print out and exit(1) to the OS if the server cannot run
	log.Fatal(r.Run(host + ":" + port))
}
