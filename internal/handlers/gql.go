package handlers

import (
	"context"
	"fmt"
	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/handler"
	"github.com/aliereno/go-graphql-server/internal/gql"
	"github.com/aliereno/go-graphql-server/internal/gql/resolvers"
	"github.com/aliereno/go-graphql-server/internal/handlers/auth"
	"github.com/aliereno/go-graphql-server/internal/orm"
	"github.com/gin-gonic/gin"
)

// GraphqlHandler defines the GQLGen GraphQL server handler
func GraphqlHandler(orm *orm.ORM) gin.HandlerFunc {
	// NewExecutableSchema and Config are in the generated.go file
	c := gql.Config{
		Resolvers: &resolvers.Resolver{
			ORM: orm, // pass in the ORM instance in the resolvers to be used
		},
	}
	c.Directives.IsAuthenticated = func(ctx context.Context, obj interface{}, next graphql.Resolver) (interface{}, error) {
		err := auth.LookToken(ctx)
		if err != nil {
			// block calling the next resolver
			return nil, fmt.Errorf(err.Error())
		}
		// or let it pass through
		return next(ctx)
	}
	h := handler.GraphQL(gql.NewExecutableSchema(c))
	//h := handler.GraphQL(gql.NewExecutableSchema(c),handler.ComplexityLimit(1))

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

// PlaygroundHandler defines a handler to expose the Playground
func PlaygroundHandler(path string) gin.HandlerFunc {
	h := handler.Playground("Go GraphQL Server", path)
	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}
