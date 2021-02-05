package server

import (
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/kalradev/review-central/graphql-server/graph"
	"github.com/kalradev/review-central/graphql-server/graph/generated"
)

// GetGraphQLHandler for graphql server
func GetGraphQLHandler() *handler.Server {
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))
	return srv
}

// GetPlaygroundHandler for api testing
func GetPlaygroundHandler(APIEndpoint string) http.HandlerFunc {
	return playground.Handler("GraphQL playground", APIEndpoint)
}
